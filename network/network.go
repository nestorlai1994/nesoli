package network

import (
	"fmt"
	"net"

	"github.com/nestorlai1994/nesoli/logger"
)

func IsPortListening(port int, proto string) (bool, error) {
	// Check if port is listening
	ln, err := net.Listen(proto, fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Info("Port is listening: " + fmt.Sprintf(":%d", port))
		return true, nil
	}
	defer ln.Close()
	return false, fmt.Errorf("port is not listening")
}
