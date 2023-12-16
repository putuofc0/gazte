package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "Selamat datang di situs web sederhana menggunakan Golang! Counter: %d", counter)
	mutex.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server berjalan pada port 8080...")
	http.ListenAndServe(":8080", nil)
}
