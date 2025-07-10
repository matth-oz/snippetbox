package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
	"github.com/matth-oz/snippetbox/internal/models"
)

type application struct {
	Logger   *slog.Logger
	Cfg      config
	Snippets *models.SnippetModel
}

type config struct {
	addr      string
	staticDir string
	dsn       string
}

func main() {

	var cfg config

	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static/", "Path to static assets")
	// user:pass@server/db_name
	flag.StringVar(&cfg.dsn, "dsn", "mattweb_snipbox:c5T9iBFD@tcp(5.23.51.100:3306)/mattweb_snipbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	db, err := openDB(cfg.dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("DB connected", "mysql", db.Stats().OpenConnections)

	defer db.Close()

	app := &application{
		Logger:   logger,
		Cfg:      cfg,
		Snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", "addr", cfg.addr)

	err = http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
