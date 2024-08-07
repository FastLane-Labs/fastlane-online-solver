package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// PoolsConfig is the configuration for pools.
// Used to generate pool instances given the FactoryAddress and DexType
type PoolsConfig struct {
	DexName        string `yaml:"dex_name"`
	DexType        string `yaml:"dex_type"`
	FactoryAddress string `yaml:"factory_address"`
}

func parsePoolsConfig(filename string) ([]*PoolsConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf struct {
		Pools []*PoolsConfig `yaml:"pools"`
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return conf.Pools, nil
}
