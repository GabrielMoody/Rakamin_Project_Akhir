package utils

import (
	"time"
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToKen(userid uint) (string, error) {
	containerConf := container.InitContainer()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userid
	claims["exp"] = time.Now().UTC().Add(1 * time.Hour).Unix()
	claims["iat"] = time.Now().UTC().Unix()

	signedToken, err := token.SignedString([]byte(containerConf.Apps.SecretJwt))

	return signedToken, err
}
