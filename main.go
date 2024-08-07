package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type ProxyConfig struct {
	ListenIP   string `json:"listen_ip"`
	SourcePort int    `json:"source_port"`
	TargetIP   string `json:"target_ip"`
	TargetPort int    `json:"target_port"`
}

type Config struct {
	Proxies []ProxyConfig `json:"proxies"`
}

func main() {
	// Read configuration
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	for _, proxy := range config.Proxies {
		go startTCPProxy(proxy)
		go startUDPProxy(proxy)
	}

	// Prevent the main function from exiting
	select {}
}

func startTCPProxy(config ProxyConfig) {
	// Start listening on the specified IP and port
	listenAddr := net.JoinHostPort(config.ListenIP, strconv.Itoa(config.SourcePort))
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", listenAddr, err)
	}
	defer listener.Close()
	log.Printf("Listening on %s for TCP", listenAddr)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleTCPConnection(clientConn, config.TargetIP, config.TargetPort)
	}
}

func handleTCPConnection(clientConn net.Conn, targetIP string, targetPort int) {
	defer clientConn.Close()

	targetAddr := net.JoinHostPort(targetIP, strconv.Itoa(targetPort))
	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Failed to connect to target %s: %v", targetAddr, err)
		return
	}
	defer targetConn.Close()

	// Forward data between client and target
	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)
}

func startUDPProxy(config ProxyConfig) {
	listenAddr := net.JoinHostPort(config.ListenIP, strconv.Itoa(config.SourcePort))
	serverAddr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address %s: %v", listenAddr, err)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", listenAddr, err)
	}
	defer conn.Close()
	log.Printf("Listening on %s for UDP", listenAddr)

	buffer := make([]byte, 65535)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Failed to read from UDP: %v", err)
			continue
		}

		go handleUDPConnection(conn, clientAddr, buffer[:n], config.TargetIP, config.TargetPort)
	}
}

func handleUDPConnection(conn *net.UDPConn, clientAddr *net.UDPAddr, data []byte, targetIP string, targetPort int) {
	targetAddr := net.JoinHostPort(targetIP, strconv.Itoa(targetPort))
	targetUDPAddr, err := net.ResolveUDPAddr("udp", targetAddr)
	if err != nil {
		log.Printf("Failed to resolve target UDP address %s: %v", targetAddr, err)
		return
	}

	_, err = conn.WriteToUDP(data, targetUDPAddr)
	if err != nil {
		log.Printf("Failed to write to target UDP address %s: %v", targetAddr, err)
		return
	}

	// Read response from target
	buffer := make([]byte, 65535)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Printf("Failed to read from target UDP: %v", err)
		return
	}

	// Send response back to client
	_, err = conn.WriteToUDP(buffer[:n], clientAddr)
	if err != nil {
		log.Printf("Failed to write to client UDP address %s: %v", clientAddr, err)
	}
}
