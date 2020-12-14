package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é o endereço de conexão com o MySQL
	StringConexaoBanco = ""

	// StringConexaoSemBanco é o endereço de conexão com o MySQL sem o banco
	StringConexaoSemBanco = ""

	// Porta onde a API vai rodar
	Porta = 0

	// SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar variaveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"),
	)

	StringConexaoSemBanco = fmt.Sprintf("%s:%s@/",
		os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
