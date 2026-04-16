package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const PORT = ":8080"

func main() {
	// Create logs directory
	os.MkdirAll("logs", 0755)

	// Set up logging
	logFile, err := os.OpenFile("logs/proxy.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Initialize cache
	cache := NewCache()

	// Start listening
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error starting proxy:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Proxy HTTP con Cache corriendo en puerto %s\n", PORT)
	fmt.Println("Presiona Ctrl+C para detener.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		// Handle each client in a goroutine (concurrencia)
		go handleClient(conn, cache)
	}
}
