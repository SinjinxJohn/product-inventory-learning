package config

import (
	"os"
	"strconv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser               string
	DBPassword           string
	DBAddress            string
	DBName               string
	JWTExpirationSeconds int64
	JWTSecret            string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost:           getEnv("PUBLIC_HOST", "localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "root"),
		DBPassword:           getEnv("DB_PASSWORD", "password"),
		DBAddress:            getEnv("DB_ADDRESS", "localhost:3306"),
		DBName:               getEnv("DB_NAME", "event"),
		JWTExpirationSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:            getEnv("JWt_SECRET", "hotlinebling"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i

	}
	return fallback
}
