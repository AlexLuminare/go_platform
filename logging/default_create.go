// Package logging предоставляет интерфейс для логирования и его стандартную реализацию.
//
// Этот пакет позволяет создавать логгеры с различными уровнями логирования,
// такими как Trace, Debug, Information, Warning и Fatal. Он также предоставляет
// возможность настройки минимального уровня логирования и поведения при фатальных ошибках.

package logging

import (
	"log"
	"os"
)

// NewDefaultLogger создает и возвращает новый экземпляр DefaultLogger с указанным минимальным уровнем логирования.
//
// Параметры:
//   - level: минимальный уровень логирования для нового логгера.
//
// Возвращает:
//   - Logger: интерфейс Logger, реализованный структурой DefaultLogger.
//
// Созданный логгер использует стандартный вывод (os.Stdout) для всех уровней логирования.
// Он включает временную метку в выводе лога и использует различные префиксы для каждого уровня логирования.
// По умолчанию, логгер вызывает панику для логов уровня Fatal.
func NewDefaultLogger(level LogLevel) Logger {
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel](*log.Logger){
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},

		triggerPanic: true,
	}
}
