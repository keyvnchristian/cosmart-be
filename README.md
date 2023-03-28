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

curl -X POST http://localhost:8080/v1/schedule
   -H 'Content-Type: application/json'
   -d '{
	"user_id": 2,
	"title": "Jesus is Lord",
	"author": "Paul the Apostle",
	"edition": 20,
	"pickup_schedule": "2023-10-11"
}'
```
