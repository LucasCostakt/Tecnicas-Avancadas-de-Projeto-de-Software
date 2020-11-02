package readfile

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

//Recebe o path com o nome do arquivo e retorna o seu conteúdo
//caso houver algum erro o conteúdo é nulo
func ReadFile(filePath string) ([]byte, error) {

	//abre o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Arquivo não encontrado em ReadFile()")
		return nil, errors.New("Arquivo não encontrado")
	}
	//lê o conteúdo arquivo
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Erro ao ler o arquivo em ReadFile()")
		return nil, errors.New("Erro ao ler o arquivo")
	}

	defer file.Close()

	return []byte(fileContents), nil
}
