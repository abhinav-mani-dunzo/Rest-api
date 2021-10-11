package main

import (
	"Rest-api/Database"
	"Rest-api/controller"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log"
	"net/http"
)

func main() {
	var PORT = ":2004"

	Database.Get()
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("RestApi"),
		newrelic.ConfigLicense("eu01xxd80e779f618032cd1c674822b17ceaNRAL"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.HandleFunc(newrelic.WrapHandleFunc(app,"/product", controller.GetAllProduct)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app,"/product/{id}", controller.GetProductById)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app,"/product", controller.AddAllProduct)).Methods("PUT")
	r.HandleFunc(newrelic.WrapHandleFunc(app,"/topfive", controller.GetTopFiveProduct)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app,"/purchase", controller.Purchase)).Methods("POST")


	log.Fatalln(http.ListenAndServe(PORT, r))

}
