package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Representa a URl para comunicacao com a API
	ApiUrl = ""
	// Porta onde a aplicacao WEB esta rodando
	Porta = 0
	// HashKey é utilizada pra autenticar o cookie
	HashKey []byte
	// BlockKey é utilizado para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar iniciliza as variaveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
