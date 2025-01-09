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

	wildCard := net.IPv4(
		^maskBytes[0],
		^maskBytes[1],
		^maskBytes[2],
		^maskBytes[3],
	)

	broadcast := net.IPv4(
		ipBytes[0]|wildCard[0],
		ipBytes[1]|wildCard[1],
		ipBytes[2]|wildCard[2],
		ipBytes[3]|wildCard[3],
	)

	fmt.Printf("IP Address: %s\n", ip.String())
	fmt.Printf("Subnet Mask: %d\n", subnetMask)
	fmt.Printf("WildCard: %s\n", wildCard.String())
	fmt.Printf("Broadcast: %s\n", broadcast.String())
}
