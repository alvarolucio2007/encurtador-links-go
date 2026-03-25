package maintest

import (
	"testing"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/database"
)

func TestPostgres(t *testing.T) {
	t.Run("Fluxo completo", func(t *testing.T) {
		testarPostgres(t)
	})
}

func testarPostgres(t *testing.T) {
	database.ConectarDataBase()
	database.MigrarBanco()

	urlEsperada := "https://github.com/alvarolucio2007"
	var codigo string

	t.Run("Escrita: Deve salvar a URL e retornar um código menor", func(t *testing.T) {
		var err error
		codigo, err = database.CriarEntradaPostgres(urlEsperada)
		if err != nil {
			t.Errorf("Erro interno! %v", err)
		}
		if codigo == "" {
			t.Error("Esperava um código, veio vazio.")
		}
	})
	t.Run("Leitura: Deve recuperar a URL original usando o código", func(t *testing.T) {
		urlRecuperada, err := database.BuscarURLOriginal(codigo)
		if err != nil {
			t.Errorf("Erro ao buscar URL: %v", err)
		}
		if urlEsperada != urlRecuperada {
			t.Errorf("Diferença detectada, esperava %s e recebi %s", urlEsperada, urlRecuperada)
		}
	})
}
