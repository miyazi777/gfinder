package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	PluginConfigs []PluginConfig `json:"plugin_configs"`
}

type PluginMode string

var (
	PLUGIN_MODE_JSON PluginMode = "json"
	PLUGIN_MODE_ROW  PluginMode = "row"
)

type PluginConfig struct {
	Path       string     `json:"path"`
	PluginMode PluginMode `json:"plugin_mode"`
}

func NewConfig() (*Config, error) {
	var config Config

	// TODO: 設定ファイルのパスを環境変数から取得するようにする
	f, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
