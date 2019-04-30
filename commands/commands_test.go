package commands

import (
	"testing"

	"github.com/xoreo/mcserver-backend/types"
)

func TestInitializeServer(t *testing.T) {
	server1, err := types.NewServer("1.7.2", "test-server-1", 25565)
	if err != nil {
		t.Fatal(err)
	}

	err = InitializeServer(server1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStartServer(t *testing.T) {

}
