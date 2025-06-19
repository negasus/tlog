package main

import (
	"context"
	"log/slog"

	"github.com/negasus/tlog"
)

func main() {
	opts := &tlog.Options{
		AttrColor:  tlog.TextColorBlack,
		ShowSource: true,
		TrimSource: true,
		Level:      slog.LevelDebug,
	}

	h := tlog.New(opts)
	h.SetLevel(slog.LevelDebug)

	slog.SetDefault(slog.New(h))

	slog.Debug("Hello, world!", "key", "value")
	slog.Info("Hello, world! Barb", "key", "value")
	slog.Error("Hello, world! Fo", "key", "value")
	slog.Log(context.Background(), -2, "Hello, world!", slog.String("key", "value"))

	slog.Warn("With tag 1", "key", "value", "a", 42.4, "tag", "TAG") // This will not show the tag
	h.TagOn("TAG")
	slog.Warn("With tag 2", "key", "value", "a", 42.4, "tag", "TAG") // This will show the tag
	h.TagOff("TAG")
	slog.Warn("With tag 3", "key", "value", "a", 42.4, "tag", "TAG") // This will not show the tag
}
