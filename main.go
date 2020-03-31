package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9009", router))
}
