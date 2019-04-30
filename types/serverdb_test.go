package types

import "testing"

func TestLoadDB(t *testing.T) {
	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(db.String())
}

func TestAddServer(t *testing.T) {
	version := "1.12"
	name := "my-test-server"
	port := 25565
	server, err := NewServer(version, name, port)
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
}

func TestClose(t *testing.T) {
	version := "1.12"
	name := "my-test-server"
	port := 25565
	server, err := NewServer(version, name, port)
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
