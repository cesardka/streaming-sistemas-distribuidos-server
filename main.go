package main

import (
	"fmt"
	"log"
	"net/http"
)

// headersParaCORS permite que qualquer caminho (incluindo localhost...!!) possa requisitar nossos arquivos
func headersParaCORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func main() {
	// TODO: permitir definir caminhos e porta alternativos via linha de comando
	diretorio_de_arquivos := "files"
	PORTA_PADRAO := 8080

	// Adiciona caminho para disponibilizar os arquivos
	http.Handle("/", headersParaCORS(http.FileServer(http.Dir(diretorio_de_arquivos))))

	// Iniciando servidor
	fmt.Printf("Iniciando servidor na porta %v\n", PORTA_PADRAO)
	log.Printf("Disponibilizando '%s' na porta: %v\n", diretorio_de_arquivos, PORTA_PADRAO)

	// Escuta por chamados e loga caso aconteça algum erro fatal
	porta_disponibilizada := fmt.Sprintf(":%v", PORTA_PADRAO)
	log.Fatal(http.ListenAndServe(porta_disponibilizada, nil))
}