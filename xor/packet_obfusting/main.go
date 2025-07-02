package main

import "fmt"

func main() {
	originalPacket := []byte("Hello Network!")
	key := byte(0x5A)

	// Obfuscating package
	obfuscatedPacket := make([]byte, len(originalPacket))
	for i := 0; i < len(originalPacket); i++ {
		obfuscatedPacket[i] = originalPacket[i] ^ key
	}

	// De-obfuscate packet (XOR is reversible!)
	deobfuscatedPacket := make([]byte, len(obfuscatedPacket))
	for i := 0; i < len(obfuscatedPacket); i++ {
		deobfuscatedPacket[i] = obfuscatedPacket[i] ^ key
	}

	fmt.Printf("Original: %s\n", originalPacket)
	fmt.Printf("Obfuscated: %s\n", string(obfuscatedPacket))
	fmt.Printf("De-obfuscated: %s\n", deobfuscatedPacket)
}
