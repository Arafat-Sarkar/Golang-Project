package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/arafat/Go-Bookstore/pkg/routes"
)

func main(){
  r:= mux.NewRouter()
  routes.RegisterBookstorRoute(r)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe("localhost:9010", r))
}
