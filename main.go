package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"` 
	Lastname string `json:"lastname"`
}

var movies []Movie

//getMovies

func getMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//delete a movie

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index , item := range movies {
		if item.ID  == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...) 
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

//get a single movie
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	   

}

//create a movie
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

 	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"]{

			movies = append(movies[:index], movies[index + 1:]...) 
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}

	}
}


func main() {
	r := mux.NewRouter()

	// append movies

	movies = append(movies, Movie{ID:"1", Isbn:"165874", Title:"MOvie 1", Director : &Director{Firstname: "abel", Lastname: "eloh"}})
	movies = append(movies, Movie{ID:"2", Isbn:"16754862", Title:"MOvie 2", Director : &Director{Firstname: "abeloh", Lastname: "abeloa"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}