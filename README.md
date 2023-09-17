# URL Shortener using GO Lang

This is a simple URL shortener project created using Golang

## Overview

This project implements a basic URL shortener application and has the following key aspects:

- HTTP routing and handling.
- Database interaction (even if just for demonstration).
- Basic error handling and validation.
- Unit tests for each module and file

## Setup

1. Install project dependencies using the command
   `go mod download`

2. Run `setup.sh` file to automatically start the server

## API Endpoints

This project exposes two basic endpoints `GET` and `POST` to create a shortened
URL and to fetch the original url from the shortened url

- `POST` endpoint
  - `/api/shorten`
  - Body :
    ```
    {
        url: "url-string"
    }
    ```
- `GET` endpoint
  - `/{short_url_hash}`
  - Path parameters - `short_url_hash` provided by `/api/shorten` endpoint

## Tests

File unit tests have been added for modules `helpers`(*fully) and
`controllers`(*partially) and can be run by using the command
`go test ./... -v -cover`

TODO

- [x] Unit tests for `helpers`
- [ ] Unit tests for `controllers`
