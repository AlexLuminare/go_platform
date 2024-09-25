// Package logging предоставляет интерфейс для логирования и его стандартную реализацию.
//
// Этот пакет позволяет создавать логгеры с различными уровнями логирования,
// такими как Trace, Debug, Information, Warning и Fatal. Он также предоставляет
// возможность настройки минимального уровня логирования и поведения при фатальных ошибках.

package logging

import (
	"go_platform/config"
	"log"
	"os"
	"strings"
)

// NewDefaultLogger создает и возвращает новый экземпляр DefaultLogger на основе предоставленной конфигурации.
//
// Параметры:
//   - cfg: объект Configuration, содержащий настройки для логгера.
//
// Возвращает:
//   - Logger: интерфейс Logger, реализованный структурой DefaultLogger.
//
// Функция выполняет следующие действия:
// 1. Устанавливает флаги для форматирования вывода лога (временная метка и префикс сообщения).
// 2. Определяет минимальный уровень логирования из конфигурации или использует Debug по умолчанию.
// 3. Создает логгеры для каждого уровня логирования, используя стандартный вывод (os.Stdout).
// 4. Настраивает логгер на вызов паники для сообщений уровня Fatal.
//
// Созданный логгер использует различные префиксы для каждого уровня логирования (TRACE, DEBUG, INFO, WARN, FATAL).
func NewDefaultLogger(cfg config.Configuration) Logger {
	flags := log.Lmsgprefix | log.Ltime

	var level LogLevel = Debug
	if configLevelString, found :=
		cfg.GetString("logging:level"); found {
		level = LogLevelFromString(configLevelString)
	}

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

// LogLevelFromString преобразует строковое представление уровня логирования в соответствующее значение LogLevel.
//
// Параметры:
//   - val: строка, представляющая уровень логирования.
//
// Возвращает:
//   - level: соответствующее значение LogLevel.
//
// Функция нечувствительна к регистру и поддерживает следующие значения:
//   - "debug" -> Debug
//   - "information" -> Information
//   - "warning" -> Warning
//   - "fatal" -> Fatal
//   - "none" -> None
//
// Если входная строка не соответствует ни одному из известных уровней,
// функция возвращает уровень Debug по умолчанию.

func LogLevelFromString(val string) (level LogLevel) {
	switch strings.ToLower(val) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":

		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}
