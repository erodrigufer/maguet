package main

import (
	"fmt"
	"os"

	"github.com/erodrigufer/maguet/internal/app"
	"github.com/erodrigufer/maguet/internal/cli"
	"github.com/erodrigufer/maguet/internal/openai"
)

func main() {
	authtoken, err := app.GetAuthToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error retrieving auth token: %v\n", err)
		os.Exit(1)
	}

	client := openai.NewClient(authtoken)

	// Mocks.
	// client := openai.NewMockClient()

	cli.DefineCommands(client)
}
