package middleware

import (
	"fmt"
	"pleasurelove/internal/dto/response"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	ctx := utils.GetContext(c)

	logger.Error(ctx, "Unhandled error", err)
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return response.SetResponseAPI(c, code, "unhandled error", err.Error(), nil)
}

func RecoverMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Ambil lokasi panic yang relevan
				location := getRelevantPanicLocation()

				// Ambil stack trace singkat
				stackBuf := make([]byte, 1024) // Batasi ukuran stack trace
				stackSize := runtime.Stack(stackBuf, false)
				stackTrace := string(stackBuf[:stackSize])

				// Format log dengan gaya box
				box := `
┌───────────────────────────────────────────────────┐
│                   Panic Occurred!                │
├───────────────────────────────────────────────────┤
│ Location   : %s
│ Error      : %v
│ Stack Trace:
│ %s
└───────────────────────────────────────────────────┘
`
				logMessage := fmt.Sprintf(box, location, r, strings.ReplaceAll(stackTrace, "\n", "\n│ "))
				fmt.Println(logMessage) // Cetak ke terminal

				// Kirim respons umum ke client
				_ = response.SetResponseAPI(c, fiber.StatusInternalServerError, "Internal Server Error", "An unexpected error occurred. Please try again later.", nil)
			}
		}()
		return c.Next()
	}
}

// getRelevantPanicLocation mencari lokasi panic yang relevan dari stack trace
func getRelevantPanicLocation() string {
	pc := make([]uintptr, 10) // Ambil hingga 10 level stack trace
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])

	// Iterasi frame untuk menemukan lokasi yang relevan
	for {
		frame, more := frames.Next()
		// Cari lokasi di package "controllers" atau lokasi yang relevan
		if strings.Contains(frame.File, "controllers") || strings.Contains(frame.File, "router") {
			return fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
		if !more {
			break
		}
	}

	// Jika tidak ditemukan, kembalikan lokasi default
	return "unknown location"
}
