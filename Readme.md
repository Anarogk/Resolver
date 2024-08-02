


# DNS Resolver

A robust DNS Resolver written in Golang, featuring caching with `go-cache` and structured logging with `zap`. The project includes a command-line interface (CLI) built using the Cobra library, allowing easy management and execution of the DNS resolver.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Usage](#usage)
- [Disclaimer](#disclaimer)
- [License](#license)

## Description

The DNS Resolver is designed to process DNS queries and responses efficiently. It integrates caching to handle frequently queried records and a robust logging system to provide detailed logs of operations. The CLI built with Cobra makes it easy to manage and run the DNS Resolver.

## Features

- Efficient DNS query processing
- Caching for frequently queried records using `go-cache`
- Structured logging with `zap`
- Log rotation using `lumberjack`
- Command-line interface with Cobra

## Project Structure

```
dns_resolver/
├── cmd/
│   ├── root.go
│   ├── start.go
│   └── main.go
├── internal/
│   ├── cache/
│   │   └── cache.go
│   ├── dns/
│   │   ├── resolver.go
│   │   ├── resolver_test.go
│   └── logger/
│       └── logger.go
├── scripts/
│   └── build.sh
└── go.mod
```

## Dependencies

- [go-cache](https://github.com/patrickmn/go-cache): In-memory key:value store/cache
- [zap](https://github.com/uber-go/zap): Blazing fast, structured, leveled logging in Go
- [lumberjack](https://github.com/natefinch/lumberjack): Rolling logger
- [Cobra](https://github.com/spf13/cobra): A library for creating powerful modern CLI applications

## Installation

### Prerequisites

- Go 1.16 or higher

### Clone the repository

```sh
git clone https://github.com/yourusername/dns_resolver.git
cd dns_resolver
```

### Install dependencies

```sh
go mod tidy
```

## Usage

### Build the application

Make the build script executable and run it:

```sh
chmod +x scripts/build.sh
./scripts/build.sh
```

### Run the DNS Resolver

Use the generated binary to start the DNS Resolver:

```sh
./dns_resolver start --address ":53" --loglevel "info"
```

### Command-Line Options

- `--address` (`-a`): Address to listen for DNS queries (default: `:53`)
- `--loglevel` (`-l`): Logging level (default: `info`)

## Disclaimer

This project is developed by me alone and in a very non-professional capacity. It is intended for educational and testing purposes only. Use it at your own risk. Ensure that you comply with local laws and regulations when deploying and using DNS servers.

<!-- ## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->

