package harperdb

import (
	"testing"
)

func TestCreateRefreshAuthenticationTokens(t *testing.T) {
	resp, err := c.CreateAuthenticationTokens(DEFAULT_USERNAME, DEFAULT_PASSWORD)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.RefreshOperationToken(resp.RefreshToken); err != nil {
		t.Fatal(err)
	}
}
