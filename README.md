# Chat-Application-with-Concurrency-in-Go
# Concurrency in Go: Chat Application

This project demonstrates Go's concurrency model by implementing a simple chat application. Multiple users can connect to the server and exchange messages simultaneously. The application uses goroutines to handle concurrent client connections.

## Project Structure

The project consists of the following files:

- `main.go`: The Go code for the chat application.
- `README.md`: The documentation for the project.

## Key Concepts

### 1. **Concurrency with Goroutines**
The chat server uses Go’s goroutines to handle multiple client connections concurrently. Each client connection is handled in a separate goroutine, allowing users to interact with the chat application without blocking the server.

### 2. **Synchronization with Mutex**
The server uses a `sync.Mutex` to safely manage access to the `clients` map, ensuring that only one goroutine modifies the map at a time. This prevents race conditions when adding or removing clients.

### 3. **TCP Server**
The server listens on port 8080 for incoming TCP connections. When a new client connects, a new goroutine is spawned to handle that connection. The client can send messages, and the server will broadcast those messages to all other connected clients.

## How It Works

1. The server starts and listens on TCP port 8080.
2. When a client connects, the server creates a new goroutine to handle that client.
3. The client sends messages, and the server broadcasts those messages to all other connected clients.
4. When a client disconnects, the server removes the client and notifies the others.

## Running the Chat Application

1. Install Go from [https://golang.org/dl/](https://golang.org/dl/).
2. Clone the project and navigate to the project folder.
3. Run the server with the following command:
   ```bash
   go run main.go
