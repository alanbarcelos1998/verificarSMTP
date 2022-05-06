package main

import (
	"fmt"
	"log"
	"net/http"
	route "smtp-api/routes"

	"github.com/gorilla/mux"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func configServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonMiddleware)
	route.Config(router)

	fmt.Println("Server is running port 8785")
	log.Fatal(http.ListenAndServe(":8785", router))
}

func main() {
	configServer()
}
