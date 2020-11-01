package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	letter, _ := ioutil.ReadFile("hellomundo.html")
	fmt.Fprintf(w, string(letter))
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerMeuCumprimento))
	if err != nil {
		log.Println(err)
	}
}
