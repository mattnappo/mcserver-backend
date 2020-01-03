package commands

import (
	"testing"

	"github.com/xoreo/mcserver-backend/types"
)

func TestGenerateService(t *testing.T) {
	version := "1.12"
	name := "test-server"
	port := 25565
	ram := 1024

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

	_, err = InstallService(service, server.Name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUninstallService(t *testing.T) {

}
