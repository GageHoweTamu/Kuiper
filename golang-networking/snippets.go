package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

const (
	bootstrapNodeIP = "127.0.0.1"
	port          = "8080"
)

type peer struct {
	ip    string
	conn  net.Conn
	peerX int
	peerY int
}

var (
	peers         = make(map[string]*peer)
	peerLock      = sync.RWMutex{}
	bootstrapPeer *peer
)


func main() {
	go startNode()

	// Connect to the bootstrap node
	conn, err := net.Dial("tcp", bootstrapNodeIP+":"+port)
	if err != nil {
		fmt.Printf("Failed to connect to bootstrap node: %v\n", err)
		return
	};
	if err == nil {
		fmt.Printf("Connected to bootstrap node: %s\n", bootstrapNodeIP)
	}

	// Add the bootstrap node to the peer list
	addPeer(bootstrapNodeIP, conn)
	bootstrapPeer = peers[bootstrapNodeIP]

	// Discover and connect to other peers
	discoveredPeers := discoverPeers(conn)
	for _, peerAddr := range discoveredPeers {
		conn, err := net.Dial("tcp", peerAddr)
		if err != nil {
			fmt.Printf("Failed to connect to peer %s: %v\n", peerAddr, err)
			continue
		}
		fmt.Printf("Connected to peer: %s\n", peerAddr)
		addPeer(peerAddr, conn)
	}

	// Request the peer list from the bootstrap node
	_, err = bootstrapPeer.conn.Write([]byte("PEER_LIST\n"))
	if err != nil {
		fmt.Printf("Failed to request peer list: %v\n", err)
	}
	select {} // Keeps the main goroutine alive
}

func getIPFromName(name string) string {
	ips, err := net.LookupIP(name)
	if err != nil { return "" }
	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 != nil {
			fmt.Printf("Found IPv4 address for %s: %s\n", name, ipv4.String())
			return ipv4.String()
		}
	}
	fmt.Printf("No IPv4 address found for %s\n", name)
	return ""
}

func startNode() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Failed to start node: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Node started and listening on %s\n", port)
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
	remoteAddr := conn.RemoteAddr().String()
	addPeer(remoteAddr, conn)
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read from connection: %v\n", err)
			removePeer(remoteAddr)
			return
		}
		message = strings.TrimSpace(message)
		handleMessage(remoteAddr, message)
	}
}

func handleMessage(sender, message string) {
	fmt.Printf("Received message from %s: %s\n", sender, message)
	parts := strings.Split(message, " ")
	switch parts[0] {
	case "BROADCAST":
		broadcastMessage(sender, strings.Join(parts[1:], " "))
	case "PEER_LIST":
		sendPeerList(sender)
	default:
		fmt.Printf("Unknown message type: %s\n", parts[0])
	}
}

func broadcastMessage(sender, message string) {
	peerLock.RLock()
	defer peerLock.RUnlock()
	for addr, peer := range peers {
		if addr != sender {
			_, err := peer.conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Printf("Failed to send message to %s: %v\n", addr, err)
				removePeer(addr)
			}
		}
	}
}

func sendPeerList(recipient string) {
	peerLock.RLock()
	defer peerLock.RUnlock()
	peerList := "PEER_LIST"
	for addr, peer := range peers {
		if addr != recipient {
			peerList += fmt.Sprintf(" %s %d %d", addr, peer.peerX, peer.peerY)
		}
	}

	peer, ok := peers[recipient]
	if ok {
		_, err := peer.conn.Write([]byte(peerList + "\n"))
		if err != nil {
			fmt.Printf("Failed to send peer list to %s: %v\n", recipient, err)
			removePeer(recipient)
		}
	}
}

func discoverPeers(conn net.Conn) []string {
	// Simulate peer discovery by returning a hardcoded list of peers
	// In a real implementation, this would involve exchanging peer lists with connected nodes
	return []string{"127.0.0.1:8001", "127.0.0.1:8002"}
}

func addPeer(addr string, conn net.Conn) {
	peerLock.Lock()
	defer peerLock.Unlock()
	x, y := generateCoordinates()
	peers[addr] = &peer{ip: addr, conn: conn, peerX: x, peerY: y}
}

func removePeer(addr string) {
	peerLock.Lock()
	defer peerLock.Unlock()

	delete(peers, addr)
}

func generateCoordinates() (int, int) {
	// Implement your logic to generate random coordinates
	// For simplicity, we'll return fixed values
	return 100, 200
}