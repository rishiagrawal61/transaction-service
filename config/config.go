package config

import "os"

type Config struct {
	Port          string
	EnableRestApi bool

	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() Config {
	return Config{Port: envStr("PORT", "8080"),
		EnableRestApi: envBool("ENABLE_REST_API", true),
		DB: DBConfig{
			Host:     envStr("DB_HOST", "127.0.0.1"),
			Port:     envStr("DB_PORT", "3306"),
			User:     envStr("DB_USER", "root"),
			Password: envStr("DB_PASSWORD", ""),
			Name:     envStr("DB_NAME", "test"),
		},
	}
}

func envBool(key string, defaultVal bool) bool {
	v := os.Getenv(key)
	if v == "true" {
		return true
	}
	if v == "false" {
		return false
	}
	return defaultVal
}

func envStr(key string, defaultVal string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	return v
}
