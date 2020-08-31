package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) InsertItem(w http.ResponseWriter, r *http.Request) {
	var (
		request  entity.Item
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

	log.Infof("Request InsertItem: %+v", request)

	err = m.item.InsertItem(ctx, request)

	writeResponse(w, response, err)
}
