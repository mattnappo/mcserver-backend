package commands

import (
	"testing"

	"github.com/xoreo/mcserver-backend/types"
)

func getTestServer() (*types.Server, error) {
	server, err := types.NewServer("1.7.2", "test-server", 25565, 1024)
	if err != nil {
		return nil, err
	}

	err = InitializeServer(server)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func TestInitializeServer(t *testing.T) {
	server1, err := types.NewServer("1.7.2", "test-server", 25565, 1024)
	if err != nil {
		t.Fatal(err)
	}

	err = InitializeServer(server1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartServer(t *testing.T) {
	server, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}
	err = StartServer(*server)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRestartServer(t *testing.T) {
	server, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	RestartServer(*server)
}

func TestStopServer(t *testing.T) {
	server, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	RestartServer(*server)
}
