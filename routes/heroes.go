package routes

import (
	"encoding/json"
	"net/http"

	"github.com/barretodaniel/hots-tier-list-api/models"

	"github.com/gorilla/mux"
)

// HeroesRouter creates a subrouter for the /heroes endpoint
func HeroesRouter(r *mux.Router) {
	s := r.PathPrefix("/heroes").Subrouter()
	s.HandleFunc("/", heroesHandler).Methods("GET")
}

func heroesHandler(w http.ResponseWriter, r *http.Request) {

	heroes, err := models.GetHeroes()
	res, err := json.Marshal(heroes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(res)
}
