package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Database() *sql.DB {
	enviroment()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao testar conexão com o banco: %v", err)
	}
	return db
}
