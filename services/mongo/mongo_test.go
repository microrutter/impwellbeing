package mongo

import (
	"os"
    "testing"
)

func TestConnection(t *testing.T) {
	os.Setenv("MongoDB", "mongodb://localhost:27017")

	connection := CheckConnection()
	expected_return := "Connected to MongoDB!"

	if connection != expected_return {
		t.Errorf("got %s, wanted %s", connection, expected_return)
	}
}