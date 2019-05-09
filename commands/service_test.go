package commands

import (
	"testing"

	"github.com/xoreo/mcserver-backend/types"
)

func TestGenerateService(t *testing.T) {
	version := "1.12"
	name := "test-server"
	var port uint32 = 25565
	var ram uint32 = 1024

	server, err := types.NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}

	service, err := GenerateService(*server)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(service)
}

func TestInstallService(t *testing.T) {
	server, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	service, err := GenerateService(*server)

	err = InstallService(service, server.Name)
	if err != nil {
		t.Fatal(err)
	}

}
