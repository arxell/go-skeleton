package config

import (
	"os"
	"strings"
)

// Config ...
type Config struct {
	Hostport string
}

// NewConfig ...
func NewConfig() (*Config, error) {
	c := &Config{
		Hostport: envString("Hostport", ":9090"),
	}
	return c, nil
}

func envString(key, defaultVal string) string {
	key = strings.ToUpper(key)
	result := os.Getenv(key)
	if len(result) == 0 {
		return defaultVal
	}
	return result
}
