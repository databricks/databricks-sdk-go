package log

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

func TestLevel_order(t *testing.T) {
	var wantOrder = []Level{
		LevelTrace,
		LevelDebug,
		LevelInfo,
		LevelWarn,
		LevelError,
	}

	for i := 1; i < len(wantOrder); i++ {
		if wantOrder[i-1] >= wantOrder[i] {
			t.Errorf("Level %s should be lower than %s", wantOrder[i-1].String(), wantOrder[i].String())
		}
	}
}

func TestLogger_Log(t *testing.T) {
	testCases := []struct {
		minLevel Level
		level    Level
		format   string
		args     []any
		want     string
	}{
		{
			minLevel: LevelInfo,
			level:    LevelTrace, // drop lower levels
			format:   "test %s",
			args:     []any{"message"},
			want:     "",
		},
		{
			minLevel: LevelInfo,
			level:    LevelError, // keep higher levels
			format:   "test %s",
			args:     []any{"message"},
			want:     "[ERROR] test message\n",
		},
		{
			minLevel: LevelTrace,
			level:    LevelTrace,
			format:   "test %s",
			args:     []any{"message"},
			want:     "[TRACE] test message\n",
		},
		{
			minLevel: LevelDebug,
			level:    LevelDebug,
			format:   "test %s",
			args:     []any{"message"},
			want:     "[DEBUG] test message\n",
		},
		{
			minLevel: LevelInfo,
			level:    LevelInfo,
			format:   "test %s",
			args:     []any{"message"},
			want:     "[INFO] test message\n",
		},
		{
			minLevel: LevelWarn,
			level:    LevelWarn,
			format:   "test %s",
			args:     []any{"message"},
			want:     "[WARN] test message\n",
		},
		{
			minLevel: LevelError,
			level:    LevelError,
			format:   "test %s",
			args:     []any{"message"},
			want:     "[ERROR] test message\n",
		},
		{
			minLevel: LevelInfo,
			level:    Level(42), // unknown level
			format:   "test %s",
			args:     []any{"message"},
			want:     "[UNKNOWN] test message\n",
		},
	}

	for _, tc := range testCases {
		out := &bytes.Buffer{}
		l := logger{
			level: tc.minLevel,
			out:   out,
		}

		l.Log(context.Background(), tc.level, tc.format, tc.args...)

		if got := out.String(); got != tc.want {
			t.Errorf("Log(): got %s, want %s", got, tc.want)
		}
	}
}

func TestNew(t *testing.T) {
	testCases := []struct {
		level Level
	}{
		{level: LevelTrace},
		{level: LevelDebug},
		{level: LevelInfo},
		{level: LevelWarn},
		{level: LevelError},
	}

	for _, tc := range testCases {
		l := New(tc.level).(*logger)

		if l.level != tc.level {
			t.Errorf("New(%s): got minLevel %s, want %s", tc.level, l.level, tc.level)
		}
	}
}

func TestSetAndGetDefaultLogger(t *testing.T) {
	old := defaultLogger
	t.Cleanup(func() { defaultLogger = old })

	m := &logger{}
	SetDefaultLogger(m)

	if DefaultLogger() != m {
		t.Error("SetDefaultLogger did not set the default logger correctly")
	}
}

type params struct {
	ctx    context.Context
	level  Level
	format string
	a      []any
}

type mockLogger struct {
	params
}

func (m *mockLogger) Log(ctx context.Context, level Level, format string, a ...any) {
	m.ctx = ctx
	m.level = level
	m.format = format
	m.a = a
}

func (m *mockLogger) MinLevel() Level {
	return LevelTrace
}

// loggedParams returns the last set of parameters that were logged (if any)
// by executing logFunc. It returns a zero params struct if logFunc did not
// log anything.
func loggedParams(logFunc func()) params {
	old := defaultLogger
	defer func() { defaultLogger = old }()
	m := &mockLogger{}
	defaultLogger = m
	logFunc()
	return m.params
}

func TestLogFunctions(t *testing.T) {
	testCases := []struct {
		logFunc func(string, ...any)
		level   Level
	}{
		{logFunc: Trace, level: LevelTrace},
		{logFunc: Debug, level: LevelDebug},
		{logFunc: Info, level: LevelInfo},
		{logFunc: Warning, level: LevelWarn},
		{logFunc: Error, level: LevelError},
	}

	for _, tc := range testCases {
		want := params{
			ctx:    context.Background(),
			level:  tc.level,
			format: "test %s",
			a:      []any{"message"},
		}

		got := loggedParams(func() {
			tc.logFunc("test %s", "message")
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestLogContextFunctions(t *testing.T) {
	type key int // dummy context key type
	testCases := []struct {
		logFunc func(context.Context, string, ...any)
		level   Level
	}{
		{logFunc: TraceContext, level: LevelTrace},
		{logFunc: DebugContext, level: LevelDebug},
		{logFunc: InfoContext, level: LevelInfo},
		{logFunc: WarningContext, level: LevelWarn},
		{logFunc: ErrorContext, level: LevelError},
	}

	for _, tc := range testCases {
		t.Run(tc.level.String(), func(t *testing.T) {
			want := params{
				ctx:    context.WithValue(context.Background(), key(42), "value"),
				level:  tc.level,
				format: "test %s",
				a:      []any{"message"},
			}

			got := loggedParams(func() {
				tc.logFunc(want.ctx, "test %s", "message")
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}
