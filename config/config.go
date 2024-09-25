// Package config предоставляет интерфейс для работы с конфигурацией приложения.
//
// Этот пакет определяет интерфейс Configuration, который может быть реализован
// различными провайдерами конфигурации для обеспечения гибкости в управлении
// настройками приложения.
package config

// Configuration - это интерфейс, определяющий методы для получения
// значений конфигурации различных типов.
//
// Реализации этого интерфейса должны обеспечивать доступ к конфигурационным
// данным, поддерживая различные типы данных и возможность получения
// значений по умолчанию.
type Configuration interface {
	// GetString возвращает строковое значение конфигурации по имени.
	// Если значение не найдено, found будет false.
	GetString(name string) (configValue string, found bool)

	// GetInt возвращает целочисленное значение конфигурации по имени.
	// Если значение не найдено, found будет false.
	GetInt(name string) (configValue int, found bool)

	// GetBool возвращает булево значение конфигурации по имени.
	// Если значение не найдено, found будет false.
	GetBool(name string) (configValue bool, found bool)

	// GetFloat возвращает значение с плавающей точкой конфигурации по имени.
	// Если значение не найдено, found будет false.
	GetFloat(name string) (configValue float64, found bool)

	// GetStringDefault возвращает строковое значение конфигурации по имени.
	// Если значение не найдено, возвращается значение по умолчанию defVal.
	GetStringDefault(name, defVal string) (configValue string)

	// GetIntDefault возвращает целочисленное значение конфигурации по имени.
	// Если значение не найдено, возвращается значение по умолчанию defVal.
	GetIntDefault(name string, defVal int) (configValue int)

	// GetBoolDefault возвращает булево значение конфигурации по имени.
	// Если значение не найдено, возвращается значение по умолчанию defVal.
	GetBoolDefault(name string, defVal bool) (configValue bool)

	// GetFloatDefault возвращает значение с плавающей точкой конфигурации по имени.
	// Если значение не найдено, возвращается значение по умолчанию defVal.
	GetFloatDefault(name string, defVal float64) (configValue float64)

	// GetSection возвращает подсекцию конфигурации по имени.
	// Если секция не найдена, found будет false.
	GetSection(sectionName string) (section Configuration, found bool)
}
