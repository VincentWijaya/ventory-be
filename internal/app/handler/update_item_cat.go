package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vincentwijaya/ventory-be/internal/entity"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) UpdateItemCategory(w http.ResponseWriter, r *http.Request) {
	var (
		request  entity.ItemCategory
		err      error
		response interface{}
	)
	ctx := r.Context()
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil && err.Error() != "EOF" {
		w.WriteHeader(400)
		log.Error("Failed to decode UpdateItemCategory request to json")
		writeResponse(w, nil, err)
		return
	}

	itemID := chi.URLParam(r, "id")
	n, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		panic("Failed to cast item id from param to int64")
	}
	request.ID = n
	log.Infof("Request UpdateItemCategory: %+v", request)

	err = m.itemCategory.UpdateItemCategory(ctx, request)

	writeResponse(w, response, err)
}
