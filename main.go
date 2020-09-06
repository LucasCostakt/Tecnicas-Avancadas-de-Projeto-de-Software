package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerMeuCumprimento))
	if err != nil {
		log.Println(err)
	}
}
