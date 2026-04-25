package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadReadsDotEnv(t *testing.T) {
	dir := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWd) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"PORT", "MONGODB_URI", "MONGODB_DATABASE", "APP_ENV", "DAILY_WORKFLOW_HOUR"} {
		t.Setenv(key, "")
		_ = os.Unsetenv(key)
	}
	env := "PORT=19090\nMONGODB_URI=mongodb://localhost:27017\nMONGODB_DATABASE=fashion_dashboard_test\nAPP_ENV=test\nDAILY_WORKFLOW_HOUR=9\n"
	if err := os.WriteFile(filepath.Join(dir, ".env"), []byte(env), 0600); err != nil {
		t.Fatal(err)
	}

	cfg := Load()
	if cfg.Port != "19090" || cfg.MongoURI == "" || cfg.MongoDatabase != "fashion_dashboard_test" || cfg.AppEnv != "test" || cfg.DailyWorkflowHour != 9 {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}
