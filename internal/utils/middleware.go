package utils

import (
	"fmt"
	"strings"
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func MiddlewareJWTAuthorization(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Unauthorized",
			"message": "You are not logged in",
		})
	}

	containerConf := container.InitContainer()

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(containerConf.Apps.SecretJwt), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)

	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid token 2",
		})
	}

	c.Locals("user", claims)

	return c.Next()
}
