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
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	// RDB is the Redis client that can be used in other parts of your code
	RDB *redis.Client
	// RedisCtx is the default context for Redis operations
	RedisCtx context.Context
)

func InitRedis() {
	LoadEnv()

	redisHost := GetEnv("REDIS_HOST", "127.0.0.1")
	redisPort := GetEnv("REDIS_PORT", "6379")
	redisPass := GetEnv("REDIS_PASSWORD", "")

	// Create the Redis client
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort), // Redis address
		Password: redisPass,                                  // password, if any
		DB:       0,                                          // database to be used
	})

	// Initialize the context for Redis operations
	RedisCtx = context.Background()
}
