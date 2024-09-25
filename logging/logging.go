package logging

// LogLevel представляет уровень важности сообщения лога.
type LogLevel int

const (
	// Trace - самый низкий уровень логирования, используется для очень подробной отладочной информации.
	Trace LogLevel = iota

	// Debug используется для отладочной информации, полезной разработчикам.
	Debug

	// Information используется для общей информации о работе приложения.
	Information

	// Warning используется для логирования потенциальных проблем или нежелательных ситуаций.
	Warning

	// Fatal используется для критических ошибок, которые могут привести к завершению работы приложения.
	Fatal

	// None используется для отключения логирования.
	None
)

// Logger - это интерфейс, определяющий методы для логирования сообщений различных уровней.
//
// Реализации этого интерфейса должны обеспечивать логирование сообщений
// в соответствии с заданным уровнем важности.
type Logger interface {
	// Trace логирует сообщение с уровнем Trace.
	Trace(string)
	// Tracef логирует форматированное сообщение с уровнем Trace.
	Tracef(string, ...interface{})

	// Debug логирует сообщение с уровнем Debug.
	Debug(string)
	// Debugf логирует форматированное сообщение с уровнем Debug.
	Debugf(string, ...interface{})

	// Info логирует сообщение с уровнем Information.
	Info(string)
	// Infof логирует форматированное сообщение с уровнем Information.
	Infof(string, ...interface{})

	// Warn логирует сообщение с уровнем Warning.
	Warn(string)
	// Warnf логирует форматированное сообщение с уровнем Warning.
	Warnf(string, ...interface{})

	// Panic логирует сообщение с уровнем Fatal и может вызвать панику.
	Panic(string)
	// Panicf логирует форматированное сообщение с уровнем Fatal и может вызвать панику.
	Panicf(string, ...interface{})
}
