// Package routes elabora e faz as rotas do encurtamento dos links.
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/cache"
	"github.com/alvarolucio2007/encurtador-links-go/src/internal/database"
	"github.com/alvarolucio2007/encurtador-links-go/src/internal/verifiers"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	url, err := cache.BuscarLinkRedis(code)
	if err == nil && url != "" {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	url, err = database.BuscarURLOriginal(code)
	if err != nil {
		http.Error(w, "link não encontrado", http.StatusNotFound)
		return
	}
	err = cache.AdicionarLinkRedis(code, url)
	if err != nil {
		http.Error(w, "Erro ao adicionar o código no Redis!", http.StatusInternalServerError)
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var input struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	if !verifiers.ValidateURL(input.URL) {
		http.Error(w, "URL inválida", http.StatusBadRequest)
		return
	}
	code, err := database.CriarEntradaPostgres(input.URL)
	if err != nil {
		http.Error(w, "Erro ao adicionar o código no postgres", http.StatusInternalServerError)
	}
	if err = cache.AdicionarLinkRedis(code, input.URL); err != nil {
		http.Error(w, "Erro ao adicionar o código no redis", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(map[string]string{"code": code}); err != nil {
		http.Error(w, "Erro no encoder do JSON!", http.StatusInternalServerError)
	}
}
