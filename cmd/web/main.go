package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	Logger *slog.Logger
	Cfg    config
}

type config struct {
	addr      string
	staticDir string
}

func main() {

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	app := &application{
		Logger: logger,
		Cfg:    cfg,
	}

	logger.Info("starting server", "addr", cfg.addr)

	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
