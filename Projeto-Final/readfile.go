package main

import (
	"errors"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Arquivo não encontrado")
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Erro ao ler o arquivo")
	}

	defer file.Close()

	return []byte(fileContents), nil
}
