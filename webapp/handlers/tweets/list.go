package tweets

import (
	"encoding/json"
	"net/http"

	"github.com/aren55555/learn/dec7/webapp/models"
)

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	tweets, err := models.AllTweets(h.PG)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
