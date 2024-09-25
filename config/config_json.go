package config

import (
	"encoding/json"
	"os"
	"strings"
)

// Load загружает конфигурацию из JSON-файла и возвращает объект Configuration.
//
// Параметры:
//   - fileName: путь к JSON-файлу конфигурации.
//
// Возвращаемые значения:
//   - config: объект Configuration, содержащий загруженную конфигурацию.
//   - err: ошибка, если произошла проблема при чтении или разборе файла.
//
// Функция выполняет следующие шаги:
// 1. Читает содержимое файла.
// 2. Декодирует JSON в map[string]interface{}.
// 3. Создает объект DefaultConfig с полученными данными.
//
// Если на любом этапе возникает ошибка, она возвращается вместе с nil-конфигурацией.
func Load(fileName string) (config Configuration, err error) {
	var data []byte
	data, err = os.ReadFile(fileName)
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err = decoder.Decode(&m)
		if err == nil {
			config = &DefaultConfig{configData: m}
		}
	}
	return
}
