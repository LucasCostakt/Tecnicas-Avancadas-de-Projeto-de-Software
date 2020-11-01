package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	myStructTestResponse := []struct {
		name  string
		send  string
		want1 []byte
		want2 error
	}{
		{name: "Teste sucesso ao ler o arquivo", send: "C:/Users/lucas/Desktop/arquivo-de-teste", want1: []byte("Esse será só um teste para a disciplina de técnicas avançadas de projeto de software"), want2: nil},
		{name: "Teste erro ao localizar o arquivo", send: "./arquivo-de-teste", want1: nil, want2: errors.New("Arquivo não encontrado")},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.send)
			AssertResponsebody(t, got, tt.want1)
			AssertResponsebody(t, err, tt.want2)
		})
	}
}

func AssertResponsebody(t *testing.T, got, expectedResponse interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expectedResponse) {
		str1 := fmt.Sprintf("%v", got)
		str2 := fmt.Sprintf("%v", expectedResponse)
		t.Errorf("Resposta errada, got %q Resposta esperada %q\n", str1, str2)
	}
}
