package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"hots-tier-list/services"

	"github.com/gorilla/mux"
)

var hs *services.HeroService

// HeroesRouter creates a subrouter for the /heroes endpoint
func HeroesRouter(r *mux.Router, db *sql.DB) {
	hs = services.GetHeroService(db)
	s := r.PathPrefix("/heroes").Subrouter()
	s.HandleFunc("/", heroesHandler)
}

func heroesHandler(w http.ResponseWriter, r *http.Request) {

	heroes, err := hs.GetHeroes()
	res, err := json.Marshal(heroes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(res)
}
