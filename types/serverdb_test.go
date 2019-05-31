package types

import (
	"testing"
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
	port := 25565
	ram := 1024

	server, err := NewServer(version, name, port, ram)
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
	port := 25565
	ram := 1024
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

func TestUpdateSever(t *testing.T) {
	testServer, err := NewServer("1.12", "test", 9999, 1024)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(testServer.String())

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AddServer(testServer)
	if err != nil {
		t.Fatal(err)
	}

	db.Close()

	db, err = LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	newServer := testServer
	newServer.Name = "new-name"

	err = db.UpdateServer(testServer, newServer)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(newServer.String())

	db.Close()
}
