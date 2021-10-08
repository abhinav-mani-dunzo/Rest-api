package controller

import (
	"Rest-api/Service"
	"Rest-api/model"
	"Rest-api/util"
	"encoding/json"
	"log"
	"net/http"
)

func Purchase(w http.ResponseWriter, r *http.Request) {
	log.Println("Purchase hit")
	w.Header().Set("Content-Type", "application/json")
	var purchase model.Purchase
	_ = json.NewDecoder(r.Body).Decode(&purchase)
	purchase.UID = util.RandSeq(10)
	err := Service.Purchase(purchase)
	if err == nil {
		w.WriteHeader(201)
	} else {
		w.WriteHeader(400)
	}

}
