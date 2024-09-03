package network

import (
	"fmt"
	"testing"
)

func TestIsPortListening(t *testing.T) {
	t.Run("Port 22 is listening", func(t *testing.T) {
		port := 22
		proto := "tcp"
		got, err := IsPortListening(port, proto, true)
		if err != nil {
			t.Errorf("IsPortListening() error = %v", err)
			return
		}
		if got != true {
			t.Errorf("IsPortListening() = %v, want %v", got, true)
		}
	})

	t.Run("Port is not listening", func(t *testing.T) {
		port := 65535
		proto := "tcp"
		got, err := IsPortListening(port, proto, true)
		if err == nil {
			t.Errorf("IsPortListening() error = %v", err)
			return
		}
		if got != false {
			t.Errorf("IsPortListening() = %v, want %v", got, false)
		}
	})
}

func TestIsDHCPPortOpen(t *testing.T) {
	t.Run("Port 67 is open", func(t *testing.T) {
		got, err := IsUdpListening(67, "udp", true)
		if err != nil {
			t.Errorf("IsPortOpen() error = %v", err)
		}
		if !got {
			t.Errorf("IsPortOpen() = %v, want %v", got, true)
		}
	})
}

func TestIsTFTPServiceListening(t *testing.T) {
	t.Run("TFTP service is not listening", func(t *testing.T) {
		got, err := IsUdpListening(69, "udp", true)
		if err == nil {
			t.Errorf("IsTFTPServiceListening() error = %v", fmt.Errorf("TFTP service is listening"))
			return
		}
		if got {
			t.Errorf("IsTFTPServiceListening() = %v, want %v", got, false)
		}
	})
}
