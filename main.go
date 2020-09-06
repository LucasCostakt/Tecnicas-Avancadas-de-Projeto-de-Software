package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	nomeFile := "telainicio.html"
	responseArquivo := lerArquivo(nomeFile)
	fmt.Fprintf(w, string(responseArquivo))

}

func lerArquivo(nomeFile string) []byte {
	conteudoFile, err := ioutil.ReadFile(nomeFile)
	if err != nil {
		fmt.Println(err)
	}
	return conteudoFile
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerMeuCumprimento))
	if err != nil {
		log.Println(err)
	}
}
