package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port              string
	MongoURI          string
	MongoDatabase     string
	AppEnv            string
	DailyWorkflowHour int
}

func Load() Config {
	loadDotEnv(".env")

	hour, err := strconv.Atoi(getenv("DAILY_WORKFLOW_HOUR", "7"))
	if err != nil || hour < 0 || hour > 23 {
		hour = 7
	}

	return Config{
		Port:              getenv("PORT", "8080"),
		MongoURI:          os.Getenv("MONGODB_URI"),
		MongoDatabase:     getenv("MONGODB_DATABASE", "fashion_dashboard"),
		AppEnv:            getenv("APP_ENV", "development"),
		DailyWorkflowHour: hour,
	}
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)
		if key == "" || os.Getenv(key) != "" {
			continue
		}
		_ = os.Setenv(key, value)
	}
}

func (c Config) Addr() string {
	return ":" + c.Port
}

func (c Config) IsTest() bool {
	return c.AppEnv == "test"
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
