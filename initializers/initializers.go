package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	shell := os.Getenv("SHELL")
	fmt.Println(shell)
}
