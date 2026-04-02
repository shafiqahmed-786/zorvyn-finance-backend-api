package middleware

import "github.com/gofiber/fiber/v2"

func RequireRoles(roles ...string) fiber.Handler {
    allowed := make(map[string]bool)
    for _, role := range roles {
        allowed[role] = true
    }

    return func(c *fiber.Ctx) error {
        role := c.Locals("role").(string)
        if !allowed[role] {
            return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
        }
        return c.Next()
    }
}