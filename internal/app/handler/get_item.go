package handler

import (
	"net/http"

	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) GetItem(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response interface{}
	)
	ctx := r.Context()

	lastID := r.URL.Query().Get("lastId")
	dataPerPage := r.URL.Query().Get("max")

	log.Infof("Request GetItem: %s", r.URL.Query().Encode())

	_, err = m.item.GetItem(ctx, dataPerPage, lastID)
	writeResponse(w, response, err)
}
