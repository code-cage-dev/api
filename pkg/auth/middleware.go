package auth

import (
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
)

func Check(msg string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("access_token")
		if token == "" {
			return result.Error(msg, fiber.StatusUnauthorized)
		}
		ctx.Locals("token", token)
		return nil
	}
}
