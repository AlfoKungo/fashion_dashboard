package scheduler

import (
	"context"
	"log"
	"time"
)

type Runner interface {
	Run(context.Context) error
}

func Start(ctx context.Context, hour int, runner Runner) {
	if runner == nil {
		return
	}
	go func() {
		for {
			wait := time.Until(nextRun(time.Now(), hour))
			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				timer.Stop()
				return
			case <-timer.C:
				if err := runner.Run(ctx); err != nil {
					log.Printf("daily workflow: %v", err)
				}
			}
		}
	}()
}

func nextRun(now time.Time, hour int) time.Time {
	run := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
	if !run.After(now) {
		run = run.Add(24 * time.Hour)
	}
	return run
}
