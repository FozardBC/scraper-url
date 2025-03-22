package main

import (
	"log/slog"
	"os"
	"scraper-url/internal/config"
	"scraper-url/internal/crawler/spider"
	"scraper-url/internal/lib/logger/slogpretty"
	"scraper-url/internal/netsrv/tcp"
	"scraper-url/internal/storage/files"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	// run cfg
	cfg := config.MustLoad()

	// run logger
	log := setupLogger(cfg.Env)

	// start db
	storage, err := files.New()
	if err != nil {
		log.Error("can't open file", "error", err.Error())
		os.Exit(1)
	}

	//run crawler
	crawler := spider.New()

	data, err := crawler.Scan(cfg.Url, cfg.Depth)
	if err != nil {
		log.Error("can't scan data", "error", err.Error())
		os.Exit(1)
	}

	err = storage.Save(data)
	if err != nil {
		log.Error("can't save in storage", "error", err.Error())
		os.Exit(1)
	}

	//run tcp server
	server := tcp.New(log, cfg.Address, crawler)

	server.ListenAndServe()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
