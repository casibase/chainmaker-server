# ChainMaker Server

A HTTP service wrapper for ChainMaker SDK, designed to work with Casibase without requiring CGO support.

## Overview

ChainMaker Server is a lightweight HTTP service that provides a bridge between Casibase and ChainMaker blockchain. Instead of directly integrating ChainMaker SDK (which requires CGO), this service offers HTTP endpoints to interact with ChainMaker blockchain, making it more flexible and easier to deploy.

## Features

- RESTful API interface for ChainMaker operations
- Easy integration with Casibase
- No CGO dependency required for client applications

## Technical Stack

- Go 1.22+
- Beego Framework
- ChainMaker SDK v2.2.0

## Configuration

The server runs on port 14000 by default. You can modify the port in the configuration file.

## API Endpoints

The service provides HTTP endpoints for ChainMaker operations. Detailed API documentation will be provided separately.

## Building and Running

1. Install Go 1.22 or later
2. Clone the repository
```bash
git clone https://github.com/casibase/chainmaker-server.git
```
3. Install dependencies
```bash
go mod tidy
```
4. Run the server
```bash
go run main.go
```

## License

Apache-2.0

## Related Projects

- [Casibase](https://github.com/casibase/casibase)
- [ChainMaker](https://git.chainmaker.org.cn/chainmaker)