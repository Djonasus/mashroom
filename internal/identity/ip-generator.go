package identity

import (
	"crypto/ed25519"
	"crypto/sha512"
	"net"
)

func MakeIPv6(pubKey ed25519.PublicKey) net.IP {
	hash := sha512.Sum512(pubKey)

	ip := make(net.IP, 16)

	ip[0] = 0x02

	copy(ip[1:], hash[:15])

	return ip
}

// For tests. Do not use
func MakeIPv4(pubKey ed25519.PublicKey) net.IP {
	hash := sha512.Sum512(pubKey)
	ip := net.IPv4(10, hash[0], hash[1], hash[2])
	return ip
}
