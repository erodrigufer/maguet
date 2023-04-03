package cli

import (
	"fmt"
	"os"
	"strings"

	cb "github.com/atotto/clipboard"
	"github.com/erodrigufer/maguet/internal/openai"
	"github.com/spf13/cobra"
)

func DefineCommands(api openai.ChatGPTResponder) {
	// Variables to store user defined flags.
	var inputFile string
	var outputFile string
	var temperature float32
	var clipboard bool

	rootCmd := &cobra.Command{
		Use:   "maguet",
		Short: "A simple CLI app for interacting with OpenAI's ChatGPT API",
		Long: `maguet is a command-line interface that enables users to interact with OpenAI's ChatGPT API. 
You can use this app to prompt the ChatGPT model for text/code generation, chat with it and save output to a file.`,
	}

	completeCmd := &cobra.Command{
		Use:   "complete",
		Short: "Prompt for text generation/completion",
		Long: `Use the complete command to generate text or chat with the OpenAI's ChatGPT API.
The generated text or chat messages can be printed to the console or saved to a file using the -o flag.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Join all the elements of the args slice into a single string,
			// in which the elements are joined using whitespaces.
			// The args slice contains the user input in the command-line after
			// all flags have been read.
			prompt := strings.Join(args, " ")
			if err := completeExecution(prompt, inputFile, outputFile, temperature, clipboard, api); err != nil {
				return err
			}
			return nil
		},
	}

	// Define the flags.
	completeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The path to the input file to read")
	completeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The path to the output file")
	completeCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.3, "The temperature value for text generation (between 0 and 1)")
	completeCmd.Flags().BoolVarP(&clipboard, "clipboard", "c", false, "Copy output to system clipboard.")

	rootCmd.AddCommand(completeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func completeExecution(prompt, inputFile, outputFile string, temperature float32, clipboard bool, api openai.ChatGPTResponder) error {
	if inputFile != "" {
		// Read the content of the file into a byte slice.
		fileData, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
		// Convert the byte slice to a string.
		content := string(fileData)

		// Modify prompt to add input file to it.
		prompt = fmt.Sprintf("%s:\n\"%s\"", prompt, content)
	}

	fmt.Printf("Sending completion request to ChatGPT API...\n")
	resp, err := api.RequestCompletion(prompt, temperature)
	if err != nil {
		return fmt.Errorf("Request for completion failed: %v", err)
	}

	// If the outputFile flag has not been set, print the completion.
	if outputFile == "" {
		fmt.Printf("--------- BEGIN ChatGPT response---------\n\n%s\n\n--------- END response ---------\n", resp)
	} else {
		fmt.Printf("Writing ChatGPT response into %s file...\n", outputFile)
		if err := os.WriteFile(outputFile, []byte(resp), 0644); err != nil {
			return fmt.Errorf("error writing response to output file: %w", err)
		}
	}

	if clipboard {
		// Copy response into system's clipboard.
		if err := cb.WriteAll(resp); err != nil {
			return fmt.Errorf("error copying response to system's clipboard: %w", err)
		}
		fmt.Printf("\nChatGPT response was successfully copied into the system's clipboard.\n")
	}

	return nil
}
