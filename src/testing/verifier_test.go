// Package test faz o teste dos arquivos principais.
package test

import (
	"testing"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/verifiers"
)

func TestVerifier(t *testing.T) {
	t.Run("Verificar", func(t *testing.T) {
		testVerifyLink(t)
	})
}

func testVerifyLink(t *testing.T) {
	testLinks := []struct {
		nome  string
		link  string
		passa bool
	}{
		{"Passa.com", "https://google.com", true},
		{"Passa.Net", "https://teste.net", true},
		{"FalhaSemHttps", "google.com", false},
		{"PassaLinkInterno", "https://localhost", true},
	}
	for _, tc := range testLinks {
		t.Run(tc.nome, func(t *testing.T) {
			if verifiers.ValidateURL(tc.link) != tc.passa {
				t.Errorf("Erro: resultado de %s diferente do esperado %t", tc.link, tc.passa)
			}
		})
	}
}
