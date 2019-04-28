package types

import "testing"

func TestServerDB(t *testing.T) {
	server1, err := NewServer("1.7.2", "test server 1")
	if err != nil {
		t.Fatal(err)
	}
	server2, err := NewServer("1.2.1", "test server 2")
	if err != nil {
		t.Fatal(err)
	}
	server3, err := NewServer("1.8", "test server 3")
	if err != nil {
		t.Fatal(err)
	}

	db, err := LoadDB()
	if err != nil {
		t.Fatal(err)
	}
	(*db).AddServer(server1)
	(*db).AddServer(server2)
	(*db).AddServer(server3)
	(*db).Close()

}
