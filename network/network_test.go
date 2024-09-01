package network

import "testing"

func TestIsPortListening(t *testing.T) {
	t.Run("Port 22 is listening", func(t *testing.T) {
		port := 22
		proto := "tcp"
		got, err := IsPortListening(port, proto)
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
		got, err := IsPortListening(port, proto)
		if err == nil {
			t.Errorf("IsPortListening() error = %v", err)
			return
		}
		if got != false {
			t.Errorf("IsPortListening() = %v, want %v", got, false)
		}
	})
}
