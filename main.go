package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Init movies var as a slice Movie struct
var movies []Movie

func main() {
	// Init Router
	r := mux.NewRouter()
	createMockData()

	// Router Handlers / Endpoints
	r.HandleFunc("/api/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", createMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8200", r))
}

// Get All movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get Single Movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	//Loop movies and find with id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		} else {
			fmt.Fprintf(w, "Movie with id:"+params["id"]+" not found ")
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Create a new Movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		} else {
			fmt.Fprintf(w, "Movie with id:"+params["id"]+" not found ")
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Update Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = strconv.Itoa(rand.Intn(100000))
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)

		} else {
			fmt.Fprintf(w, "Movie with id:"+params["id"]+" not found ")
		}
	}
	json.NewEncoder(w).Encode(movies)
}
