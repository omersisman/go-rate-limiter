package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const requestsPerIP = 5

var endpoints = []string{"http://localhost:8080/endpoint1", "http://localhost:8080/endpoint2", "http://localhost:8080/endpoint3"}
var ipAddresses = []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}

func main() {
	var wg sync.WaitGroup
	for _, ip := range ipAddresses {
		for _, endpoint := range endpoints {
			wg.Add(1)
			go func(ip, endpoint string) {
				defer wg.Done()
				sendRequest(ip, endpoint)
			}(ip, endpoint)
		}
	}
	wg.Wait()
}

func sendRequest(ip, endpoint string) {
	client := &http.Client{}
	for i := 0; i < requestsPerIP; i++ {
		req, _ := http.NewRequest("GET", endpoint, nil)
		req.Header.Set("X-Forwarded-For", ip)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Request %d from %s to %s: Error %v\n", i+1, ip, endpoint, err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Printf("Request %d from %s to %s: HTTP Status %d, Response: %s\n", i+1, ip, endpoint, resp.StatusCode, body)
		time.Sleep(100 * time.Millisecond)
	}
}
