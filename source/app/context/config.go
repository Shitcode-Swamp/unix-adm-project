package appctx

import (
	"os"
	"strconv"
)

var Cfg = newConfig()

type Config struct {
	MongoURI      string
	MongoDB       string
	JWTSecret     string
	JWTTTLHours   int
	AdminUsername string
	AdminPassword string
	ServerPort    string
	TLSCert       string
	TLSKey        string
}

func newConfig() Config {
	return Config{
		MongoURI:      mustEnv("MONGO_URI"),
		MongoDB:       envOrDefault("MONGO_DB", "secrets_registry"),
		JWTSecret:     mustEnv("JWT_SECRET"),
		JWTTTLHours:   atoiOrDefault(os.Getenv("JWT_TTL_HOURS"), 24),
		AdminUsername: mustEnv("ADMIN_USERNAME"),
		AdminPassword: mustEnv("ADMIN_PASSWORD"),
		ServerPort:    envOrDefault("SERVER_PORT", "80"),
		TLSCert:       os.Getenv("TLS_CERT"),
		TLSKey:        os.Getenv("TLS_KEY"),
	}
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("required env var " + key + " is not set")
	}
	return v
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func atoiOrDefault(s string, def int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return def
}
