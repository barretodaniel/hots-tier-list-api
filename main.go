package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"hots-tier-list/routes"

	"github.com/gorilla/mux"
	// required for db access in routes
	_ "github.com/lib/pq"
)

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Sorry, the resource you are requesting doesn't exist.")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	connStr := "dbname=hots sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	check(err)

	defer db.Close()

	var h http.Handler = http.HandlerFunc(pageNotFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	routes.HeroesRouter(r, db)
	routes.RolesRouter(r, db)

	fmt.Println("Serving on port 3000")
	http.ListenAndServe(":3000", r)
}
