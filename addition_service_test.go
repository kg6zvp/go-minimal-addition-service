package main

import (
	"net/http"
	"testing"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

func setup() *http.Server {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/add", additionHandler)
	go func() {
		server.ListenAndServe()
	}()
	return server
}

func TestAddition(t *testing.T) {
	srv := setup()
	resp, _ := http.Post("http://localhost:8080/add", "application/json", bytes.NewBuffer([]byte(`{"a": 3,"b": 7}`)))
	body, _ := ioutil.ReadAll(resp.Body)
	var b map[string]int
	json.Unmarshal(body, &b)
	srv.Shutdown(nil)
	if b["sum"] == 10 {
		return
	}
	panic(nil)
}
