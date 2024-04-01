package config

import (
	"errors"
	"os"
	"path"

	"github.com/naoina/toml"
)

type apiInfo struct {
	url string
	key string
}

type tradeInfo struct {
	upperLine uint
	lowerLine uint
	ticker    string
}

type Config struct {
	api   apiInfo
	trade tradeInfo
}

func NewConfig(filePath string) (*Config, error) {
	if filePath == "" {
		filePath = "./config.toml"
	}

	if file, err := os.Open(path.Join(filePath)); err != nil {
		return nil, err
	} else {
		defer file.Close()
		cfg := new(Config)
		if err := toml.NewDecoder(file).Decode(cfg); err != nil {
			return nil, err
		} else {
			if !cfg.checkValidInputs() {
				return nil, errors.New("invalid config input")
			}
			return cfg, nil
		}
	}
}

func (c *Config) Url() string {
	return c.api.url
}

func (c *Config) UpperLine() uint {
	return c.trade.upperLine
}

func (c *Config) LowerLine() uint {
	return c.trade.lowerLine
}

func (c *Config) Ticker() string {
	return c.trade.ticker
}

func (c *Config) checkValidInputs() bool {
	if c.api.key == "" ||
		c.api.url == "" ||
		c.trade.lowerLine <= c.trade.lowerLine ||
		c.trade.ticker == "" {
		return false
	}
	return true
}
