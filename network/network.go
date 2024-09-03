package network

import (
	"fmt"
	"net"

	"github.com/nestorlai1994/nesoli/logger"
)

func IsPortListening(port int, proto string, showlog bool) (bool, error) {
	// Check if port is listening
	ln, err := net.Listen(proto, fmt.Sprintf(":%d", port))
	if err != nil {
		if showlog {
			logger.Info("Port is listening: " + fmt.Sprintf(":%d", port))
		}
		return true, nil
	}
	defer ln.Close()
	return false, fmt.Errorf("port is not listening")
}

func IsUdpListening(port int, proto string, showlog bool) (bool, error) {
	// Check if DHCP service is running
	conn, err := net.ListenPacket(proto, fmt.Sprintf(":%d", port))
	if err != nil {
		if showlog {
			logger.Info(fmt.Sprintf("%d/%s service is running", port, proto))
		}
		return true, nil
	}
	defer conn.Close()
	return false, fmt.Errorf("%d/%s service is not running", port, proto)
}
