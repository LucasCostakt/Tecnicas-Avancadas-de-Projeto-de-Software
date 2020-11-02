package readfile

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

//Teste feito para função ReadFile()
func TestRead(t *testing.T) {
	myStructTestResponse := []struct {
		name  string //Nome do teste
		send  string //Path do arquivo a ser lido
		want1 []byte //O que esperamos de retorno da função, conteúdo do arquivo. Retorna nulo se houver algum erro.
		want2 error  //O tipo de erro caso ele exista se não existir tem que ser nulo nulo
	}{
		{ //Esse caso de teste verifica se a função leu corretamente o conteúdo do arquivo
			name:  "Teste sucesso ao ler o arquivo",
			send:  "C:/Users/lucas/Desktop/arquivo-de-teste.txt",
			want1: []byte("Esse será só um teste para a disciplina de técnicas avançadas de projeto de software"),
			want2: nil,
		},
		{ //Esse caso de teste verifica se o erro retornado está correto
			name:  "Teste erro ao localizar o arquivo",
			send:  "./arquivo-errado",
			want1: nil,
			want2: errors.New("Arquivo não encontrado"),
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.send)
			AssertResponsebody(t, got, tt.want1)
			AssertResponsebody(t, err, tt.want2)
		})
	}
}

//Compara o que foi respondido pela função com o que é esperado
func AssertResponsebody(t *testing.T, got, expectedResponse interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, expectedResponse) {
		str1 := fmt.Sprintf("%v", got)
		str2 := fmt.Sprintf("%v", expectedResponse)
		t.Errorf("Resposta errada, got %q Resposta esperada %q\n", str1, str2)
	}
}
