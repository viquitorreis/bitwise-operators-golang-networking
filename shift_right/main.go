package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	ip := net.IPv4(192, 162, 191, 241).To4()
	ipUint32 := binary.BigEndian.Uint32(ip)
	octets := ipToOctets(ipUint32)
	fmt.Printf("Octet - 1: %08b, 2: %08b, 3: %08b, 4: %08b\n",
		octets[0], octets[1], octets[2], octets[3],
	)
}

func ipToOctets(ip uint32) [4]byte {
	return [4]byte{
		byte(ip >> 24), // first 8 bits
		byte(ip >> 16), // 9 bit - 15 bit
		byte(ip >> 8),  // 16 bit - 24 bit
		byte(ip),       // 25 - 32 bit
	}
}
