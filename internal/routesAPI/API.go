package routes

import (
	"fmt"
	"log"
	"net/http"
)

func SetupAPI() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", HandleShorten)
	mux.HandleFunc("GET /{code}", HandleRedirect)
	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Erro ao subir o servidor: ", err)
	}
}
