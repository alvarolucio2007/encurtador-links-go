package maintest

import (
	"fmt"
	"os"
	"testing"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/cache"
	"github.com/alvarolucio2007/encurtador-links-go/src/internal/database"
)

func TestMain(m *testing.M) {
	fmt.Println("Preparando testes")
	database.ConectarDataBase()
	cache.ConectarRedis()
	_, err := database.DB.Exec("TRUNCATE TABLE links RESTART IDENTITY")
	if err != nil {
		fmt.Printf("Erro ao limpar DB: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Banco pronto para testes")
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestFluxoCompletoApp(t *testing.T) {
	t.Run("Postgres", func(t *testing.T) {
		TestPostgres(t)
	})
	t.Run("Redis", func(t *testing.T) {
		TestRedisFluxoCompleto(t)
	})
	t.Run("Verifier", func(t *testing.T) {
		TestVerifier(t)
	})
	t.Run("Shortener", func(t *testing.T) {
		TestEncodeDecode(t)
	})
}
