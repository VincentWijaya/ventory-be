package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) InsertItemCategory(w http.ResponseWriter, r *http.Request) {
	var (
		request  entity.ItemCategory
		err      error
		response interface{}
	)
	ctx := r.Context()
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(400)
		log.Error("Failed to decode InsertItemCategory request to json")
		writeResponse(w, nil, err)
		return
	}

	log.Infof("Request InsertItemCategory: %+v", request)

	err = m.itemCategory.InsertItemCategory(ctx, request)

	writeResponse(w, response, err)
}
