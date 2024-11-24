// main.go

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// Create a map to store the clients
var clients = make(map[net.Conn]bool)
var mu sync.Mutex

// Function to handle communication with each client
func handleClient(conn net.Conn) {
	defer conn.Close()
	// Add the client to the map
	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	// Notify other clients
	message := fmt.Sprintf("New user has joined the chat: %s\n", conn.RemoteAddr().String())
	broadcast(message, conn)

	// Read messages from the client and broadcast them to others
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// Remove the client from the map when they disconnect
			mu.Lock()
			delete(clients, conn)
			mu.Unlock()

			// Notify other clients about the disconnection
			broadcast(fmt.Sprintf("User %s has left the chat.\n", conn.RemoteAddr().String()), conn)
			break
		}
		message := fmt.Sprintf("%s: %s", conn.RemoteAddr().String(), string(buffer[:n]))
		broadcast(message, conn)
	}
}

// Broadcast messages to all connected clients except the sender
func broadcast(message string, sender net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	// Send the message to all clients
	for client := range clients {
		if client != sender {
			client.Write([]byte(message))
		}
	}
}

func main() {
	// Listen on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Server started on port 8080...")

	// Accept new connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each client in a new goroutine
		go handleClient(conn)
	}
}
