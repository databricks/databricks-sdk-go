package main

import (
	"context"
	"flag"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/databricks/log"
	"golang.org/x/exp/slog"
)

// slogAdapter makes an slog.Logger usable with the Databricks SDK.
type slogAdapter struct {
	*slog.Logger
}

func (s *slogAdapter) Log(ctx context.Context, level log.Level, format string, a ...any) {
	slevel := slog.LevelDebug
	switch level {
	case log.LevelInfo:
		slevel = slog.LevelInfo
	case log.LevelWarn:
		slevel = slog.LevelWarn
	case log.LevelError:
		slevel = slog.LevelError
	}
	s.Logger.Log(ctx, slevel, format, a...)
}

func main() {
	flag.Parse()

	// Define global logger for the SDK.
	// This instance is used when no context is present (e.g. upon construction),
	// or when a context is present but doesn't hold a logger to use.
	log.SetDefaultLogger(&slogAdapter{
		slog.Default().With("global", true),
	})

	logKey := &struct{}{} // define a unique key for the context value
	ctx := context.WithValue(context.Background(), logKey, &slogAdapter{
		slog.Default().With("global", false),
	})

	// Construct client and make a request.
	// Observe a mix of global=true and global=false in the log output.
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(ctx)
	if err != nil {
		panic(err)
	}
}
