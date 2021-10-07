package main

import (
	"Rest-api/Database"
	"Rest-api/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	var PORT = ":2004"

	Database.Get()

	r:=mux.NewRouter()
	r.HandleFunc("/product",controller.GetAllProduct).Methods("GET")
	r.HandleFunc("/product/{id}",controller.GetProductById).Methods("GET")
	r.HandleFunc("/product",controller.AddAllProduct).Methods("POST")
	r.HandleFunc("/topfive",controller.GetTopFiveProduct).Methods("GET")
	r.HandleFunc("/purchase",controller.Purchase).Methods("POST")
	log.Fatalln(http.ListenAndServe(PORT,r))

}
