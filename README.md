# Gator in Golang
This is a command-line tool for aggregating and managing RSS feeds, built with Golang. It uses SQLC for database interaction, Goose for migrations, and Postgres as the database.

## Features
- Aggregate RSS feeds from various sources
- Stores feeds in a Postgres Database
- Use command-line interface to manage your RSS feeds

## Prerequisites
Before running this tool, ensure that the following dependencies are installed:

- Go (version 1.20 or higher)
- PostgreSQL
- Goose (for migrations)
- SQLC (for generating database access layer)

## Installing Dependencies
1. *Go*: Download and install Go from the official site or through your package manager.
```bash
sudo apt install golang  # Linux
brew install go          # macOS
```
2. *PostgreSQL*: Install Postgres:
```bash
sudo apt install postgresql  # Linux
brew install postgresql      # macOS
```
3. *Goose*: Install Goose for database migrations:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
4. *SQLC*: Install SQLC for code generation:
```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

## Set up config file
*.gatorconfig.json*: Create `.gatorconfig.json` file in your machine's home directory. The program will use it to track current user.
```json
{"db_url":"","current_user_name":""}
```

## Installation
Clone the repository and navigate into the project directory:
```bash
git clone https://github.com/Xavier-Hsiao/rss-aggregator.git
cd rss-aggregator
```

## Setting Up the Database
1. Create a new Postgres database:
```bash
createdb rss_aggregator
```

2. Apply database migrations using Goose:
```bash
goose postgres "postgres://user:password@localhost/rss_aggregator?sslmode=disable" up
```
Replace user and password with your PostgreSQL credentials.

3. Save the database URL to *gatorconfig.json* file.

## Running the Tool
```bash
go build -o rss-cli
./rss-cli
```

## Usage




