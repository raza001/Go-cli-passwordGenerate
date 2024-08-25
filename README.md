# Password Generator CLI

This project provides a simple CLI tool to generate random passwords with customizable options using [Cobra](https://github.com/spf13/cobra).

## Features

- Generate random passwords of a specified length.
- Include digits and special characters in the password.

## Installation

To use this CLI tool, you need to have Go installed on your system. Follow the steps below to get started:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/password-generator-cli.git
   cd password-generator-cli
   ```

2. Build the CLI tool:

   ```bash
   go build x
   ```

3. Run the tool:

   ```bash
   ./generate
   ```

## Usage

This CLI tool allows you to generate random passwords with different options.

### Basic Usage

```bash
./generate -l 12 -d -s
```

- `-l` or `--length`: Specifies the length of the password.
- `-d` or `--digits`: Includes digits in the password.
- `-s` or `--special-chars`: Includes special characters in the password.

### Examples

1. Generate a 12-character password with digits and special characters:

   ```bash
   ./generate -l 12 -d -s
   ```

2. Generate a 16-character password with only letters:

   ```bash
   ./generate -l 16
   ```

## Commands

- `generate`: Generates a random password based on the provided options.

### Command Options

- `-l`, `--length int`: Specifies the length of the password (default is 12).
- `-d`, `--digits`: Include digits in the password.
- `-s`, `--special-chars`: Include special characters in the password.
