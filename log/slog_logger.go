package log

import (
	"sync"
	"time"

	"golang.org/x/exp/slog"
)

// NewSlogLogger return Logger implementer that use stdlib slog as the backend.
func NewSlogLogger(wr ...Writer) Logger {
	singletonLogger = &slogLog{wr: wr, mu: sync.Mutex{}}
	return singletonLogger
}

type slogLog struct {
	log *multiSlog
	wr  []Writer
	mu  sync.Mutex
}

func (s *slogLog) Init(dur time.Duration) {
	var slogs multiSlog
	for _, w := range s.wr {
		switch w.Output() {
		case CONSOLE:
			opt := &slog.HandlerOptions{Level: toSlogLevel(w.Level())}
			slogs.loggers = append(slogs.loggers, slog.New(slog.NewTextHandler(w.Writer(), opt)))
		case FILE, NEWRELIC:
			opt := &slog.HandlerOptions{Level: toSlogLevel(w.Level())}
			slogs.loggers = append(slogs.loggers, slog.New(slog.NewJSONHandler(w.Writer(), opt)))
		}
		w.Wait(dur)
	}
	s.log = &slogs
}
func (s *slogLog) Flush(dur time.Duration) {
	for _, w := range s.wr {
		w.Flush(dur)
	}
}
func (s *slogLog) With(pr ...Log) Logger {
	if len(pr) == 0 {
		return s
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.log = s.log.With(toSlogAttr(pr)...)
	return s
}
func (s *slogLog) Dbg(msg string, pr ...Log) {
	if len(pr) > 0 {
		s.log.Debug(msg, toSlogAttr(pr)...)
		return
	}
	s.log.Debug(msg)
}
func (s *slogLog) Inf(msg string, pr ...Log) {
	if len(pr) > 0 {
		s.log.Info(msg, toSlogAttr(pr)...)
		return
	}
	s.log.Info(msg)
}
func (s *slogLog) Wrn(msg string, pr ...Log) {
	if len(pr) > 0 {
		s.log.Warn(msg, toSlogAttr(pr)...)
		return
	}
	s.log.Warn(msg, toSlogAttr(pr)...)
}
func (s *slogLog) Err(msg string, pr ...Log) {
	if len(pr) > 0 {
		s.log.Error(msg, toSlogAttr(pr)...)
		return
	}
	s.log.Error(msg)
}

// toSlogLevel transform local log Level to slog level.
func toSlogLevel(lvl Level) slog.Level {
	switch lvl {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	}
	return -1
}

// toSlogAttr transform local Log to slog field.
func toSlogAttr(pr []Log) []any {
	var attrs []any
	for _, p := range pr {
		switch p.typ {
		case StringType:
			attrs = append(attrs, slog.String(p.key, p.str))
		case NumType:
			attrs = append(attrs, slog.Int(p.key, p.num))
		case FloatType:
			attrs = append(attrs, slog.Float64(p.key, p.flt))
		case BoolType:
			attrs = append(attrs, slog.Bool(p.key, p.b))
		case AnyType:
			attrs = append(attrs, slog.Any(p.key, p.any))
		case ErrorType:
			attrs = append(attrs, slog.Any(p.key, p.err))
		}
	}
	return attrs
}

// multiSlog add support to write logs to multiple slog.Logger.
type multiSlog struct {
	loggers []*slog.Logger
}

func (m *multiSlog) With(args ...any) *multiSlog {
	for i := range m.loggers {
		m.loggers[i] = m.loggers[i].With(args...)
	}
	return m
}
func (m *multiSlog) Debug(msg string, args ...any) {
	for _, log := range m.loggers {
		log.Debug(msg, args...)
	}
}
func (m *multiSlog) Info(msg string, args ...any) {
	for _, log := range m.loggers {
		log.Info(msg, args...)
	}
}
func (m *multiSlog) Warn(msg string, args ...any) {
	for _, log := range m.loggers {
		log.Warn(msg, args...)
	}
}
func (m *multiSlog) Error(msg string, args ...any) {
	for _, log := range m.loggers {
		log.Error(msg, args...)
	}
}
