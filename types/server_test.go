package types

import "testing"

func TestNewServer(t *testing.T) {
	version := "1.12"
	name := "test-server"
	var port uint32 = 25565
	var ram uint32 = 1024

	server, err := NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}
	t.Log((*server).String())
}
