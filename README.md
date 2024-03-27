# Multithreaded gRPC Server and Client

This implementation features a multithreaded server and client setup. The server provides two services:

- Fibonacci function
- Primality check of a number

Clients are limited by the number of connections, and they are served for an exact maximum number of times fixed by the user at the launch of the server.

## Test Instructions

1. Navigate to the server and client directories.
2. Run `go run .` in each directory, starting with the server.

For example:
```bash
cd server
go run .
