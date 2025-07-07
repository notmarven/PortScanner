package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"
)

type PortResult struct {
	Port int
	Open bool
}

var mu sync.Mutex
var results []PortResult

func scanPort(host string, port int, wg *sync.WaitGroup) {

	// defer means run at the end no matter what
	defer wg.Done()

	// Create a address ("mc.example.net:25565")
	address := host + ":" + strconv.Itoa(port)

	// Try to connect to it
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	// Lock it to prevent multiple goroutines to access at the same time
	mu.Lock()
	defer mu.Unlock() // Unlock it at last

	// For errors, assume port is closed & filter it
	if err != nil {
		// Add to list as closed port
		results = append(results, PortResult{Port: port, Open: false})
		return
	}

	// No errors
	conn.Close()

	// Add to list as open port
	results = append(results, PortResult{Port: port, Open: true})

}

// Main entry point

func main() {

	var host string
	var minPort int
	var maxPort int

	fmt.Print("Enter your host: ")
	fmt.Scan(&host)

	fmt.Print("Enter min port: ")
	fmt.Scan(&minPort)

	fmt.Print("Enter min port: ")
	fmt.Scan(&maxPort)

	// Track time taken since start
	var start = time.Now()

	var wg sync.WaitGroup

	// Scan a range of ports
	for port := minPort; port <= maxPort; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}

	wg.Wait()

	// Sort the results from the PortResult list
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})

	// Print the function
	for _, r := range results {
		if r.Open {
			fmt.Printf("Port %d is open\n", r.Port)
		} else {
			fmt.Printf("Port %d is closed\n", r.Port)
		}
	}

	var elapsed = time.Since(start)

	fmt.Println("Time taken: ", &elapsed)

}
