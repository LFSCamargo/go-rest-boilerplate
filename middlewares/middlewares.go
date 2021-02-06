package middlewares

import (
	"github.com/LFSCamargo/go-rest-boilerplate/auth"
	"github.com/LFSCamargo/go-rest-boilerplate/types"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware - A function that handles the token parsing and stores the user as a local variable
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		c.Status(401)
		c.JSON(types.Error{
			Message: "Not Authorized",
		})
		return err
	}
	c.Locals("user", user)
	return c.Next()
}
