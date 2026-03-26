// Package database monitora e ajusta o Postgresql
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/alvarolucio2007/encurtador-links-go/internal/shortener"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarDataBase() {
	var err error
	dsn := "host=pg_encurtador user=user password=password dbname=shortener sslmode=disable"
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Não foi possível se conectar ao postgres:", err)
	}
	fmt.Println("Postgres conectado.")
}

func MigrarBanco() {
	query := `
    CREATE TABLE IF NOT EXISTS links (
        id BIGSERIAL PRIMARY KEY,
        url_original TEXT NOT NULL,
        criado_em TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );
    CREATE INDEX IF NOT EXISTS idx_url_original ON links(url_original);
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Erro na migração do banco: ", err)
	}
	fmt.Println("✅ Tabelas prontas para uso.")
}

func CriarEntradaPostgres(url string) (string, error) {
	var id uint64
	query := "INSERT INTO links (url_original) VALUES ($1) RETURNING id"
	err := DB.QueryRow(query, url).Scan(&id)
	if err != nil {
		return "", errors.New("erro ao criar a entrada no postgres")
	}
	codigoCurto := shortener.Encode(id)
	return codigoCurto, nil
}

func BuscarURLOriginal(codigoCurto string) (string, error) {
	id := shortener.Decode(codigoCurto)
	var urlOriginal string
	query := "SELECT url_original FROM links WHERE id=$1"
	err := DB.QueryRow(query, id).Scan(&urlOriginal)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("link não encontrado")
		}
		return "", err
	}
	return urlOriginal, nil
}
