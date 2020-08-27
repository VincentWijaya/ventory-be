package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vincentwijaya/ventory-be/constant/errs"

	"github.com/vincentwijaya/ventory-be/constant"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	ctxHelp "github.com/vincentwijaya/ventory-be/pkg/context"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) Register(w http.ResponseWriter, r *http.Request) {
	var (
		request  entity.RegisterRequest
		err      error
		response interface{}
	)
	ctx := r.Context()
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(400)
		log.Error("Failed to decode register wallet request to json")
		writeResponse(w, nil, err)
		return
	}

	log.Infof("Request Register: %+v", request)

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

	err = m.user.Register(ctx, request)

	writeResponse(w, response, err)
}
