package types

import (
	"testing"

	"github.com/xoreo/mcserver-backend/commands"
)

func TestLoadDB(t *testing.T) {
	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(db.String())
}

func TestAddServer(t *testing.T) {
	version := "1.12"
	name := "test server"
	var port uint32 = 25565
	var ram uint32 = 1024
	server, err := NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}

	err = commands.InitializeServer(server)
	if err != nil {
		t.Fatal(err)
	}

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[ before AddServer ] %s\n", db.String())

	err = db.AddServer(server)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[ after AddServer ] %s\n", db.String())

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClose(t *testing.T) {
	version := "1.12"
	name := "test server"
	var port uint32 = 25565
	var ram uint32 = 1024
	server, err := NewServer(version, name, port, ram)
	if err != nil {
		t.Fatal(err)
	}

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[ before AddServer ] %s\n", db.String())

	db.AddServer(server)

	t.Logf("[ after AddServer ] %s\n", db.String())

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
}
