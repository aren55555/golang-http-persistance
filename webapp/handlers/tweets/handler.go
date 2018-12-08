package tweets

import (
	"database/sql"
	"net/http"

	"github.com/aren55555/learn/dec7/webapp/handlers/helpers"
)

type Handler struct {
	PG *sql.DB
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var action http.HandlerFunc
	switch r.Method {
	case http.MethodGet:
		action = h.list
	case http.MethodPost:
		action = h.create
	default:
		action = helpers.MethodNotAllowed
	}
	action(w, r)
}
