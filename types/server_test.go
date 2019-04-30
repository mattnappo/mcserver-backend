package types

import "testing"

func TestNewServer(t *testing.T) {
	version := "1.12"
	name := "my-test-server"
	port := 25565

	_, err := NewServer(version, name, port)
	if err != nil {
		t.Fatal(err)
	}

}
