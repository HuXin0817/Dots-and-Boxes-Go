package config

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

var configFilePath = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "config.json"
	}
	dir := filepath.Join(home, ".dots-and-boxes")
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
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
	if err = os.WriteFile(configFilePath, content, 0644); err != nil {
		return err
	}
	return nil
}

var defaultConf = Config{
	AI1:     false,
	AI2:     false,
	AI1Name: "L3",
	AI2Name: "L3",
}

var Conf = func() Config {
	var c Config
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		return defaultConf
	}
	if err = sonic.Unmarshal(content, &c); err != nil {
		return defaultConf
	}
	return c
}()
