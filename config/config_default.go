package config

import "strings"

// DefaultConfig представляет собой реализацию интерфейса Configuration,
// которая хранит конфигурационные данные в виде вложенных map[string]interface{}.
type DefaultConfig struct {
	configData map[string]interface{}
}

// get является внутренним методом для получения значения из конфигурации по имени.
// Он поддерживает доступ к вложенным секциям через разделитель ":".
//
// Параметры:
//   - name: строка, представляющая путь к значению в конфигурации.
//
// Возвращаемые значения:
//   - result: интерфейс{}, содержащий найденное значение.
//   - found: булево значение, указывающее, было ли найдено значение.
func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

// GetSection возвращает подсекцию конфигурации по указанному имени.
//
// Параметры:
//   - name: строка, представляющая путь к секции в конфигурации.
//
// Возвращаемые значения:
//   - section: объект Configuration, представляющий найденную секцию.
//   - found: булево значение, указывающее, была ли найдена секция.
func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := c.get(name)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		}
	}
	return
}

// GetString возвращает строковое значение из конфигурации по указанному имени.
//
// Параметры:
//   - name: строка, представляющая путь к значению в конфигурации.
//
// Возвращаемые значения:
//   - result: строковое значение из конфигурации.
//   - found: булево значение, указывающее, было ли найдено значение.
func (c *DefaultConfig) GetString(name string) (result string, found bool) {
	value, found := c.get(name)
	if found {
		result = value.(string)
	}
	return
}

// GetInt возвращает целочисленное значение из конфигурации по указанному имени.
//
// Параметры:
//   - name: строка, представляющая путь к значению в конфигурации.
//
// Возвращаемые значения:
//   - result: целочисленное значение из конфигурации.
//   - found: булево значение, указывающее, было ли найдено значение.
func (c *DefaultConfig) GetInt(name string) (result int, found bool) {
	value, found := c.get(name)
	if found {
		result = int(value.(float64))
	}
	return
}

// GetBool возвращает булево значение из конфигурации по указанному имени.
//
// Параметры:
//   - name: строка, представляющая путь к значению в конфигурации.
//
// Возвращаемые значения:
//   - result: булево значение из конфигурации.
//   - found: булево значение, указывающее, было ли найдено значение.
func (c *DefaultConfig) GetBool(name string) (result bool, found bool) {
	value, found := c.get(name)
	if found {
		result = value.(bool)
	}
	return
}

// GetFloat возвращает значение с плавающей точкой из конфигурации по указанному имени.
//
// Параметры:
//   - name: строка, представляющая путь к значению в конфигурации.
//
// Возвращаемые значения:
//   - result: значение с плавающей точкой из конфигурации.
//   - found: булево значение, указывающее, было ли найдено значение.
func (c *DefaultConfig) GetFloat(name string) (result float64, found bool) {
	value, found := c.get(name)
	if found {
		result = value.(float64)
	}
	return
}
