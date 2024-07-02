package harperdb

import (
	"fmt"
	"testing"
)

const payload = "console.log('hello world')"

func TestAddDropComponent(t *testing.T) {
	project := randomID()
	if _, err := c.AddComponent(project); err != nil {
		t.Fatal(err)
	}

	if _, err := c.DropComponent(project, ""); err != nil {
		t.Fatal(err)
	}
}

func TestGetSetComponents(t *testing.T) {
	project := randomID()
	if _, err := c.AddComponent(project); err != nil {
		t.Fatal(err)
	}

	if _, err := c.SetComponentFile(project, "test.js", payload); err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetComponentFile(project, "test.js")

	if err != nil {
		t.Fatal(err)
	}
	if resp.Message != payload {
		t.Fatal(fmt.Errorf("Expected file contents to match defined const"))
	}

	if _, err := c.GetComponents(); err != nil {
		t.Fatal(err)
	}
}

func TestPackageComponent(t *testing.T) {
	project := randomID()
	if _, err := c.AddComponent(project); err != nil {
		t.Fatal(err)
	}

	if _, err := c.SetComponentFile(project, "test.js", payload); err != nil {
		t.Fatal(err)
	}

	if _, err := c.PackageComponent(project, false); err != nil {
		t.Fatal(err)
	}
}

func TestDeployComponent(t *testing.T) {
	if _, err := c.DeployComponent("my-project", DeployComponentOptions{Package: "HarperDB/application-template"}); err != nil {
		t.Fatal(err)
	}

	if _, err := c.DropComponent("my-project", ""); err != nil {
		t.Fatal(err)
	}
}
