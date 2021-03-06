package controller

import (
	"Rest-api/Service"
	"Rest-api/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Product hit")
	res := Service.GetAllProduct()
	payload, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Product By ID hit")
	vars := mux.Vars(r)
	id := vars["id"]
	res, err := Service.GetProductById(id)
	if err != nil {
		w.WriteHeader(404)
	} else {
		payload, err := json.Marshal(res)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func AddAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Add All Product hit")
	var products []model.Product
	_ = json.NewDecoder(r.Body).Decode(&products)
	err := Service.AddAllProducts(products)
	if err == nil {
		w.WriteHeader(201)
	}
	w.WriteHeader(400)
}

func GetTopFiveProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Top 5 Product hit")
	res := Service.TopFiveProduct()
	payload, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
