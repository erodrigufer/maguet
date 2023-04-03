package main

import (
	"fmt"
	"os"

	"github.com/erodrigufer/maguet/internal/cli"
	"github.com/erodrigufer/maguet/internal/openai"
	"github.com/joho/godotenv"
)

func main() {
	// TODO: relocate these commands ?
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading .env file")
	}

	authtoken := "mock_auth_token"
	client := openai.NewClient(authtoken)

	// Mocks.
	// client := openai.NewMockClient()

	cli.DefineCommands(client)
}
