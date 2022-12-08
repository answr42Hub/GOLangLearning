package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UselessFact struct {
	Facts string `json:"text"`
	Id    int    `json:"id"`
	Lang  string `json:"language"`
}

var facts = []UselessFact{
	{Facts: "The average person spends 6 months of their life waiting on a red light to turn green.", Lang: "en", Id: 0},
	{Facts: "The average person falls asleep in seven minutes.", Lang: "en", Id: 1},
	{Facts: "The average person has over 1,460 dreams a year.", Lang: "en", Id: 2},
	{Facts: "The average person laughs 13 times a day.", Lang: "en", Id: 3},
	{Facts: "The average person will spend 6 months of their life waiting on a red light to turn green.", Lang: "en", Id: 4},
}

func main() {
	routeur := mux.NewRouter()
	routeur.HandleFunc("/facts", returnAll).Methods(http.MethodGet)
	routeur.HandleFunc("/facts/{id}", returnOne).Methods(http.MethodGet)
	//routeur.HandleFunc("/facts", create).Methods(http.MethodPost)
	//routeur.HandleFunc("/facts/{id}", update).Methods(http.MethodPut)
	//routeur.HandleFunc("/facts/{id}", delete).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", routeur)
}

func returnAll(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facts)
}

func returnOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	json.NewEncoder(w).Encode(facts[id])
}
