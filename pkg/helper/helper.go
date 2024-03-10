package helper

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/volkankocali/hotel-store-case-go/pkg/config"
	"github.com/volkankocali/hotel-store-case-go/pkg/schema"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Helper struct {
	cfg config.Config
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		cfg: *config,
	}
}

type AutoCustomClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (helper *Helper) GenerateTokenClients(user schema.UserSchemaResponse) (string, error) {
	claims := &AutoCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1000).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("hotelstorecasego"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (helper *Helper) GeneratePasswordHash(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return string(hashedPass), nil
}

func (helper *Helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func (helper *Helper) GenerateReferenceCode() string {
	return "REF-" + time.Now().Format("20060102150405")
}

func (helper *Helper) ParseToken(tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("hotelstorecasego"), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	fmt.Println("claims", claims)
	return claims, nil
}
