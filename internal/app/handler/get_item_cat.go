package handler

import (
	"net/http"
	"strconv"

	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) GetItemCategory(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response interface{}
	)
	ctx := r.Context()

	var (
		lastID      int64 = 0
		dataPerPage int64 = 10
	)

	lastIDFromQuery := r.URL.Query().Get("lastId")
	dataPerPageFromQuery := r.URL.Query().Get("max")

	log.Infof("Request GetItemCategory: %s", r.URL.Query().Encode())

	if lastIDFromQuery != "" {
		n, err := strconv.ParseInt(lastIDFromQuery, 10, 64)
		if err != nil {
			panic("Failed to cast last id from query to int64")
		}
		lastID = n
	}
	if dataPerPageFromQuery != "" {
		n, err := strconv.ParseInt(dataPerPageFromQuery, 10, 64)
		if err != nil {
			panic("Failed to cast data per page from query to int64")
		}
		dataPerPage = n
	}

	response, err = m.itemCategory.GetItemCategory(ctx, dataPerPage, lastID)
	writeResponse(w, response, err)
}
