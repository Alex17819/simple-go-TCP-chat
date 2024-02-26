package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func receiveMessages(conn net.Conn) {
    defer conn.Close()

    // Primim mesajele de la server și le afișăm în consolă
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("Mesaj de la server:", scanner.Text())
    }
}

func main() {
    // Stabilim conexiunea cu serverul
    conn, err := net.Dial("tcp", "localhost:12345")
    if err != nil {
        fmt.Println("Eroare la conectarea la server:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Conectat la serverul localhost:12345")

    // Pornim o goroutine pentru a primi și afișa mesajele de la server
    go receiveMessages(conn)

    // Trimitem mesaje către server
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        message := scanner.Text()
        _, err := conn.Write([]byte(message + "\n"))
        if err != nil {
            fmt.Println("Eroare la trimiterea mesajului către server:", err)
            break
        }
    }
}
