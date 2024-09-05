package ioutil

import "testing"

func TestDelete(t *testing.T) {
	CheckOrCreateFile("/go/.test/tmp/test.txt")
	RemoveFolder("/go/.test/tmp/test.txt")
}
