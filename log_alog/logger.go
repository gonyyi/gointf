package log_alog 

import (
	"github.com/gonyyi/alog"
	"io"
)

func NewLogWithAlog(w io.Writer) *Logger {
	log := alog.New(w)
	log.Control.Level = alog.TraceLevel
	
	l := Logger{
		alog: &log,
	}
	
	return &l
}

type Logger struct {
	alog *alog.Logger
	tag  struct {
		USR, STO, REQ, SYS int64
	}
}

func (l *Logger) Logger() *alog.Logger {
	return l.alog
}

func (l *Logger) Trace(tag int64, action string, detail string) {
	l.alog.Trace(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
func (l *Logger) Debug(tag int64, action string, detail string) {
	l.alog.Debug(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
func (l *Logger) Info(tag int64, action string, detail string) {
	l.alog.Info(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
func (l *Logger) Warn(tag int64, action string, detail string) {
	l.alog.Warn(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
func (l *Logger) Error(tag int64, action string, detail string) {
	l.alog.Error(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
func (l *Logger) Fatal(tag int64, action string, detail string) {
	l.alog.Fatal(alog.Tag(tag)).Str("action", action).Str("detail", detail).Write("")
}
