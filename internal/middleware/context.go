package middleware

import (
	"context"
	"pleasurelove/internal/constanta"
	"pleasurelove/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CopyLocalsToContext(c *fiber.Ctx, keys ...constanta.ContextKey) {
	ctx := c.UserContext()

	for _, key := range keys {
		if val := c.Locals(key); val != nil {
			ctx = context.WithValue(ctx, key, val)
		}
	}

	c.SetUserContext(ctx)
}

func SetTraceIDAndRequestIDMiddleware(c *fiber.Ctx) error {
	traceID := utils.GenerateTraceID()     // Fungsi untuk membuat trace_id unik
	requestID := utils.GenerateRequestID() // Fungsi untuk membuat request_id unik

	// Simpan trace_id dan request_id ke dalam context
	ctx := c.UserContext()
	ctx = context.WithValue(ctx, constanta.TraceID, traceID)
	ctx = context.WithValue(ctx, constanta.RequestID, requestID)
	c.SetUserContext(ctx)

	return c.Next()
}
