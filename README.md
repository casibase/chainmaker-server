<h1 align="center" style="border-bottom: none;">ChainServer</h1>
<p align="center">
  <a href="#badge">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="https://hub.docker.com/r/casbin/chainserver">
    <img alt="docker pull casbin/chainserver" src="https://img.shields.io/docker/pulls/casbin/chainserver.svg">
  </a>
  <a href="https://github.com/casibase/chainserver/actions/workflows/build.yml">
    <img alt="GitHub Workflow Status (branch)" src="https://github.com/casibase/chainserver/workflows/Build/badge.svg?style=flat-square">
  </a>
  <a href="https://github.com/casibase/chainserver/releases/latest">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/casibase/chainserver.svg">
  </a>
  <a href="https://hub.docker.com/r/casbin/chainserver">
    <img alt="Docker Image Version (latest semver)" src="https://img.shields.io/badge/Docker%20Hub-latest-brightgreen">
  </a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/casibase/chainserver">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/casibase/chainserver?style=flat-square">
  </a>
  <a href="https://github.com/casibase/chainserver/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/casibase/chainserver?style=flat-square" alt="license">
  </a>
  <a href="https://github.com/casibase/chainserver/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/casibase/chainserver?style=flat-square">
  </a>
  <a href="#">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/casibase/chainserver?style=flat-square">
  </a>
  <a href="https://github.com/casibase/casibase/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/casibase/chainserver?style=flat-square">
  </a>
  <a href="https://discord.gg/devUNrWXrh">
    <img alt="Discord" src="https://img.shields.io/discord/1022748306096537660?logo=discord&label=discord&color=5865F2">
  </a>
</p>

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

## Docker Deployment Guide

The recommended way to deploy is to use the official image from Docker Hub:

```sh
docker run -d --name chainserver -p 13900:13900 casbin/chainserver:latest
```

Alternatively, you can use Docker Compose (recommended for production or multi-container setups):

```sh
docker-compose up -d
```

You can also build the image locally if you want to use your own changes:

```sh
docker build -t chainserver .
docker run -d --name chainserver -p 13900:13900 chainserver
```

The default service port is 13900. For troubleshooting or advanced usage, see the Dockerfile and build.sh, or open an issue if you encounter problems.

## API Documentation (Swagger)

After starting the service, you can view the API documentation via Swagger UI:

- Open your browser and visit: [http://localhost:13900/swagger/index.html](http://localhost:13900/swagger/index.html)

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