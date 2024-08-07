# Overview
kvdb is an in-memory key-value store written in the Go programming language.

It uses a simple client-server model and is designed to be easy to use. It can be used as a database, session storage, or cache. It may not be suitable for complex needs.

Data is stored at keys of different types. Each data type allows you to store different kind of data such as string values or objects.

Instances are easily configurable with environment variables and a simple YAML file.

Components:
- `kvdbserver` - The server process
- `kvdb-cli` - CLI tool to manage the server
  
**Note!** This project is in early development. The `main` branch is currently unstable and getting breaking changes! Use a specific release version if you want to use this.

The name of the project will be changed before full release.

# Data types

Currently supported data types:
- `String`
- `HashMap`

# Documentation

- [Using the CLI](./docs/kvdb-cli.md)
- [Using the server](./docs/kvdbserver.md)
- [Databases](./docs/databases.md)
- [Data types](./docs/datatypes.md)
- [API](./docs/api.md)

# Releases

Release notes are available [here](./docs/changelog/).

Binaries are available for download. Multiple platforms supported. You can download them [here](https://github.com/hollowdll/kvdb/releases).

However, it is recommended to build the binaries from source. Instructions on how to build from source below.

Releases are managed with [GoReleaser](https://goreleaser.com/).

# Build binaries

To build the binaries, you first need to install Go. Minimum version required is go1.22.

Instructions for installing Go can be found [here](https://go.dev/doc/install).

You may also need tools to work with gRPC and Protocol Buffers in Go. This is needed if you want to compile `.proto` files and generate Go code.

- [Protocol Buffer compiler](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)
- [Quickstart](https://grpc.io/docs/languages/go/quickstart/)

If you just want to compile the binaries then installing only Go is enough.

After you have successfully installed go, clone this repository.

Cloning with git:
```sh
git clone https://github.com/hollowdll/kvdb.git
```

Note: You can also download the source code for a specific release [here](https://github.com/hollowdll/kvdb/releases).

Change directory to the project root:
```sh
cd kvdb
```

Get the dependencies:
```sh
go mod tidy
```

Build the server:
```sh
go build -o ./bin/kvdbserver/ ./cmd/kvdbserver/
```

Build the CLI:
```sh
go build -o ./bin/kvdb-cli/ ./cmd/kvdb-cli/
```

These will build the binaries to `bin/` directory in the project root. You can change the output directory and binary names to whatever you like by modifying the path with `-o` flag.

For more advanced build, use `go help build` to see more build options.

# Docker

Images are available in Docker Hub with multiple tags. Links below.

- [Repository](https://hub.docker.com/r/hakj/kvdb)
- [Tags](https://hub.docker.com/r/hakj/kvdb/tags)

## Pull the server image

```sh
docker pull hakj/kvdb
```

## Build the server image

Make sure to be in the project root
```sh
cd kvdb
```
Latest tag
```sh
docker build -f "./Dockerfile.bookworm" -t kvdb:latest .
```
Debian based image
```sh
docker build -f "./Dockerfile.bookworm" -t kvdb:bookworm .
```
Alpine Linux based image
```sh
docker build -f "./Dockerfile.alpine" -t kvdb:alpine .
```

These commands build the image only for a single architecture. If you want to build multi-arch images for other platforms, read [this](https://docs.docker.com/build/building/multi-platform/).

## Start a container

Example of starting a container
```sh
docker run -p 12345:12345 --rm -it kvdb
```
This binds the host's port `12345` to the container's port `12345` so you can access the server outside the container.

# Running tests

Change directory to the project root:
```sh
cd kvdb
```

Run all tests:
```sh
go test ./...
```

Run only integration tests:
```sh
go test ./tests/integration
```

Show verbose test result output:
```sh
go test -v ./...
```

## Integration tests

Running the integration tests starts test servers both with and without TLS. The servers are stopped after the tests have finished. 

The operating system assigns random TCP ports for the test servers. You can output the assigned ports to the console by running the integration tests with `-v` flag.

For example:
```sh
go test -v ./tests/integration
```

# License

This project is licensed under MIT license. It is free and open source software.
