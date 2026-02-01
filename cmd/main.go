package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":3333",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Failed to start", "error", err)
		os.Exit(1)
	}
}
