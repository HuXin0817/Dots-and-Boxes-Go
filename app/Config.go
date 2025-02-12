package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/bytedance/sonic"
)

const ServerAddr = "127.0.0.1:8080"

type Config struct {
	AI1     bool
	AI2     bool
	AI1Name string
	AI2Name string
}

var ConfigFilePath = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "config.json"
	}
	dir := filepath.Join(home, ".dots-and-boxes")
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "config.json"
		}
	} else if err != nil {
		return "config.json"
	}
	return filepath.Join(dir, "config.json")
}()

func (c *Config) Save() error {
	if runtime.GOOS == `js` {
		return nil
	}
	content, err := sonic.Marshal(c)
	if err != nil {
		return err
	}
	if err = os.WriteFile(ConfigFilePath, content, 0644); err != nil {
		return err
	}
	return nil
}

var defaultConf = Config{
	AI1:     false,
	AI2:     false,
	AI1Name: "L4",
	AI2Name: "L4",
}

var Conf = func() (c Config) {
	content, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return defaultConf
	}
	if err = sonic.Unmarshal(content, &c); err != nil {
		return defaultConf
	}
	return c
}()
