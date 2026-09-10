package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	JWTSecret      string
	JWTExpireHours int

	CORSAllowOrigin string
	RateLimitPerMin int

	WAProvider   string // "stub" or "fonnte"
	FonnteToken  string
	FonnteAPIURL string
}

func Load() *Config {
	_ = godotenv.Load() // optional for local dev

	return &Config{
		AppPort:        getEnv("APP_PORT", "8080"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "portaljob"),
		DBPassword:     getEnv("DB_PASSWORD", "portaljob_secret"),
		DBName:         getEnv("DB_NAME", "portaljob"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		JWTSecret:      getEnv("JWT_SECRET", "change-this-in-production"),
		JWTExpireHours: getEnvInt("JWT_EXPIRE_HOURS", 24),

		CORSAllowOrigin: getEnv("CORS_ALLOW_ORIGIN", "*"),
		RateLimitPerMin: getEnvInt("RATE_LIMIT_PER_MIN", 120),

		WAProvider:   getEnv("WA_PROVIDER", "stub"),
		FonnteToken:  getEnv("FONNTE_TOKEN", ""),
		FonnteAPIURL: getEnv("FONNTE_API_URL", "https://api.fonnte.com/send"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
