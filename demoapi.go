package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/demoapi", jsonMessage)
	fmt.Println("Running on :8000!")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type DemoApi struct {
	Title string `json:"title"`
	Message string `json:"Message"`
}

func jsonMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DemoApi{
		Title: "Random title",
		Message: "Random message",
	})
}
