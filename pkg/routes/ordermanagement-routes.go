package routes

import (
	"github.com/gorilla/mux"
	"github.com/shadheeraNa/go-ordermanagement/pkg/controllers"
)

var RegisterOrderManagementRoutes = func(router *mux.Router) { //this function create handelers
	router.HandleFunc("/order/", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/order/", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/order/{orderID}", controllers.GetOrderById).Methods("GET")
	router.HandleFunc("/order/{orderId}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/{orderId}", controllers.DeleteOrder).Methods("DELETE")
}
