package maintest

import (
	"testing"

	"github.com/alvarolucio2007/encurtador-links-go/src/internal/shortener"
)

func TestEncodeDecode(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		testarEncoding(t)
	})
	t.Run("Decode", func(t *testing.T) {
		testarDecoding(t)
	})
}

func testarEncoding(t *testing.T) {
	testIDs := []struct {
		nome    string
		entrada uint64
		saida   string
	}{
		{"ID Mil", 1000, "g8"},
		{"ID Milhão", 10000000, "FXsk"},
		{"ID Grande", 3810923821908, "155NdZ6A"},
	}
	for _, tc := range testIDs {
		t.Run(tc.nome, func(t *testing.T) {
			resultado := shortener.Encode(tc.entrada)
			if resultado != tc.saida {
				t.Errorf("Falha no teste [%s]: entrada %d gerou %s, mas queríamos %s",
					tc.nome, tc.entrada, resultado, tc.saida)
			}
		})
	}
}

func testarDecoding(t *testing.T) {
	testIDs := []struct {
		nome    string
		entrada string
		saida   uint64
	}{
		{"Nome pequeno", "g8", 1000},
		{"Nome grande", "FXsk", 10000000},
		{"Nome gigante", "155NdZ6A", 3810923821908},
	}
	for _, tc := range testIDs {
		t.Run(tc.nome, func(t *testing.T) {
			resultado := shortener.Decode(tc.entrada)
			if resultado != tc.saida {
				t.Errorf("Falha no teste [%s]: entrada %s gerou %v, mas queríamos %v",
					tc.nome, tc.entrada, resultado, tc.saida)
			}
		})
	}
}
