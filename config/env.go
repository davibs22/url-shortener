// Package main implements the entry point for the Url-Shortener application.
// 
// Url-Shortener is a simple URL shortening service written in Go using the Fiber framework.
// This service shortens long URLs and stores them in Redis, providing a fast and scalable 
// solution for URL management.
//
// Author: Davi Bomfim Santiago
// License: MIT License
// Date: August 16, 2024

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetEnv returns the value of an environment variable or a default value
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
