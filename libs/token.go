package libs

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	UserId uint64
	Role   string
	jwt.RegisteredClaims
}

func NewToken(id uint64, role string) *claims {

	return &claims{
		UserId: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

}

func (c *claims) Create() (string, error) {

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecret)

}

func CheckToken(token string) (*claims, error) {

	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)

	return claims, nil

}

func CodeCrypt(len int) (string, error) {

	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)

	if err != nil {

		return "", fmt.Errorf("could not hash password %w", err)

	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:len], nil

}
