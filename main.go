package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// The recipe Type (more like an object)
type Recipe struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Method      string `json:"method,omitempty"`
	Ingredients string `json:"ingredients,omitempty"`
}

var recipes []Recipe

// Display all from the recipes var
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(recipes)
}

// Display a single data
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range recipes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Recipe{})
}

// create a new item
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe Recipe
	_ = json.NewDecoder(r.Body).Decode(&recipe)
	recipe.ID = params["id"]
	recipes = append(recipes, recipe)
	json.NewEncoder(w).Encode(recipes)
}

// Delete an item
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range recipes {
		if item.ID == params["id"] {
			recipes = append(recipes[:index], recipes[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(recipes)
	}
}

// main function to boot up everything
func main() {
	//	router := mux.NewRouter()
	recipes = append(recipes, Recipe{ID: "1", Name: "John", Method: "mix stuff", Ingredients: "Stuff"})

	//log.Fatal(http.ListenAndServe(":8000", router))

	//mux := http.NewServeMux()
	//mux.HandleFunc("/recipes", GetRecipes).Methods("GET")
	//mux.HandleFunc("/recipes/{id}", GetRecipe).Methods("GET")
	//mux.HandleFunc("/recipes/{id}", CreateRecipe).Methods("POST")
	//mux.HandleFunc("/recipes/{id}", DeleteRecipe).Methods("DELETE")

	//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//       w.Header().Set("Content-Type", "application/json")
	//       w.Write([]byte("{\"hello\": \"world\"}"))
	//   })

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	//handler := cors.Default().Handler(mux)
	//http.ListenAndServe(":8000", handler)

	mux := http.NewServeMux()
	mux.HandleFunc("/recipes", GetRecipes)
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write([]byte("{\"hello\": \"world\"}"))
	//})

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)

}
