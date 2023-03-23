package harperdb

import (
	"testing"
)

func TestSetRoutes(t *testing.T) {
	t.Skip("Route API not implemented")
	routes := []Route{
		{
			Host: "127.0.0.1",
			Port: 1112,
		},
	}
	req := OpSetRoutes{
		Server: "hub",
		Routes: routes,
	}
	resp, err := c.SetRoutes(req)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Skipped) > 0 {
		t.Fatalf("Skipped routes is %d, should be 0", len(resp.Skipped))
	}
}

func TestGetRoutes(t *testing.T) {
	_, err := c.GetRoutes()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteRoutes(t *testing.T) {
	t.Skip("Not implemented")
}
