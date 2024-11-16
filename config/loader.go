package config

import (
	"encoding/json"
	"io"
	"os"
)

func Load(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}
	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
