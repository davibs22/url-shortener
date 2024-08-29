// Package main implements the entry point for the Url-Shortener application.
// 
// Url-Shortener is a simple URL shortening service written in Go using the Fiber framework.
// This service shortens long URLs and stores them in Redis, providing a fast and scalable 
// solution for URL management.
//
// Author: Davi Bomfim Santiago
// License: MIT License
// Date: August 16, 2024

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/howeyc/crc16"
	"url-shortener/config"
)

type URLPayload struct {
	URL string `json:"url"`
}

func StoreReduce(c *fiber.Ctx) error {

	config.LoadEnv()

	// Initialize Redis
	config.InitRedis()

	// Decoding the request body into the URLPayload struct
	var payload URLPayload
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&payload)
	if err != nil {
		c.SendString("Invalid request body")
		return c.SendStatus(400)
	}

	// Checking if the URL field is empty
	if payload.URL == "" {
		c.SendString("URL is required")
		return c.SendStatus(400)
	}

	// Example data to calculate the CRC16
	data := []byte(payload.URL)

	// Calculate the CRC16
	crc := crc16.Checksum(data, crc16.IBMTable)

	// Converting the CRC16 to a string
	crcStr := fmt.Sprintf("%d", crc)

	// Example of how to set a value
	err = config.RDB.Set(config.RedisCtx, config.GetEnv("REDUCE_HOST", "127.0.0.1")+"/"+crcStr, payload.URL, 0).Err()
	if err != nil {
		c.SendString(fmt.Sprintf("Failed to set key: %v", err))
		return c.SendStatus(400)
	}

	link := config.GetEnv("REDUCE_HOST", "127.0.0.1") + "/" + crcStr

	return c.SendString(link)

}

func ShowReduce(c *fiber.Ctx) error {

	config.LoadEnv()

	// Initialize Redis
	config.InitRedis()

	link := config.GetEnv("REDUCE_HOST", "127.0.0.1") + "/" + c.Params("hash")

	// Example of how to get a value
	val, err := config.RDB.Get(config.RedisCtx, link).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to get key: %v", err))
	}

	return c.Redirect(val)

}

func DestroyReduce(c *fiber.Ctx) error {
	// Initialize Redis
	config.InitRedis()

	// Decoding the request body into the URLPayload struct
	var payload URLPayload
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&payload)
	if err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// Checking if the URL field is empty
	if payload.URL == "" {
		return c.Status(400).SendString("URL is required")
	}

	// Example of how to delete a value
	err = config.RDB.Del(config.RedisCtx, payload.URL).Err()
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("Failed to delete key: %v", err))
	}

	c.Status(200)
	return c.SendString(fmt.Sprintf("Url successfully deleted: %s", payload.URL))
}
