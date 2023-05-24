package config

import (
	"fmt"
	"os"
)

var (
	// return the listening port for the server
	PORT = getEnv("PORT", "5000")

	// DB return the sqlite database name
	DB = getEnv("DB", "todo.db")
)

func getEnv(name, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf("ENV variable not found :: %v", name))
}
