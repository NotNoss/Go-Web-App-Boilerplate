package middleware

import "github.com/gofiber/fiber/v2"

func RequireAuth(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strics-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SOMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// Go to next middleware:
	return c.Next()
}