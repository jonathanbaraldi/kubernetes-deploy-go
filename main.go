package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Test of automatic deployment on Kubernetes integrating the Github on Rancher :) \nv0.1"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
