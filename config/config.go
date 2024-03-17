package config

import (
	"os"
	"strconv"
)

var (
	Port             = getEnv("METRICS_PORT", "9090")
	ProtectedMetrics = getEnvAsBool("METRICS_PROTECTED", true)
	MetricsUsername  = getEnv("METRICS_USERNAME", "metricsUser")
	MetricsPassword  = getEnv("METRICS_PASSWORD", "MetricsVeryHardPassword")
	UpdateInterval   = getEnvAsInt("UPDATE_INTERVAL", 30)
	TimeZone         = getEnv("TIMEZONE", "UTC")
	InactivityTime   = getEnvAsInt("INACTIVITY_TIME", 2)
	BaseURL          = getEnv("MARZBAN_BASE_URL", "")
	ApiUsername      = getEnv("MARZBAN_USERNAME", "")
	ApiPassword      = getEnv("MARZBAN_PASSWORD", "")
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue
}
