package main

import (
	"flag"
	"fmt"
	"projeto-final/services/readfile"
)

func main() {
	//Flag coleta o path do arquivo passado junto ao comando do console
	flag.Parse()
	flagArg := flag.Arg(0)
	//Chama função de leitura de arquivo
	fileContentes, _ := readfile.ReadFile(flagArg)
	//Imprime no console o conteúdo de dentro do arquivo
	fmt.Println(string(fileContentes))
}
