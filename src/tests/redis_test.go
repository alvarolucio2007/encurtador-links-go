package maintest

import (
	"testing"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/cache"
	"github.com/alvarolucio2007/encurtador-links-go/src/internal/shortener"
)

func TestRedisFluxoCompleto(t *testing.T) {
	t.Run("Fluxo Completo", func(t *testing.T) {
		TestRedis(t)
	})
}

func TestRedis(t *testing.T) {
	cache.ConectarRedis()
	urlOriginal := "https://google.com"

	ID := uint64(1000)
	encodedURL := shortener.Encode(ID)
	err := cache.AdicionarLinkRedis(encodedURL, urlOriginal)
	if err != nil {
		t.Error("Erro no redis!", err)
	}
	urlRecuperada, err := cache.BuscarLinkRedis(encodedURL)
	if err != nil {
		t.Fatal("Erro ao buscar no Redis:", err)
	}
	if urlRecuperada != urlOriginal {
		t.Errorf("Erro! Esperava %s, recebi %s", urlOriginal, urlRecuperada)
	}
}
