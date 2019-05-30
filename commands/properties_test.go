package commands

import "testing"

func TestDefaultProperties(t *testing.T) {
	properties := DefaultProperties(7071, "cool test server!")

	t.Log(properties)
}

func TestGetFile(t *testing.T) {
	properties := DefaultProperties(7071, "cool test server!")

	data := properties.GetFile()

	t.Log(data)
}

func TestChangeProperty(t *testing.T) {
	properties := DefaultProperties(7071, "cool test server!")

	data := properties.GetFile()
	t.Log(data)

	properties.ChangeProperty("LevelType", "new value")

	data = properties.GetFile()
	t.Log(data)
}
