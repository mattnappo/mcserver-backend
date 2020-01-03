package types

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func getTestServer() (*Server, error) {
	rand.Seed(time.Now().UnixNano())
	min := 25565
	max := 26000
	random := rand.Intn(max-min) + min

	server, err := NewServer("1.7.2", "server-"+strconv.Itoa(random), random, 1024)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func TestLoadDB(t *testing.T) {
	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(db.String())
}

func TestAddServer(t *testing.T) {
	testServer, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[ before AddServer ] %s\n", db.String())

	err = db.AddServer(testServer)
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
	testServer, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("[ before AddServer ] %s\n", db.String())

	db.AddServer(testServer)

	t.Logf("[ after AddServer ] %s\n", db.String())

	err = db.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSever(t *testing.T) {
	testServer, err := getTestServer()
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

func TestDeleteServer(t *testing.T) {
	testServer, err := getTestServer()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(testServer.String())
	t.Log(testServer.Hash.String())

	/* -----------call one---------------- */
	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	err = db.AddServer(testServer)
	if err != nil {
		t.Fatal(err)
	}

	db.Close()
	/* ------------------------------ */

	/* -----------call two---------------- */
	db, err = LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	deletedServer, err := db.DeleteServer(testServer.Hash.String())
	if err != nil {
		t.Fatal(err)
	}

	db.Close()
	t.Log(deletedServer.String())
	/* ------------------------------ */

	/* -----------call three---------------- */
	db, err = LoadDB()
	if err != nil {
		t.Fatal(err)
	}

	for i, server := range db.Servers {
		t.Logf("SERVER %d:\n%s\n\n", i, server.String())
	}

	db.Close()
	/* ------------------------------ */

}
