package main

import (
	"encoding/json"
	"os"
	"path"
)

const (
	perm = 750

	DEFAULT_MaxLayer = 100
)

func cfgPath() string {
	if path := os.Getenv("ATWS_CONFIG_PATH"); path != "" {
		return path
	}

	if home := os.Getenv("HOME"); home != "" {
		return path.Join(home, ".config/atws/config.json")
	}

	panic("can't get config path")
}

type Config struct {
	Layer     int
	Workspace int

	MaxLayer int
}

func openConfig() (*os.File, error) {
	p := cfgPath()

	dir, _ := path.Split(p)
	err := os.MkdirAll(dir, perm)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(p, os.O_CREATE|os.O_RDWR, perm)
}

func GetConfig() (cfg *Config, err error) {
	f, err := openConfig()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg = new(Config)
	dec := json.NewDecoder(f)
	err = dec.Decode(cfg)

	if cfg.MaxLayer == 0 {
		cfg.MaxLayer = DEFAULT_MaxLayer
	}

	return cfg, err
}

func SetConfig(cfg *Config) error {
	f, err := openConfig()
	if err != nil {
		return err
	}
	defer f.Close()

	f.Seek(0, 0)

	enc := json.NewEncoder(f)
	return enc.Encode(cfg)
}
