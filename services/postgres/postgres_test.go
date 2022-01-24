package postgres

import (
	"os"
	"testing"
)

func TestConnection(t *testing.T) {
	os.Setenv("PostGres_Host", "localhost")
	os.Setenv("PostGres_Port", "5432")
	os.Setenv("PostGres_User", "postgres")
	os.Setenv("PostGres_Db", "impwellbeing")
	os.Setenv("PostGres_Password", "example")

	_, err := createConnection()

	if err != nil {
		t.Errorf("Got Error %s", err.Error())
	}
}

