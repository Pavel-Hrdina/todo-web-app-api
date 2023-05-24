# Todo backend

## Table of content

- [Todo backend](#todo-backend)
  - [Table of content](#table-of-content)
  - [Getting Started](#getting-started)
    - [Prerequisite](#prerequisite)
    - [Installation](#installation)
      - [Prefer `make file`](#prefer-make-file)
      - [Via `go`](#via-go)
  - [Running localy](#running-localy)
    - [Using air](#using-air)
    - [Using Docker](#using-docker)
  - [API references](#api-references)
  - [Contributing](#contributing)
  - [Contact](#contact)
  - [Acknowledgments](#acknowledgments)

## Getting Started

### Prerequisite

Make sure you have the following installed outside the current project directory and available in your GOPATH

- golang
- [air](https://github.com/cosmtrek/air) for hot reloading
- [godotenv](https://github.com/joho/godotenv) for loading .env file

### Installation 

> :mailbox: **Notice**: The make file works only on unix

#### Prefer `make file`

```sh
# the binary will be in $(root)/tmp/main
make build
```

#### Via `go`

Make sure you have go 1.18 or higher:

* Clone the repo... 

* Run  
  ```sh
  go get
  ```

* install using makefile
  ```sh
  # make install not yet working
  make install
  ```

## Running localy

> :warning: **Important**:  because this is a CGO enabled package, you are required to set the environment variable CGO_ENABLED=1 and **have a gcc compiler present within your path** if you are building locally

### Using air

1. Type ``air`` in the command line

2. Use the command ``make watch``  


### Using Docker

* Run 
   ```sh
  make docker-build
  ```
* and
  ```sh
  docker run main:main
  ```

## API references

TODO

## Contributing

TODO

## Contact

TODO

## Acknowledgments

TODO



