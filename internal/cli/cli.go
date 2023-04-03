package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/erodrigufer/maguet/internal/openai"
	"github.com/spf13/cobra"
)

func DefineCommands(api openai.ChatGPTResponder) {
	// Variables to store user defined flags.
	var filePath string
	var outputFile string
	var temperature float32

	rootCmd := &cobra.Command{
		Use:   "maguet",
		Short: "A simple CLI app for interacting with OpenAI's ChatGPT API",
		Long: `maguet is a command-line interface that enables users to interact with OpenAI's ChatGPT API. 
You can use this app to prompt the ChatGPT model for text/code generation, chat with it and save output to a file.`,
	}

	promptCmd := &cobra.Command{
		Use:   "prompt",
		Short: "Prompt for text generation or chat",
		Long: `Use the prompt command to generate text or chat with the OpenAI's ChatGPT API.
The generated text or chat messages can be printed to the console or saved to a file using the -o flag.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Join all the elements of the args slice into a single string,
			// in which the elements are joined using whitespaces.
			// The args slice contains the user input in the command-line after
			// all flags have been read.
			prompt := strings.Join(args, " ")

			// fmt.Printf("maguet prompt subcommand!\nFile path: %s\nOutput file: %s\nTemperature: %.2f\nprompt: %s\n", filePath, outputFile, temperature, prompt)
			fmt.Printf("Sending completion request to ChatGPT API...\n")
			resp, err := api.RequestCompletion(prompt, temperature)
			if err != nil {
				return fmt.Errorf("Request for completion failed: %v", err)
			}

			// If the outputFile flag has not been set, print the completion.
			if outputFile == "" {
				fmt.Printf("--------- BEGIN ChatGPT response---------\n\n%s\n\n--------- END response ---------\n", resp)
			}
			return nil
		},
	}

	// Define the flags.
	promptCmd.Flags().StringVarP(&filePath, "file", "f", "", "The path to the file to read")
	promptCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The path to the output file")
	promptCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.3, "The temperature value for text generation (between 0 and 1)")

	rootCmd.AddCommand(promptCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
