package main

import (
	"log/slog"
	"os"
	"scraper-url/internal/config"
	"scraper-url/internal/crawler/spider"
	"scraper-url/internal/lib/logger/slogpretty"
	"scraper-url/internal/tcp"
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

	//run crawler
	crawler := spider.New()

	//run tcp server
	server := tcp.New(log, cfg.Address, crawler)

	crawler.Scan(cfg.Url, cfg.Depth)

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
