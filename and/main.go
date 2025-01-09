package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.IPv4(192, 168, 1, 10)
	subnetMask := net.IPv4Mask(255, 255, 255, 0)

	ipBytes := ip.To4()
	maskBytes := subnetMask

	network := net.IPv4(
		ipBytes[0]&maskBytes[0],
		ipBytes[1]&maskBytes[1],
		ipBytes[2]&maskBytes[2],
		ipBytes[3]&maskBytes[3],
	)

	fmt.Printf("IP Address: %s\n", ip.String())
	fmt.Printf("Subnet Mask: %d\n", subnetMask)
	fmt.Printf("Network: %s\n", network.String())
}
