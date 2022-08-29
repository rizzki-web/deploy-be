package routes

import (
	"waysbuck/handlers"
	mysql "waysbuck/pkg"
	"waysbuck/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	r.HandleFunc("/transaction/", h.UpdateTransaction).Methods("PATCH")
	r.HandleFunc("/transaction/{id}",h.DeleteTransaction).Methods("DELETE")

}
