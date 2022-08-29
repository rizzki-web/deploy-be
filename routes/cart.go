package routes

import (
	"waysbuck/handlers"
	mysql "waysbuck/pkg"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", h.FindCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
	r.HandleFunc("/cart", h.CleaningCart).Methods("DELETE")

}
