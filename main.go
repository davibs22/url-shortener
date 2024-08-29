// Package main implements the entry point for the Url-Shortener application.
// 
// Url-Shortener is a simple URL shortening service written in Go using the Fiber framework.
// This service shortens long URLs and stores them in Redis, providing a fast and scalable 
// solution for URL management.
//
// Author: Davi Bomfim Santiago
// License: MIT License
// Date: August 16, 2024

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"url-shortener/config"
	"url-shortener/router"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:               "Url-Shortener",
		DisableStartupMessage: false,
	})

	config.LoadEnv()

	router.InitializeRoutes(app)

	app.Listen(fmt.Sprintf("%s:%s", config.GetEnv("APP_HOST", "127.0.0.1"), config.GetEnv("APP_PORT", "80")))
}
