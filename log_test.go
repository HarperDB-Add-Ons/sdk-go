package harperdb

import (
	"testing"
	"time"
)

func TestReadDeleteAuditLog(t *testing.T) {
	database := randomID()
	table := randomID()

	if err := c.CreateDatabase(database); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.CreateTable(database, table, "id"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	record := createTestRecord()

	if _, err := c.Insert(database, table, []interface{}{record}); err != nil {
		t.Fatal(err)
	}

	wait()

	resp, err := c.ReadAuditLog(database, table, "", nil)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)

	if len(resp) == 0 {
		t.Fatal("There should be at least one record in the audit log")
	}

	if message, err := c.DeleteAuditLogsBefore(database, table, time.Now()); err != nil {
		t.Log(message)
		t.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	resp, err = c.ReadAuditLog(database, table, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if len(resp) != 0 {
		t.Fatal("There should be at no records in the audit log")
	}

	if err := c.DropSchema(database); err != nil {
		t.Fatal(err)
	}
}
