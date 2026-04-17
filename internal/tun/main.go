package tun

import (
	"fmt"
	"log"

	"golang.zx2c4.com/wireguard/tun"
)

func CreateTun() {
	const tunName = "mytun0"
	const mtu = 1420

	// Создаем TUN-устройство.
	// На Linux это создаст стандартный tun-интерфейс.
	// На Windows потребуется наличие wintun.dll.
	dev, err := tun.CreateTUN(tunName, mtu)
	if err != nil {
		log.Fatalf("Ошибка создания TUN: %v", err)
	}
	defer dev.Close()

	fmt.Printf("Интерфейс %s успешно запущен!\n", tunName)

	// Читаем пакеты в бесконечном цикле
	// Пакет в wireguard-go читается в [][]byte для поддержки многопоточности (batching)
	packet := make([]byte, mtu)
	for {
		// Метод Read принимает срез буферов (для одного пакета достаточно одного буфера)
		sizes := make([]int, 1)
		n, err := dev.Read([][]byte{packet}, sizes, 0)
		if err != nil {
			log.Printf("Ошибка чтения: %v", err)
			continue
		}

		HandlePacket(packet)

		if n > 0 {
			fmt.Printf("Получен пакет длиной %d байт\n", sizes[0])
			// Здесь будет логика обработки (Этап 2 и 3)
		}
	}
}
