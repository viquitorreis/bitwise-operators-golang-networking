package main

import "fmt"

func main() {
	fmt.Printf("/24 network has %d addresses\n", calculateNetworkSize(24))
	fmt.Printf("/16 network has %d addresses\n", calculateNetworkSize(16))
}

func calculateNetworkSize(cidr int) int {
	hostBits := 32 - cidr
	return 1 << hostBits // 2^hostBits
}
