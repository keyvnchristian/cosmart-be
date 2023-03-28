# Cosmart-Books

Cosmart-Books is created to fulfill Cosmart Backend Test

## Local

### Install tools

1. Install go: [read here](https://golang.org/dl/)

### Usage

1. To build binaries, execute `make build` or `make build-linux`
2. To run server, execute `make run`
3. To run test, execute `make test`, and it will generate you coverage metrics
4. To send HTTP request and view the response in Visual Studio Code directly, I've prepared `req.http` files or

```bash
curl http://localhost:8080/v1/books/love HTTP/1.1
```
