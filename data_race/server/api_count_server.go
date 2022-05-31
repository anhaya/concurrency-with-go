package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var httpCount int
var mu sync.Mutex

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/count", IncorrectCount).Methods("GET")
	router.HandleFunc("/v2/count", CorrectCountv2).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}

func IncorrectCount(w http.ResponseWriter, r *http.Request) {
	httpCount++ //in a concurrent cenario, we are going to have data race problem, consequently, incosistent information
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(httpCount)))
}

func CorrectCountv2(w http.ResponseWriter, r *http.Request) {
	mu.Lock() //this way we are sayind that no one can access this memory address until we unlock. Mutex says that only
	//one goroutine can access by time. If another goroutine comes and try to lock, it wont throw deadlock
	//because Mutex will wait until an unlock. This happens because each request is a goroutine, but If I put mu.lock()
	//twice at the same goroutine I would receive deadlock
	defer mu.Unlock()
	httpCount++
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(httpCount)))
}
