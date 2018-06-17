package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"hots-tier-list/services"

	"github.com/gorilla/mux"
)

var rs *services.RoleService

// RolesRouter creates a subrouter for the /roles path
func RolesRouter(r *mux.Router, db *sql.DB) {
	rs = services.GetRoleService(db)
	s := r.PathPrefix("/roles").Subrouter()
	s.HandleFunc("/", rolesHandler)
}

func rolesHandler(w http.ResponseWriter, r *http.Request) {

	roles, err := rs.GetRoles()
	res, err := json.Marshal(roles)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(res)

}
