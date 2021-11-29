# ByteBank API

Simple API server for [ByteBank](https://github.com/davidgaspardev/bytebank) mobile app add-on. This server was developed 100% in the programming language Go. 

## Installations

Prerequisite to install Bytebank API:
 - [GoLang](https://golang.org) to build the application.
 - [MongoDB](
https://docs.mongodb.com/manual/installation/) local or [MongoDB Atlas](https://www.mongodb.com/atlas/database) to store the transactions.

### MongoDB Atlas

If you use MongoDB Atlas, you need to declare and initialize the following environment variables:
 - `MONGO_USER`: Username to connect with MongoDB
 - `MONGO_PASS`: Password to connect with MongoDB
 - `MONGO_HOST`: MongoDB address (DNS/IP)
 - `MONGO_DB`: Database name declared in the MongoDB
> Obs: You can use a Makefile to set the environment variables.

### Authetication configurations

You need to declare and initialize the following environment variables:
- Basic HTTP Authentication Scheme. (See more: [RFC 7617](https://datatracker.ietf.org/doc/html/rfc7617))
    - `AUTH_USER`: Anything username (Ex: elonmusk)
    - `AUTH_PASS`: Anything password (Ex: jeff%1993)
    > Obs: External applications that have consumed this API will need to know this username and password for successful communication.
- `PASS_POST`: Password to validate transfer storage in database

### Application build

Download the necessary dependencies to run the application:
```bash
go get
```

Now just run the application with the following command:

(with MongoDB Atlas)
```bash
MONGO_USER=<your user> MONGO_PASS=<your pass> MONGO_HOST=<your host> MONGO_DB=<your database name>  AUTH_USER=<anything> AUTH_PASS=<anything> POST_PASS=<anything> go run main.go
```
(with local MongoDB)
```bash
AUTH_USER=<anything> AUTH_PASS=<anything> POST_PASS=<anything> go run main.go
```