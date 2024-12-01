package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	db := configDb()
	fmt.Println(db)

	col, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal("erro ao executar a query")
	}
	fmt.Println(col)

	defer db.Close()

}

func configEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envFilePath := filepath.Join(cwd, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal(err)
	}

}

func configDb() *sql.DB {
	configEnv()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	fmt.Printf("Conectando ao banco de dados com: %s\n", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao testar conexão com o banco: %v", err)
	}
	fmt.Println("Conexão estabelecida com sucesso!")

	return db
}
