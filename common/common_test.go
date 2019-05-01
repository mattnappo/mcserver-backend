package common

import "testing"

func TestNewServerPath(t *testing.T) {
	name := "test"
	path, err := NewServerPath(name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
}
