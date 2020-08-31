package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vincentwijaya/ventory-be/pkg/log"
)

func (m *Module) DeleteItemCategory(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		response interface{}
	)
	ctx := r.Context()

	itemCategoryID := chi.URLParam(r, "id")
	log.Infof("Request DeleteItemCategory: %s", itemCategoryID)

	n, err := strconv.ParseInt(itemCategoryID, 10, 64)
	if err != nil {
		panic("Failed to cast item category id from params to int64")
	}
	err = m.itemCategory.DeleteItemCategory(ctx, n)
	writeResponse(w, response, err)
}
