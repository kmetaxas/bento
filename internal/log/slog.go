//go:build go1.21

package log

import (
	"fmt"
	"log/slog"
	"os"
)

type logHandler struct {
	slog *slog.Logger
}

func NewBentoLogAdapter(l *slog.Logger) *logHandler {
	return &logHandler{slog: l}
}

func (l *logHandler) WithFields(fields map[string]string) Modular {
	tmp := l.slog
	for k, v := range fields {
		tmp = tmp.With(slog.String(k, v))
	}

	c := l.clone()
	c.slog = tmp
	return c
}

func (l *logHandler) With(keyValues ...any) Modular {
	c := l.clone()
	c.slog = l.slog.With(keyValues...)
	return c
}

func (l *logHandler) Fatal(format string, v ...any) {
	l.slog.Error(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *logHandler) Error(format string, v ...any) {
	l.slog.Error(fmt.Sprintf(format, v...))
}

func (l *logHandler) Warn(format string, v ...any) {
	l.slog.Warn(fmt.Sprintf(format, v...))
}

func (l *logHandler) Info(format string, v ...any) {
	l.slog.Info(fmt.Sprintf(format, v...))
}

func (l *logHandler) Debug(format string, v ...any) {
	l.slog.Debug(fmt.Sprintf(format, v...))
}

func (l *logHandler) Trace(format string, v ...any) {
	l.slog.Debug(fmt.Sprintf(format, v...))
}

func (l *logHandler) clone() *logHandler {
	c := *l
	return &c
}
