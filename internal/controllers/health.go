package controllers

import (
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var startTime = time.Now() // Waktu aplikasi mulai berjalan

type HealthDependencies struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func HealthCheck(deps HealthDependencies) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Periksa koneksi database
		dbStatus := "OK"
		sqlDB, err := deps.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "DOWN"
		}

		// Periksa koneksi Redis
		redisStatus := "OK"
		if deps.Redis != nil {
			_, err := deps.Redis.Ping(context.Background()).Result()
			if err != nil {
				redisStatus = "DOWN"
			}
		}

		// Response status kesehatan
		healthStatus := map[string]interface{}{
			"status":  "OK",
			"uptime":  time.Since(startTime).String(),
			"version": "1.0.0", // Ganti dengan versi aplikasi Anda
			"dependencies": map[string]string{
				"database": dbStatus,
				"redis":    redisStatus,
			},
		}

		// Jika ada dependency yang "DOWN", ubah status menjadi 500
		if dbStatus == "DOWN" || redisStatus == "DOWN" {
			return c.Status(fiber.StatusInternalServerError).JSON(healthStatus)
		}

		return c.Status(fiber.StatusOK).JSON(healthStatus)
	}
}
