package database

import (
	"context"
	"testing"

	"github.com/jimlawless/whereami"
	"github.com/nouney/randomstring"

	"github.com/zozoee27/cookbook/backend/testutil"
)

func TestInitializeConnection(t *testing.T) {
	connection := &Connection{}
	dbName := randomstring.Generate(5)
	connection.InitializeConnection(dbName)

	err := connection.DatabaseClient.Ping(context.TODO(), nil)

	if err != nil {
		t.Logf("Cannot ping database: %s", err.Error())
	}

	testutil.CompareString(t, connection.Database.Name(), dbName, "Database names are different", whereami.WhereAmI())

	connection.Disconnect()
}
