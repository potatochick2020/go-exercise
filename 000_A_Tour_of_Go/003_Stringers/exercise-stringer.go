package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// String returns the IP address in the format "x.x.x.x".
func (ipaddr IPAddr) String() string {
	// Convert each byte to a string and join with dots
	var parts []string
	for _, b := range ipaddr {
		parts = append(parts, strconv.Itoa(int(b)))
	}
	return strings.Join(parts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
