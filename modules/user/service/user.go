package userService

import (
	"github.com/LFSCamargo/go-rest-boilerplate/auth"
	userModel "github.com/LFSCamargo/go-rest-boilerplate/database/models/user"
	userStructs "github.com/LFSCamargo/go-rest-boilerplate/modules/user/entity"
	"github.com/LFSCamargo/go-rest-boilerplate/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// HandleLogin - function that handles the login
func HandleLogin(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(userStructs.LoginInput)

	err := c.BodyParser(body)

	if err != nil {
		c.Status(403)
		return c.JSON(types.Error{
			Message: "The request body is malformed",
		})
	}

	user, error := userModel.FindByEmail(body.Email)

	if error != nil {
		c.Status(401)
		return c.JSON(types.Error{
			Message: "Invalid email or password",
		})
	}

	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if passwordError != nil {
		c.Status(401)
		return c.JSON(types.Error{
			Message: "Invalid email or password",
		})
	}

	token, tokenErr := auth.SignToken(body.Email)

	if tokenErr != nil {
		c.Status(500)
		return c.JSON(types.Error{
			Message: "Internal Server Error",
		})
	}

	c.Status(200)
	return c.JSON(userStructs.AuthenticationOutput{
		Token: token,
	})
}

// HandleSignUp - function that handles the sign up
func HandleSignUp(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := new(userStructs.RegisterInput)

	err := c.BodyParser(body)

	if err != nil {
		c.Status(403)
		return c.JSON(types.Error{
			Message: "The request body is malformed",
		})
	}

	user, _ := userModel.FindByEmail(body.Email)

	if user != nil {
		c.Status(401)
		return c.JSON(types.Error{
			Message: "User Already Registered",
		})
	}

	registered, err := userModel.CreateNewUser(body.Name, body.Email, body.Password)

	if err != nil {
		c.Status(500)
		return c.JSON(types.Error{
			Message: "Internal Server Error",
		})
	}

	token, tokenErr := auth.SignToken(registered.Email)

	if tokenErr != nil {
		c.Status(500)
		return c.JSON(types.Error{
			Message: "Internal Server Error",
		})
	}

	c.Status(200)
	return c.JSON(userStructs.AuthenticationOutput{
		Token: token,
	})
}

// HandleMe - function that handles getting the logged user from the context
func HandleMe(c *fiber.Ctx) error {
	c.Accepts("application/json")
	user := c.Locals("user").(*userModel.User)

	c.Status(200)
	return c.JSON(userStructs.User{
		Email: user.Email,
		ID:    user.ID.Hex(),
		Name:  user.Username,
	})
}
