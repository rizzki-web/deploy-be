package routes

import (
	"github.com/gorilla/mux"
)

func RouterInit(r *mux.Router) {
	UserRoutes(r)
	ProductRoutes(r)
	TopingRoutes(r)
	TransactionRoutes(r)
	CartRoutes(r)
	AuthRoutes(r)

}
