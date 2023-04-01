package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func DefineCommands() {
	var filePath string
	var outputFile string
	var temperature float64

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
		// TODO: handle errors, for example when API connection fails.
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("maguet prompt!\nFile path: %s\nOutput file: %s\nTemperature: %.2f\n", filePath, outputFile, temperature)
		},
	}

	// Define the flags.
	promptCmd.Flags().StringVarP(&filePath, "file", "f", "", "The path to the file to read")
	promptCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The path to the output file")
	promptCmd.Flags().Float64VarP(&temperature, "temperature", "t", 0.7, "The temperature value for text generation (between 0 and 1)")

	rootCmd.AddCommand(promptCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
