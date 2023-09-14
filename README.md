# maguet

**maguet** is a Go script that allows you to communicate with ChatGPT to request text completions.
The name "maguet" comes from the Catalan language and means "small wizard" (it can also be translated to the Spanish word "_maguito_").

## Requirements

1. Go v1.21
2. An OpenAI API key to access the ChatGPT API.
3. [glow](https://github.com/charmbracelet/glow) (_used as the default pager_)

## Installation

To install maguet, you can clone the repository from GitHub and build the executable using the following commands:

```bash
git clone https://github.com/erodrigufer/maguet.git
cd maguet
make install
```

This will create an executable file named `maguet` in the current directory and copy it to `~/bin`

Alternatively, you can install maguet using the `go install` command:

```bash
go install github.com/erodrigufer/maguet/cmd/maguet@latest
```

After installing maguet, you can verify that it is working properly by running the following command:

```bash
maguet --version
```

## Usage

To use maguet, you must first set your OpenAI API key as an environment variable by storing your OpenAI API key at `~/.maguet.env` as `MAGUET_TOKEN=<your-api-key>`.

Once you have set your API key, you can use maguet to request text completions from ChatGPT. For example:

```bash
maguet complete "Complete this sentence: The quick brown fox"
```

This will send a request to ChatGPT to complete the sentence "The quick brown fox" and return the completed text.

For more information execute the help menu:

```bash
maguet complete -h
```

## License

maguet is licensed under the MIT License. See the LICENSE file for more information.
