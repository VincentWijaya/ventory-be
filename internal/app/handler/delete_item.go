package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response interface{}
	)
	ctx := r.Context()

	itemID := chi.URLParam(r, "id")
	log.Infof("Request DeleteItem: %s", itemID)

	n, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		panic("Failed to cast item id from params to int64")
	}
	err = m.item.DeleteItem(ctx, n)
	writeResponse(w, response, err)
}
