package harperdb

import (
	"fmt"
	"testing"
)

func TestListUsers(t *testing.T) {
	users, err := c.ListUsers()
	if err != nil {
		t.Fatal(err)
	}

	// We expect at least to have the HDB_ADMIN
	var found bool
	for _, user := range users {
		if user.Username == DEFAULT_USERNAME {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("expected to find HDB_ADMIN user")
	}
}

func TestAddUser(t *testing.T) {
	superUser, err := findRole(SUPER_USER)
	if err != nil {
		t.Fatal(err)
	}

	testUser := randomID()
	err = c.AddUser(testUser, randomID(), superUser.Role, true)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropUser(testUser)

	user, err := findUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal(fmt.Sprintf("expected to find user %s", testUser))
	}
}

func TestAlterUser(t *testing.T) {
	superUser, err := findRole(SUPER_USER)
	if err != nil {
		t.Fatal(err)
	}

	testUser := randomID()
	err = c.AddUser(testUser, randomID(), superUser.Role, true)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropUser(testUser)

	// set user to inactive
	err = c.AlterUser(testUser, randomID(), superUser.Role, false)
	if err != nil {
		t.Fatal(err)
	}

	user, err := findUser(testUser)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal(fmt.Sprintf("expected to find user %s", testUser))
	}
	if user.Active {
		t.Fatal("did not expect user to be active")
	}
}

func TestUserInfo(t *testing.T) {
	user, err := c.UserInfo()
	if err != nil {
		t.Fatal(err)
	}

	if user.Username != DEFAULT_USERNAME {
		t.Fatal(fmt.Errorf("expected user to be %s: %s", DEFAULT_USERNAME, user.Username))
	}
}

func findUser(username string) (*User, error) {
	users, err := c.ListUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, nil
}
