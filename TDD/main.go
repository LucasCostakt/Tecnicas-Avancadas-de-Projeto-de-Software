package TDD

import (
	"encoding/json"
	"log"
)

func returnJson() []byte {
	mystruct := struct {
		Name string `json:"name"`
		Code string `json:"code"`
		Body string `json:"body"`
	}{
		Name: "Lucas",
		Code: "Success",
		Body: "Hello",
	}
	myjson, err := json.Marshal(mystruct)
	if err != nil {
		log.Println("Json Marshal error")
	}
	return myjson
}
