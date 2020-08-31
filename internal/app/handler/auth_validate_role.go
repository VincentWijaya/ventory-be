package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vincentwijaya/ventory-be/constant"
	"github.com/vincentwijaya/ventory-be/constant/errs"
	ctxHelp "github.com/vincentwijaya/ventory-be/pkg/context"
)

func (m *Module) OnlyAdmin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var (
			err      error
			response interface{}
		)

		if ctxHelp.UserRoleFromContext(ctx) != constant.AdminRole {
			session, ok := ctxHelp.GetSessionFromContext(ctx)
			if !ok || session == "" {
				err = errs.InvalidJwt
				return
			}

			err = errors.New(fmt.Sprintf("ADA YANG NYOBA-NYOBA ===> %s", session))
			writeResponse(w, response, err)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
