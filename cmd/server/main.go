package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fashion_dashboard/internal/config"
	"fashion_dashboard/internal/db"
	"fashion_dashboard/internal/repository"
	"fashion_dashboard/internal/scheduler"
	"fashion_dashboard/internal/web"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	client, err := db.Connect(ctx, cfg.MongoURI)
	if err != nil {
		log.Printf("mongodb unavailable, using in-memory store: %v", err)
	}
	defer func() {
		if err := db.Disconnect(context.Background(), client); err != nil {
			log.Printf("mongodb disconnect: %v", err)
		}
	}()

	var store repository.Store = repository.NewMemoryStore()
	if client != nil {
		mongoStore := repository.NewMongoStore(client, cfg.MongoDatabase)
		if err := mongoStore.EnsureIndexes(ctx); err != nil {
			log.Printf("mongodb indexes unavailable, using in-memory store: %v", err)
		} else {
			store = mongoStore
		}
	}

	app, err := web.NewServer(store)
	if err != nil {
		log.Fatal(err)
	}

	if !cfg.IsTest() {
		scheduler.Start(context.Background(), cfg.DailyWorkflowHour, scheduler.NewWorkflow(store))
	}

	server := &http.Server{Addr: cfg.Addr(), Handler: app.Handler(), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		log.Printf("listening on http://localhost%s", cfg.Addr())
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
