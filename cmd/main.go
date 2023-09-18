package main

import (
	"golang.org/x/exp/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	sqlite "url-shortener/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	os.Setenv("CONFIG_PATH", "/Users/v/go/src/url-shortener/config/local.yaml")
	cfg := config.MustLoad()

	//fmt.Println(cfg)
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Info("test", slog.String("perem", "perem1"))
	log.Debug("debug massages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	//TODO: init config: cleanenv

	//TODO: init logger: slog
	//TODO: init storage: sqllite
	//TODO: init router: chi, "chi render"
	//TODO: init server:

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
