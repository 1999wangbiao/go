package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// CheckURL checks if a URL is valid by sending a GET request and returning the status code and any error encountered.
func CheckURL(url string) (bool, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Check if the status code indicates success (2xx range)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true, nil
	}
	return false, fmt.Errorf("received status code %d", resp.StatusCode)
}

// GenerateIPs generates a list of IP addresses in the given range.
func GenerateIPs(startIP, endIP string) ([]string, error) {
	var ips []string

	ip := net.ParseIP(startIP)
	for ; !ip.Equal(net.ParseIP(endIP)); ip = nextIP(ip) {
		ips = append(ips, ip.String())
	}
	ips = append(ips, endIP) // include the end IP

	return ips, nil
}

// nextIP generates the next IP address.
func nextIP(ip net.IP) net.IP {
	next := net.IPv4(ip[12], ip[13], ip[14], ip[15]+1)
	return next
}
func exe(url string, wg3 *sync.WaitGroup) {
	valid, err := CheckURL(url)
	if valid {
		fmt.Printf("URL %s is valid\n", url)
	} else {
		fmt.Printf("URL %s is invalid: %v\n", url, err)
	}
	wg3.Done()
}

func main() {
	startIP := "192.168.1.250"
	endIP := "192.168.1.254"
	wg3 := sync.WaitGroup{}
	ips, err := GenerateIPs(startIP, endIP)
	if err != nil {
		fmt.Printf("Error generating IPs: %v\n", err)
		return
	}

	for _, ip := range ips {
		wg3.Add(1)
		url := fmt.Sprintf("https://%s", ip)
		go exe(url, &wg3)

	}
	wg3.Wait()
}
