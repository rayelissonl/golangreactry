package main

import (
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar tern migrate: %v\nSaída: %s", err, output)
	}
}
