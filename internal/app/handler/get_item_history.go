package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) GetItemHistory(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response interface{}
	)
	ctx := r.Context()

	itemID := chi.URLParam(r, "id")
	log.Infof("Request GetItemHistory: %s", itemID)

	n, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		panic("Failed to cast item id from param to int64")
	}

	response, err = m.item.GetItemHistoryByItemID(ctx, n)
	writeResponse(w, response, err)
}
