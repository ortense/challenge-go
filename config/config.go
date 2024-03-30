package config

import "os"

type Config struct {
	Hostname       string
	DataSourceName string
}

func getEnv() Config {
	host := os.Getenv("HOSTNAME")
	database := os.Getenv("DATABASE")

	return Config{
		Hostname:       host,
		DataSourceName: database,
	}
}

func GetConfig() Config {
	defaultConfig := Config{
		Hostname:       "localhost:9090",
		DataSourceName: "./bin/challenge.db",
	}

	env := getEnv()

	if env.Hostname != "" {
		defaultConfig.Hostname = env.Hostname
	}

	if env.DataSourceName != "" {
		defaultConfig.DataSourceName = env.DataSourceName
	}

	return defaultConfig
}
