package server

import (
	"fmt"
	"net"
)

const (
    port = "8080"
)

func main() {
    listener, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
        return
    }
    defer listener.Close()
    fmt.Printf("Server started and listening on %s\n", port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Printf("Failed to accept connection: %v\n", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    // Handle the connection
}