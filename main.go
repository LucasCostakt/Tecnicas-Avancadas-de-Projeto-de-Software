package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Imprime o conteúdo do arquivo na página web
func handlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	nomeFile := "telainicio.html"
	responseArquivo := lerArquivo(nomeFile)
	fmt.Fprintf(w, string(responseArquivo))

}

//Recebe o nome do arquivo com o path e retorna o conteúdo contido nele
func lerArquivo(nomeFile string) []byte {
	conteudoFile, err := ioutil.ReadFile(nomeFile)
	if err != nil {
		fmt.Println(err)
	}
	return conteudoFile
}

//Inicia o servidor no endereço localhost:5000
func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerMeuCumprimento))
	if err != nil {
		log.Println(err)
	}
}
