package routes

import (
	"github.com/arafat/Go-Bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstorRoute = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.Createbook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.BookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
