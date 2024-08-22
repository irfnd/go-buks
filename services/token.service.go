package services

import (
	"errors"
	"fmt"
	"go-buks/configs"
	"go-buks/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJwt(payload *models.TokenPayload) string {
	v, err := time.ParseDuration(configs.Env.JWT_EXP)
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"id":  payload.ID,
		"email": payload.Email,
	})

	token, err := t.SignedString([]byte(configs.Env.JWT_KEY))
	if err != nil {
		panic(err)
	}

	return token
}

func VerifyJwt(token string) (*models.TokenPayload, error) {
	var (
		id    uuid.UUID
		email string
	)

	parsed, err := parse(token)
	if err != nil {
		return nil, errors.New("token invalid or expired")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	for key, value := range claims {
		switch key {
		case "id":
			idStr, ok := value.(string)
			if !ok {
				return nil, errors.New("ID claim missing or invalid")
			}
			id, err = uuid.Parse(idStr)
			if err != nil {
				return nil, errors.New("invalid UUID format")
			}
		case "email":
			email, ok = value.(string)
			if !ok {
				return nil, errors.New("email claim missing or invalid")
			}
		}
	}

	return &models.TokenPayload{
		ID:    id,
		Email: email,
	}, nil
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(configs.Env.JWT_KEY), nil
	})
}