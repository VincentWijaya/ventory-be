package handler

import (
	"context"
	"net/http"

	"github.com/vincentwijaya/ventory-be/constant/errs"
	ctxHelp "github.com/vincentwijaya/ventory-be/pkg/context"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) SessionCheck(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionString := r.Header.Get("Authorization")
		log.Info("VALIDATE SESSION --> ", sessionString)

		userData, err := m.middleware.ValidateSession(ctx, sessionString)
		if err != nil {
			log.Error("Failed to authenticate user session with error: ", err)
			writeResponse(w, nil, errs.InvalidJwt)
			return
		}

		contextWithUserID := context.WithValue(r.Context(), ctxHelp.UserIDKey, userData.UserID)
		contextWithUserRole := context.WithValue(contextWithUserID, ctxHelp.UserRoleKey, userData.Role)
		contextWithSession := context.WithValue(contextWithUserRole, ctxHelp.SessionCTKey, sessionString)
		r = r.WithContext(contextWithSession)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
