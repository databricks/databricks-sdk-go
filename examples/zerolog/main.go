package main

import (
	"context"
	"flag"
	"os"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/log"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var json bool

func init() {
	flag.BoolVar(&json, "json", false, "log json messages")
}

// zerologAdapter makes an zerolog.Logger usable with the Databricks SDK.
type zerologAdapter struct {
	zerolog.Logger
}

func (z *zerologAdapter) Log(ctx context.Context, level log.Level, format string, a ...any) {
	switch level {
	case log.LevelTrace:
		z.Logger.Trace().Msgf(format, a...)
	case log.LevelDebug:
		z.Logger.Debug().Msgf(format, a...)
	case log.LevelInfo:
		z.Logger.Info().Msgf(format, a...)
	case log.LevelWarn:
		z.Logger.Warn().Msgf(format, a...)
	case log.LevelError:
		z.Logger.Error().Msgf(format, a...)
	}
}

func main() {
	flag.Parse()

	var l zerolog.Logger
	if json {
		l = zlog.Output(os.Stderr)
	} else {
		l = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Define global logger for the SDK.
	// This instance is used when no context is present (e.g. upon construction),
	// or when a context is present but doesn't hold a logger to use.
	log.SetDefaultLogger(&zerologAdapter{l.With().Bool("global", true).Logger()})

	logKey := &struct{}{} // define a unique key for the context value
	ctx := context.WithValue(context.Background(), logKey, &zerologAdapter{
		l.With().Bool("global", false).Logger(),
	})

	// Construct client and make a request.
	// Observe a mix of global=true and global=false in the log output.
	w := databricks.Must(databricks.NewWorkspaceClient())
	_, err := w.CurrentUser.Me(ctx)
	if err != nil {
		panic(err)
	}
}
