package main

import (
	"fmt"
	"net"
)

func main() {
	ip1 := "10.0.0.10"
	ip2 := "10.0.0.20"
	mask := "255.255.255.0"
	ip3 := "10.0.2.20"

	fmt.Printf("IPs %s and %s are in the same subnet? %t\n", ip1, ip2, checkSameSubnet(ip1, ip2, mask))
	fmt.Printf("IPs %s and %s are in the same subnet? %t\n", ip1, ip3, checkSameSubnet(ip1, ip3, mask))
}

func checkSameSubnet(ip1, ip2 string, mask string) bool {
	ip1Addr := net.ParseIP(ip1).To4()
	ip2Addr := net.ParseIP(ip2).To4()
	subnetMask := net.IPMask(net.ParseIP(mask).To4())

	// OR between the IP addresses to get the difference between them
	xorResult := make([]byte, 4)
	for i := range 4 {
		xorResult[i] = ip1Addr[i] ^ ip2Addr[i]
	}

	fmt.Printf("Difference in bits: %08b\n", xorResult)
	fmt.Printf("Difference in decimal: %d.%d.%d.%d\n", xorResult[0], xorResult[1], xorResult[2], xorResult[3])

	andResult := make([]byte, 4)
	for i := range 4 {
		fmt.Printf("XOR result: %08b Subnet mask: %08b\n", xorResult[i], subnetMask[i])
		andResult[i] = xorResult[i] & subnetMask[i]
	}

	// checking if they are all zero (equal)
	for i := range 4 {
		if andResult[i] != 0 {
			return false
		}
	}

	return true
}
