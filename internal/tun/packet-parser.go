package tun

import (
	"fmt"

	"golang.org/x/net/ipv6"
)

func HandlePacket(data []byte) {
	// Парсим заголовок IPv6
	header, err := ipv6.ParseHeader(data)
	if err != nil {
		// Возможно, прилетел IPv4 пакет, его пока игнорируем
		return
	}

	fmt.Printf("--- Пакет пойман ---\n")
	fmt.Printf("Откуда: %s\n", header.Src)
	fmt.Printf("Куда:   %s\n", header.Dst)
	fmt.Printf("Длина данных (Payload): %d байт\n", header.PayloadLen)
	fmt.Printf("Протокол: %d (например, 17=UDP, 6=TCP, 58=ICMPv6)\n", header.NextHeader)
}
