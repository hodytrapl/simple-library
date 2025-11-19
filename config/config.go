package config

import "errors"

func GetPortFromConfig(config map[string]string) (string, error) {
	if val, ok := config["PORT"]; ok {
		return val, nil
	} else {
		return "", errors.New("Ключ 'PORT' отсутсвует в конфигурации")
	}
}
