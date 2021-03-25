package log_alog

import (
	"github.com/gonyyi/alog"
)

func AlogAdopter(l alog.Logger) *Logger {
	return &Logger{alog: &l}
}

type Logger struct {
	alog *alog.Logger
}

func (l *Logger) Trace(tag int64, action string, detail string) {
	if detail != "" {
		l.alog.Trace(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Trace(alog.Tag(tag)).Str("action", action).Write("")
}
func (l *Logger) Debug(tag int64, action string, detail string) {
	if detail != ""{
		l.alog.Debug(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Debug(alog.Tag(tag)).Str("action", action).Write("")
}
func (l *Logger) Info(tag int64, action string, detail string) {
	if detail != "" {
		l.alog.Info(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Info(alog.Tag(tag)).Str("action", action).Write("")
}
func (l *Logger) Warn(tag int64, action string, detail string) {
	if detail != "" {
		l.alog.Warn(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Warn(alog.Tag(tag)).Str("action", action).Write("")
}
func (l *Logger) Error(tag int64, action string, detail string) {
	if detail != "" {
		l.alog.Error(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Error(alog.Tag(tag)).Str("action", action).Write("")
}
func (l *Logger) Fatal(tag int64, action string, detail string) {
	if detail != "" {
		l.alog.Fatal(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
		return
	}
	l.alog.Fatal(alog.Tag(tag)).Str("action", action).Write("")
}
