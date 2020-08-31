package middleware

import (
	"context"

	"github.com/vincentwijaya/ventory-be/lib/jwt"
)

func (m *Module) ValidateSession(ctx context.Context, sessionString string) (res *jwt.JWTValidateResponse, err error) {
	res, err = jwt.JWTValidate(sessionString, m.jwtSecret)

	return
}
