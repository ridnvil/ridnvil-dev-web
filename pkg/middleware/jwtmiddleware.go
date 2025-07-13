package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"ridnvil-dev/pkg/auth"
)

var jwtSecretKey = []byte("ridwan")

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Cookies("token")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errors.New("missing token")})
		}

		tokenString := authHeader

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errors.New("invalid token")})
		}

		if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
			c.Locals("ID", claims.ID)
			c.Locals("email", claims.Email)
			c.Locals("name", claims.Name)
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": errors.New("invalid token")})
	}
}
