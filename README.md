# Chat Application using TCP in Go

This is a simple chat application implemented in Go using TCP. It consists of a server and multiple clients that can connect to the server to exchange messages.

## Instructions to Run

### Prerequisites

- Go installed on your machine. If not, you can download it from [here](https://golang.org/dl/).

### Steps to Run

1. **Clone the Repository:**

```bash
    git clone https://github.com/Alex17819/simple-go-TCP-chat.git
    cd simple-go-TCP-chat
```
2. **Compile the Server and Client:**

```bash
    go build -o server server.go
    go build -o client client.go
```
3. **Run the Server:**

```bash
    ./server
```

4. **Run Multiple Instances of the Client:**

```bash
    ./client
```
