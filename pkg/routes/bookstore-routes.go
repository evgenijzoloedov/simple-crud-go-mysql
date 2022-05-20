package routes

import (
	"github.com/evgenijzoloedov/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
}
