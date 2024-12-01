package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func enviroment() {
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
