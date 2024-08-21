package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Utilizando o writer do FMT
func FmtWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Utilizando o writter do fmt!")
}

// Exemplos de JSON
func JsonEncode(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Resposta por encoder!")
}

func JsonDirectlyWrited(w http.ResponseWriter, r *http.Request) {
	mensagem, _ := json.Marshal("Resposta escrita diretamente!")
	w.Write(mensagem)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Called!")
}
