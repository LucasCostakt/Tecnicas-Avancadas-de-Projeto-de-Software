package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Imprime o conteúdo do arquivo na página web
func handlerHello(w http.ResponseWriter, r *http.Request) {
	nameFile := "hellomundo.html"
	responseArchive := readArchive(nameFile)
	fmt.Fprintf(w, string(responseArchive))

}

//Recebe o nome do arquivo com o path e retorna o conteúdo contido nele
func readArchive(nameFile string) []byte {
	file, err := ioutil.ReadFile(nameFile)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

//Inicia o servidor no endereço localhost:5000
func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(handlerHello))
	if err != nil {
		log.Println(err)
	}
}
