# Simrab Translate

A command-line tool for managing translations using Google Cloud Translate API. This tool helps you manage and generate translations for different languages in your project.

## Prerequisites

- Go 1.20 or later
- Google Cloud Platform account with Cloud Translate API enabled
- Google Cloud credentials properly configured

## Installation

1. Clone the repository:
```bash
git clone https://bitbucket.org/simrab/simrab-translate.git
cd simrab-translate
```

2. Install dependencies:
```bash
go mod download
```

3. Build the binary:
```bash
go build -o simrab-translate
```

## Configuration

Before using the tool, make sure you have:

### If you want to translate the missing keys:

1. Set up a Google Cloud Project 
2. Enabled the Cloud Translate API
3. Set up authentication by either:
   - Setting the `GOOGLE_APPLICATION_CREDENTIALS` environment variable pointing to your service account key file
   - Using gcloud authentication if running locally

## Usage

The tool uses a command-line interface built with Cobra. Here are some basic commands:

```bash
# Run the application
./simrab-translate [command]
```

For detailed usage instructions and available commands, run:
```bash
./simrab-translate --help
```

2. Clone the repository:
https://github.com/simrab/simrab-translate/blob/main/cmd/root.go
