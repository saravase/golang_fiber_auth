package middleware

import (
	"golang_fiber_auth/auth-api/handler"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JWTMiddleware used to validate JWT is valid or invalid
func JWTMiddleware(c *fiber.Ctx) error {

	authStatus := c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Can't access without authorization",
	})

	tokenString := c.Cookies("token")
	if tokenString == "" {
		return authStatus
	}

	token, _ := jwt.ParseWithClaims(tokenString, &handler.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	claims, ok := token.Claims.(*handler.AuthClaims)
	if !(ok && token.Valid) {
		return authStatus
	}

	c.Locals("user_id", claims.UserID)
	return c.Next()
}
