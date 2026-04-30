package main

import (
	"fmt"
	"mashroom/internal/identity"
)

func main() {
	// tun.CreateTun()
	// identity.SaveKeys(identity.GenerateKey())

	pub, _ := identity.LoadPublicKey("keys.pub")
	// fmt.Println(pub)

	// priv, _ := identity.LoadPrivateKey("keys")
	// fmt.Println(priv)

	ip := identity.MakeIPv4(pub)
	fmt.Println(ip)
}
