package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

const CACHE_TTL = 30 * time.Second

func handleClient(clientConn net.Conn, cache *Cache) {
	defer clientConn.Close()

	reader := bufio.NewReader(clientConn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	cacheKey := req.Method + ":" + req.URL.String()

	if req.Method == http.MethodGet {
		if cached, found := cache.Get(cacheKey); found {
			log.Printf("[CACHE HIT]  %s %s", req.Method, req.URL)
			fmt.Printf("[CACHE HIT]  %s %s\n", req.Method, req.URL)
			clientConn.Write(cached)
			return
		}
		log.Printf("[CACHE MISS] %s %s", req.Method, req.URL)
		fmt.Printf("[CACHE MISS] %s %s\n", req.Method, req.URL)
	} else {
		log.Printf("[BYPASS]     %s %s", req.Method, req.URL)
		fmt.Printf("[BYPASS]     %s %s\n", req.Method, req.URL)
	}

	host := req.Host
	if host == "" {
		host = req.URL.Host
	}

	if _, _, err := net.SplitHostPort(host); err != nil {
		host = host + ":80"
	}

	serverConn, err := net.Dial("tcp", host)
	if err != nil {
		log.Printf("Error connecting to origin server %s: %v", host, err)
		clientConn.Write([]byte("HTTP/1.1 502 Bad Gateway\r\n\r\nProxy: could not connect to origin server\r\n"))
		return
	}
	defer serverConn.Close()

	req.Header.Set("Connection", "close")
	req.Header.Del("Proxy-Connection")
	if err := req.Write(serverConn); err != nil {
		log.Println("Error forwarding request:", err)
		return
	}

	var responseBuf bytes.Buffer
	tee := io.TeeReader(serverConn, &responseBuf)

	written, err := io.Copy(clientConn, tee)
	if err != nil && written == 0 {
		log.Println("Error forwarding response:", err)
		return
	}

	if req.Method == http.MethodGet && responseBuf.Len() > 0 {
		cache.Set(cacheKey, responseBuf.Bytes(), CACHE_TTL)
		log.Printf("[CACHED]     %s (TTL: %v, cache size: %d)", cacheKey, CACHE_TTL, cache.Size())
	}
}
