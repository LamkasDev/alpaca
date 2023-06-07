package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/LamkasDev/alpaca/cmd/common/arch"
)

type AlpacaConfig struct{}

func NewAlpacaConfig() AlpacaConfig {
	return AlpacaConfig{}
}

func ReadConfig(config AlpacaConfig) (AlpacaConfig, error) {
	path := filepath.Join(arch.DataPath, "config", "ai.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		if raw, err = json.Marshal(config); err != nil {
			return AlpacaConfig{}, err
		}
		if err = os.MkdirAll(filepath.Join(arch.DataPath, "config"), 0755); err != nil {
			return AlpacaConfig{}, err
		}
		if err = os.WriteFile(path, raw, 0755); err != nil {
			return AlpacaConfig{}, err
		}
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return AlpacaConfig{}, err
	}

	return config, nil
}
