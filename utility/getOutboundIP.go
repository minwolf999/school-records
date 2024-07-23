package utility

import (
	"net"
)

// These function get the local IPV4 address and return if as a string (if the IPV4 isn't found he return "localhost")
func GetOutboundIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	return "localhost"
}
