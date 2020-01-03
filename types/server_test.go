package types

import "testing"

func TestNewServer(t *testing.T) {
	version := "1.12"
	name := "test server 1  1 1"
	port := 25565
	ram := 1024

	server, err := NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(server.String())
}

func TestPurge(t *testing.T) {
	version := "1.12"
	name := "test server to purge"
	port := 25565
	ram := 1024

	server, err := NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(server.String())
}
