package main

import (
	"fmt"
	"net/http"

	"github.com/barretodaniel/hots-tier-list-api/db"
	"github.com/barretodaniel/hots-tier-list-api/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	defer db.Get().Close()

	var h http.Handler = http.HandlerFunc(pageNotFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	routes.HeroesRouter(r)
	routes.RolesRouter(r)

	fmt.Println("Serving on port 3000")
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":3000", handler)
}
