package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// BuildTokenClaims ...
// tokenExpiry in minutes
func BuildTokenClaims(userID int64, tokenExpiry int64, email, userRole string) jwt.MapClaims {
	expireAt := time.Now().Add(time.Duration(tokenExpiry) * time.Minute)
	claims := jwt.MapClaims{
		"exp": expireAt.Unix(),
		"uid": userID,
		"r":   userRole,
		"em":  email,
	}
	return claims
}

func CreateToken(claims jwt.MapClaims, jwtSignKey string) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := newToken.SignedString([]byte(jwtSignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTValidate(requestToken, jwtSignKey string) (*JWTValidateResponse, error) {
	parsedToken, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		return &JWTValidateResponse{
			UserID:    int64(claims["uid"].(float64)),
			ExpiresIn: int64(claims["exp"].(float64)),
			Role:      claims["r"].(string),
			Email:     claims["em"].(string),
		}, nil
	}

	return nil, fmt.Errorf("token invalid or expired")
}
