package transport

import (
	"net"
)

type UDPTransport struct {
	conn *net.UDPConn
}

func NewUDP(port int) (*UDPTransport, error) {
	addr := &net.UDPAddr{Port: port}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPTransport{conn: conn}, nil
}

// Отправка сырого пакета на внешний адрес пира
func (t *UDPTransport) Send(data []byte, to *net.UDPAddr) error {
	_, err := t.conn.WriteToUDP(data, to)
	return err
}

func (t *UDPTransport) Receive(buf []byte) (int, *net.UDPAddr, error) {
	return t.conn.ReadFromUDP(buf)
}
