package tun

import "golang.zx2c4.com/wireguard/tun"

type Device struct {
	tun.Device
	MTU int
}

func New(name string, mtu int) (*Device, error) {
	dev, err := tun.CreateTUN(name, mtu)
	if err != nil {
		return nil, err
	}
	return &Device{Device: dev, MTU: mtu}, nil
}
