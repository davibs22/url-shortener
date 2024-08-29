// Package main implements the entry point for the Url-Shortener application.
// 
// Url-Shortener is a simple URL shortening service written in Go using the Fiber framework.
// This service shortens long URLs and stores them in Redis, providing a fast and scalable 
// solution for URL management.
//
// Author: Davi Bomfim Santiago
// License: MIT License
// Date: August 16, 2024

package router

import (
	"github.com/gofiber/fiber/v2"
	"url-shortener/app/controllers"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/:hash", controllers.ShowReduce)
	app.Post("/shorten", controllers.StoreReduce)
	app.Delete("/shorten", controllers.DestroyReduce)
}
