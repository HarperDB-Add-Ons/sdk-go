package harperdb

import (
	"testing"
)

func TestSystemStatus(t *testing.T) {
	guest := randomID()

	// try with standard super user
	_, err := c.SystemInformation()
	if err != nil {
		t.Fatal(err)
	}

	// try with non-super user
	p := Permission{}
	p.SetSuperUser(false)
	p.SetClusterUser(false)
	role, err := c.AddRole(guest, p)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropRole(role.ID)

	err = c.AddUser(guest, guest, role.Role, true)
	if err != nil {
		t.Fatal(err)
	}
	defer c.DropUser(guest)

	guestClient := NewClient(DEFAULT_ENDPOINT, guest, guest)
	_, err = guestClient.SystemInformation()
	if e, ok := err.(*OperationError); ok && e.IsNotAuthorizedError() {
		return
	}

	t.Fatal("expected SystemInfo call with guest user to fail")
}

func TestRestart(t *testing.T) {
	if _, err := c.Restart(); err != nil {
		t.Fatal(err)
	}
}

func TestRestartServic(t *testing.T) {
	if _, err := c.RestartService("http_workers"); err != nil {
		t.Fatal(err)
	}
}

func TestGetConfiguration(t *testing.T) {
	if _, err := c.GetConfiguration(); err != nil {
		t.Fatal(err)
	}
}

func TestInstallNodeModule(t *testing.T) {
	if _, err := c.DeployComponent("my-project", DeployComponentOptions{Package: "HarperDB/application-template"}); err != nil {
		t.Fatal(err)
	}

	wait()

	if _, err := c.InstallNodeModules([]string{"my-project"}, false); err != nil {
		t.Fatal(err)
	}

	wait()

	if _, err := c.DropComponent("my-project", ""); err != nil {
		t.Fatal(err)
	}
}
