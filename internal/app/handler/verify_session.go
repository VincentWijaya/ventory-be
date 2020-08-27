package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vincentwijaya/ventory-be/constant/errs"

	"github.com/vincentwijaya/ventory-be/internal/entity"
	ctxHelp "github.com/vincentwijaya/ventory-be/pkg/context"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) VerifySession(w http.ResponseWriter, r *http.Request) {
	var (
		request entity.VerifySession
		err     error
	)

	ctx := r.Context()
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(400)
		log.Error("Failed to decode register wallet request to json")
		writeResponse(w, nil, err)
		return
	}
	log.Infof("Request Verify Session: %+v", request)

	session, ok := ctxHelp.GetSessionFromContext(ctx)
	if !ok || session == "" {
		err = errs.InvalidJwt
		writeResponse(w, request, err)
		return
	}
	if ctxHelp.UserIDFromContext(ctx) != request.UserID || ctxHelp.UserRoleFromContext(ctx) != request.Role {
		err = errs.InvalidJwt
		writeResponse(w, request, err)
		return
	}

	writeResponse(w, request, err)
}
