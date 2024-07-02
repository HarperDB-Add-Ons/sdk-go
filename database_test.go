package harperdb

import (
	"testing"
)

func TestCreateDropDatabase(t *testing.T) {
	database := randomID()

	if err := c.CreateDatabase(database); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.DropDatabase(database); err != nil {
		t.Fatal(err)
	}
}

func TestCreateDuplicateDatabase(t *testing.T) {
	database := randomID()
	if err := c.CreateDatabase(database); err != nil {
		t.Fatal(err)
	}

	wait()

	err := c.CreateDatabase(database)
	if e, ok := err.(*OperationError); ok && e.IsAlreadyExistsError() {
		return
	} else {
		t.Log(e)
		t.Fatalf("should have raised AlreadyExistsError")
	}

	wait()

	if err := c.DropDatabase(database); err != nil {
		t.Fatal(err)
	}
}
