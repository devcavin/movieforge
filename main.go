package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/devcavin/movieforge/movie"
	"github.com/gorilla/mux"
)

var movies []Movie

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/home", home).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", addMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Movie server started at port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}  