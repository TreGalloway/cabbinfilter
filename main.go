package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    address := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    
    if err != nil {
        // Port is closed or filtered
        return
    }
    
    conn.Close()
    fmt.Printf("Port %d: Open\n", port)
}

func main() {
    host := "localhost"
    startPort := 1
    endPort := 1000
    
    var wg sync.WaitGroup
    
    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(host, port, &wg)
    }
    
    wg.Wait()
    fmt.Println("Scan complete")
}

