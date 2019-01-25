package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, name)
}


func main() {
	http.HandleFunc("/", handler)
	log.Fatal (http.ListenAndServe("0.0.0.0:8080", nil))
}
