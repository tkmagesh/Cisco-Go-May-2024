package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products = []Product{
	Product{Id: 100, Name: "Pen", Cost: 10},
	Product{Id: 101, Name: "Pencil", Cost: 5},
	Product{Id: 102, Name: "Marker", Cost: 50},
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	if id, err := strconv.Atoi(idString); err == nil {
		for _, p := range products {
			if p.Id == id {
				if err := json.NewEncoder(w).Encode(p); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
}

// middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%s\t%s\n", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", GetAllProductsHandler).Methods(http.MethodGet)
	r.HandleFunc("/products", CreateProductHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{id:[0-9]+}", GetProductByIdHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
