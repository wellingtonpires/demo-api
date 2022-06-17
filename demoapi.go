package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
	http.HandleFunc("/demoapi", jsonMessage)
	fmt.Println("Running on :8080!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type DemoApi struct {
	serialNumber string `json:"serialNumber"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Category int `json:"category"`
}

func jsonMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	productList := `[{"serialNumber":"9020F4K30001SN","Name":"Teclado com fio Hightech USB Y920","Price":61.9,"Category":1},{"serialNumber":"3532F4K30120SN","Name":"Mouse Gamer Hightech H980","Price":179.9,"Category":1},{"serialNumber":"9765F4K30123SN","Name":"SSD StonKing K256 256GB","Price":149.9,"Category":2},{"serialNumber":"9324F4K35434SN","Name":"Memória GPX Trix RGB 8GB 3600MHz DDR4","Price":279.9,"Category":2},{"serialNumber":"2344F4K32342SN","Name":"Processador DMZ Zyzen 5 990X 4.1GHz","Price":1199.9,"Category":2},{"serialNumber":"1234F4K35432SN","Name":"HD XD Red 2TB 3.5","Price":299.9,"Category":2}]`

	w.Write([]byte(productList))
}
