package jwt

import (
	"errors"
	"golang-agnostic-template/src/application/domain/utils"
	"golang-agnostic-template/src/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenCustomClaims struct {
	ID *string `json:"id"`
	jwt.StandardClaims
}

func CreateToken(id *string) (string, error) {
	now := time.Now().Local()
	claims := TokenCustomClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * time.Duration(config.Params.ExpiredHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Params.TokenSecret))

	return signedToken, err
}

func DecodeToken(tokenPart string) (*TokenCustomClaims, error) {
	claims := &TokenCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenPart, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Params.TokenSecret), nil
	})

	if err != nil {
		apierr := errors.New(err.Error())
		return nil, apierr
	}

	if !token.Valid {
		apierr := errors.New(utils.ErrMsgInvalidToken)
		return nil, apierr
	}

	claims, ok := token.Claims.(*TokenCustomClaims)

	if !ok {
		apierr := errors.New(utils.ErrMsgUnproccessableToken)
		return nil, apierr
	}

	return claims, err
}
