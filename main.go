package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
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
	// Define a command-line flag for the config file path
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	// Open the config file
	file, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	// Decode the config file
	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	// Rest of the application logic...
}
