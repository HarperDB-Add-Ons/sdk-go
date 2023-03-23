package harperdb

import (
	"fmt"
	"testing"
)

const (
	CLUSTER_USER = "cluster_user"
	SUPER_USER   = "super_user"
)

func TestListRoles(t *testing.T) {
	roles, err := c.ListRoles()
	if err != nil {
		t.Fatal(err)
	}

	// We should find two roles, super_user and cluster_user
	var foundSU, foundCU bool
	for _, role := range roles {
		switch role.Role {
		case CLUSTER_USER:
			foundCU = true
		case SUPER_USER:
			foundSU = true
		}
	}

	if !(foundSU && foundCU) {
		t.Fatalf("did not find super_user(%t) or cluster_user(%t) role", foundSU, foundCU)
	}
}

func TestAddAndDropRole(t *testing.T) {
	roleName := randomID()
	perms := Permission{}
	perms.SetClusterUser(false)
	perms.SetSuperUser(false)
	r, err := c.AddRole(roleName, perms)
	if err != nil {
		t.Fatal(err)
	}

	if err := c.DropRole(r.ID); err != nil {
		t.Fatal(err)
	}
}

func TestAlterRole(t *testing.T) {
	foundCU, err := findRole(CLUSTER_USER)
	if err != nil {
		t.Fatal(err)
	}

	if foundCU == nil {
		t.Fatal("did not find cluster user role")
	}

	newName := foundCU.Role
	newRole, err := c.AlterRole(foundCU.ID, newName, foundCU.Permission)
	if err != nil {
		t.Fatal(err)
	}
	if newRole.Role != newName {
		t.Fatal(fmt.Errorf("expected altered role to have name %s", newName))
	}

	// Change name back to original
	_, err = c.AlterRole(foundCU.ID, foundCU.Role, foundCU.Permission)
	if err != nil {
		t.Fatal(err)
	}
}

func findRole(name string) (*Role, error) {
	roles, err := c.ListRoles()
	if err != nil {
		return nil, err
	}

	var found *Role
	for _, role := range roles {
		if role.Role == name {
			found = &role
			break
		}
	}
	return found, nil
}
