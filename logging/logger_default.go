// Package logging предоставляет интерфейс для логирования и его стандартную реализацию.
//
// Этот пакет позволяет создавать логгеры с различными уровнями логирования,
// такими как Trace, Debug, Information, Warning и Fatal. Он также предоставляет
// возможность настройки минимального уровня логирования и поведения при фатальных ошибках.

package logging

import (
	"fmt"
	"log"
)

// DefaultLogger представляет собой стандартную реализацию интерфейса Logger.
//
// Он поддерживает несколько уровней логирования и позволяет настраивать
// минимальный уровень для вывода сообщений. Также имеет возможность
// вызывать панику при логировании фатальных ошибок.
type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel](*log.Logger)
	triggerPanic bool
}

// MinLogLevel возвращает текущий минимальный уровень логирования.
func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

// write записывает сообщение в лог, если его уровень не ниже минимального.
func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

// Trace записывает сообщение с уровнем Trace.
func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

// Tracef записывает форматированное сообщение с уровнем Trace.
func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

// Debug записывает сообщение с уровнем Debug.
func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

// Debugf записывает форматированное сообщение с уровнем Debug.
func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

// Info записывает сообщение с уровнем Information.
func (l *DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

// Infof записывает форматированное сообщение с уровнем Information.
func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(template, vals...))
}

// Warn записывает сообщение с уровнем Warning.
func (l *DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

// Warnf записывает форматированное сообщение с уровнем Warning.
func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(template, vals...))
}

// Panic записывает сообщение с уровнем Fatal и вызывает панику, если это настроено.
func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

// Panicf записывает форматированное сообщение с уровнем Fatal и вызывает панику, если это настроено.
func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formattedMsg := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMsg)
	if l.triggerPanic {
		panic(formattedMsg)
	}
}
