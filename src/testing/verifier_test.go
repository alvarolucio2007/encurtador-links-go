// Package test faz o teste dos arquivos principais.
package test

import "testing"

func testVerifyLink(t *testing.T) {
	testLinks := []struct {
		nome  string
		link  string
		passa bool
	}{
		{"Passa.com", "https://google.com", true},
		{"Passa.Net", "https://teste.net", true},
		{"FalhaSemHttps", "google.com", true},
		{"FalhaSem.", "https://google", true},
	}
	for _, tc := range testLinks {
	}
}
