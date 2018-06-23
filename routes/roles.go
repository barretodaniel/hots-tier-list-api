package routes

import (
	"encoding/json"
	"net/http"

	"github.com/barretodaniel/hots-tier-list-api/models"

	"github.com/gorilla/mux"
)

// RolesRouter creates a subrouter for the /roles path
func RolesRouter(r *mux.Router) {
	s := r.PathPrefix("/roles").Subrouter()
	s.HandleFunc("/", rolesHandler).Methods("GET")
}

func rolesHandler(w http.ResponseWriter, r *http.Request) {

	roles, err := models.GetRoles()
	res, err := json.Marshal(roles)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(res)

}
