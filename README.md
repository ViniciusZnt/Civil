# Temperature and Humidity Monitoring Web Project

This project was created to monitor temperature and humidity using ESP32, Golang, Templ, and Tailwind.

## Installation

### Tailwind CSS
Follow the instructions on [Tailwind CSS Standalone CLI](https://tailwindcss.com/blog/standalone-cli) to install Tailwind CSS.

### Templ
Follow the instructions on [Templ Quick Start](https://templ.guide/quick-start/installation/) for installation.

### Golang
Follow the instructions on [Go Installation Guide](https://go.dev/doc/install) to install Golang.

### Visual Studio Code Extensions
Install the following extensions for Visual Studio Code:
- Templ-vscode
- Go
- Tailwind CSS IntelliSense

## Usage

I have created an alias for ease of use:

alias buildapp='templ generate && go build -o ./tmp/main . && go run main.go'

# Run the alias
buildapp

After running the alias, your application should be built and running.

- Note: Make sure to replace the paths and commands according to your project structure.
