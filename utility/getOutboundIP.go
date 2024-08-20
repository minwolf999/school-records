package utility

import (
	"net"
)

// These function get the local IPV4 address and return if as a string (if the IPV4 isn't found he return "localhost")
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "localhost"
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
