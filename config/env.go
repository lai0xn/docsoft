package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	JWTSecret   string
	TemplateDIR string
	SERVER_PORT string
	DSN         string
)

func LoadENV() {
	godotenv.Load("./.env")
	JWTSecret = os.Getenv("JWTSecret")
	SERVER_PORT = os.Getenv("PORT")
	TemplateDIR = os.Getenv("TemplateDIR")
	DSN = os.Getenv("DSN")
}
