package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var clients map[net.Conn]struct{}
var mutex sync.Mutex

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Adăugăm clientul la map-ul de clienți conectați
	mutex.Lock()
	clients[conn] = struct{}{}
	mutex.Unlock()

	// Buclă pentru a primi și retransmite mesajele de la client
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()

		// Afisăm mesajul în consola serverului
		fmt.Printf("Mesaj de la %s: %s\n", conn.RemoteAddr(), message)

		// Retransmitem mesajul către toți ceilalți clienți conectați
		mutex.Lock()
		for client := range clients {
			if client != conn {
				client.Write([]byte(message + "\n"))
			}
		}
		mutex.Unlock()
	}

	// Ștergem clientul din map-ul de clienți conectați la deconectare
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
}

func main() {
	// Inițializăm map-ul de clienți conectați
	clients = make(map[net.Conn]struct{})

	// Pornim serverul și ascultăm conexiuni pe portul specificat
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Eroare la pornirea serverului:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Serverul este pornit și ascultă pe portul 12345")

	// Buclă infinită pentru acceptarea conexiunilor de la clienți
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Eroare la acceptarea conexiunii:", err)
			continue
		}
		fmt.Println("S-a conectat un nou client:", conn.RemoteAddr())

		// Gestionează clientul într-o goroutine separată
		go handleClient(conn)
	}
}
