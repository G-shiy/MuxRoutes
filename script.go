package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [up|down|reload]")
		os.Exit(1)
	}

	action := os.Args[1]

	switch action {
	case "up":
		runCommand("docker-compose", "up", "-d")
	case "down":
		runCommand("docker-compose", "down")
	case "reload":
		runCommand("docker-compose", "down")
		runCommand("docker-compose", "up", "-d")
	default:
		fmt.Println("Comando inválido. Use: up, down ou reload")
	}
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o comando %s: %v\n", command, err)
		os.Exit(1)
	}
}
