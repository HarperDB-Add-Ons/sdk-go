package harperdb

import (
	"testing"
)

func TestAddAndRemoveNode(t *testing.T) {
	t.Skip("Test not currently working")
	// Try to remove node first in case previous tests failed
	nodeName := "TEST_NODE_NAME"

	err := c.AddNode(nodeName, "127.0.0.1", 1112, []Subscription{
		{
			Channel:   "dev:dog",
			Publish:   true,
			Subscribe: false,
			Schema:    "invalid",
			Table:     "invalid",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.UpdateNode(nodeName, "127.0.0.1", 1112, []Subscription{
		{
			Channel:   "dev:dog",
			Publish:   true,
			Subscribe: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	err = c.RemoveNode(nodeName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClusterStatus(t *testing.T) {
	_, err := c.ClusterStatus()
	if err != nil {
		t.Fatal(err)
	}

	// TODO Here we could do a little more testing
	// with an active cluster
}
