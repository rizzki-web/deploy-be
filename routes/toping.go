package routes

import (
	"waysbuck/handlers"
	mysql "waysbuck/pkg"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	topingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(topingRepository)

	r.HandleFunc("/topings", h.FindTopings).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	r.HandleFunc("/toping", h.CreateToping).Methods("POST")
	r.HandleFunc("/toping/{id}", h.UpdateToping).Methods("PATCH")
	r.HandleFunc("/toping/{id}", h.DeleteToping).Methods("DELETE")

}
