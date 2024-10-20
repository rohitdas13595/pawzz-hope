package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AdminCreateToken struct {
	Id    string
	Email string
}

func CreateAdminToken(data *AdminCreateToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"id":    data.Id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Unix(),
		"sub":   "admin",
		"aud":   "admin",
		"iss":   "encore.dev",
	})

	tokenString, err := token.SignedString(GetSettings().JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyAdminToken(tokenString string) (*AdminCreateToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSettings().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &AdminCreateToken{
			Id:    claims["id"].(string),
			Email: claims["email"].(string),
		}, nil
	}
	return nil, err
}

//create User token

type UserAccesTokenData struct {
	Id    string
	Email string
	Type  string
}

type UserRefreshTokenData struct {
	Id  string
	Tid string
}

func CreateUserAccessToken(data *UserAccesTokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"type":  data.Type,
		"id":    data.Id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Unix(),
		"sub":   "user",
		"aud":   "user",
		"iss":   "encore.dev",
	})

	tokenString, err := token.SignedString([]byte(GetSettings().JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyUserAccessToken(tokenString string) (*UserAccesTokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSettings().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &UserAccesTokenData{
			Id:    claims["id"].(string),
			Email: claims["email"].(string),
			Type:  claims["type"].(string),
		}, nil
	}
	return nil, err
}

// User  refresh token

func CreateUserRefreshToken(data *UserRefreshTokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"tid": data.Tid,
		"id":  data.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": "user",
		"aud": "user",
		"iss": "encore.dev",
	})

	tokenString, err := token.SignedString([]byte(GetSettings().JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyUserRefreshToken(tokenString string) (*UserRefreshTokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSettings().JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &UserRefreshTokenData{
			Id:  claims["id"].(string),
			Tid: claims["tid"].(string),
		}, nil
	}
	return nil, err
}
