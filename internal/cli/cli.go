package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	cb "github.com/atotto/clipboard"
	"github.com/erodrigufer/maguet/internal/openai"
	"github.com/spf13/cobra"
)

const maguet_version = "v0.3.0"

func DefineCommands(api openai.ChatGPTResponder) {
	// Variables to store user defined flags.
	var inputFile string
	var outputFile string
	var openaiModel float32
	var temperature float32
	var clipboard bool
	var pager bool

	rootCmd := &cobra.Command{
		Use:     "maguet",
		Version: maguet_version,
		Short:   "A simple CLI app for interacting with OpenAI's ChatGPT API",
		Long: `maguet is a command-line interface that enables users to interact with OpenAI's ChatGPT API. 
You can use this app to prompt the ChatGPT model for text/code generation, chat with it and save output to a file.`,
	}

	completeCmd := &cobra.Command{
		Use:     "complete",
		Aliases: []string{"comp", "c", "co"},
		Short:   "Prompt for text generation/completion",
		Long: `Use the complete command to generate text or chat with the OpenAI's ChatGPT API.
The generated text or chat messages can be printed to the console or saved to a file using the -o flag.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Join all the elements of the args slice into a single string,
			// in which the elements are joined using whitespaces.
			// The args slice contains the user input in the command-line after
			// all flags have been read.
			prompt := strings.Join(args, " ")
			if err := completeExecution(prompt, inputFile, outputFile, temperature, openaiModel, clipboard, pager, api); err != nil {
				return err
			}
			return nil
		},
	}

	// Define the flags.
	completeCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The path to the input file to read")
	completeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "The path to the output file")
	completeCmd.Flags().Float32VarP(&openaiModel, "model", "m", 3.5, "OpenAI model used to request a response from a prompt (e.g. '4.0' for GPT 4).")
	completeCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.3, "The temperature value for text generation (between 0 and 1)")
	completeCmd.Flags().BoolVarP(&clipboard, "clipboard", "c", false, "Copy output to system clipboard.")
	completeCmd.Flags().BoolVarP(&pager, "pager", "p", false, "Show response in a pager.")

	rootCmd.AddCommand(completeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func completeExecution(prompt, inputFile, outputFile string, temperature, openaiModel float32, clipboard, pager bool, api openai.ChatGPTResponder) error {
	if openaiModel != 4.0 && openaiModel != 3.5 {
		return fmt.Errorf("invalid ChatGPT model (invalid 'model' flag values)")
	}

	if inputFile != "" {
		fmt.Println("Using input file to aggregate prompt...")
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

	fmt.Printf("Sending completion request to ChatGPT API...\nUsing ChatGPT %.1f\n", openaiModel)
	resp, err := api.RequestCompletion(prompt, temperature, openaiModel)
	if err != nil {
		return fmt.Errorf("Request for completion failed: %w", err)
	}

	// If the outputFile flag has not been set, print the completion.
	if outputFile == "" {
		if pager {
			if err := openPager(resp); err != nil {
				return fmt.Errorf("error while displaying text in pager: %w", err)
			}
		} else { // Do not use a pager to display output.
			fmt.Printf("--------- BEGIN ChatGPT response---------\n\n%s\n\n--------- END response ---------\n", resp)
		}
	} else { // Store in output file.
		fmt.Printf("Writing ChatGPT response into %s file...\n", outputFile)
		if err := os.WriteFile(outputFile, []byte(resp), 0644); err != nil {
			return fmt.Errorf("error writing response to output file: %w", err)
		}
		if err := openPager(resp); err != nil {
			fmt.Printf("error while displaying text in pager: %v\n", err)
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

func openPager(text string) error {
	pagerCmd := exec.Command("glow")
	pagerCmd.Stdin = strings.NewReader(text)
	pagerCmd.Stdout = os.Stdout

	// Fork off a process and wait for it to terminate.
	err := pagerCmd.Run()
	if err != nil {
		return fmt.Errorf("error while forking process with pager: %w", err)
	}
	return nil
}
