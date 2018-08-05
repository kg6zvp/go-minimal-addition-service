package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type pair struct {
	A int `json:"a"`
	B int `json:"b"`
}

func additionHandler(w http.ResponseWriter, re *http.Request) {
	var req pair
	json.NewDecoder(re.Body).Decode(&req)
	json.NewEncoder(w).Encode(map[string]int{"sum": req.A + req.B})
}

func main() {
	http.HandleFunc("/add", additionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
