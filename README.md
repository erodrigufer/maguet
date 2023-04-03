# maguet

**maguet** is a Go script that allows you to communicate with ChatGPT to request text completions.
The name "maguet" comes from the Catalan language and means "small wizard" (it can also be translated to the Spanish word "_maguito_").

## Requirements

To use maguet, you must have Go installed on your system.
You will also need an OpenAI API key to access the ChatGPT API.

## Installation

To install maguet, you can clone the repository from GitHub and build the executable using the following commands:

```
git clone https://github.com/erodrigufer/maguet.git
cd maguet
make build
```

This will create an executable file named `maguet` in the current directory.

Alternatively, you can install maguet using the `go install` command:

```
go install github.com/erodrigufer/maguet
```

After installing maguet, you can verify that it is working properly by running the following command:

```
maguet
```

## Usage

To use maguet, you must first set your OpenAI API key as an environment variable:

```
export MAGUET_TOKEN=<your-api-key>
```

Once you have set your API key, you can use maguet to request text completions from ChatGPT. For example:

```
maguet complete "Complete this sentence: The quick brown fox"
```

This will send a request to ChatGPT to complete the sentence "The quick brown fox" and return the completed text.

For more information execute the help menu:

```
maguet complete -h
```

## License

maguet is licensed under the MIT License. See the LICENSE file for more information.
