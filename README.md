# ChainServer

A HTTP service wrapper for ChainMaker SDK, designed to work with Casibase without requiring CGO support.

## Overview

ChainServer is a lightweight HTTP service that provides a bridge between Casibase and ChainMaker blockchain. Instead of directly integrating ChainMaker SDK (which requires CGO), this service offers HTTP endpoints to interact with ChainMaker blockchain, making it more flexible and easier to deploy.

## Features

- RESTful API interface for ChainMaker operations
- Easy integration with Casibase
- No CGO dependency required for client applications

## Compatibility Notice

**Important:**  
Currently, ChainServer only supports ChainMaker blockchains with **TLS disabled** and **authtype set to `permissionedWithCert`**. If you encounter errors when invoking transaction or query APIs, please verify that your ChainMaker server configuration meets these requirements.  
Additionally, ensure that the `provider` configuration in your Casibase setup is consistent with your ChainMaker network settings.

For details on how to check and modify your ChainMaker chain configuration, please refer to the official documentation:  
- [ChainMaker Configuration Guide](https://docs.chainmaker.org.cn/manage/%E9%95%BF%E5%AE%89%E9%93%BE%E9%85%8D%E7%BD%AE%E7%AE%A1%E7%90%86.html)

If your configuration does not meet the above requirements, the service may not function as expected.

## Technical Stack

- Go 1.22+
- Beego Framework
- ChainMaker SDK v2.2.0

## Configuration

The server runs on port 13900 by default. You can modify the port in the configuration file.

## API Endpoints

The service provides HTTP endpoints for ChainMaker operations. Detailed API documentation will be provided separately.

## Building and Running

1. Install Go 1.22 or later
2. Clone the repository
```bash
git clone https://github.com/casibase/chainserver.git
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