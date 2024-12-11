package harperdb

type CreateAuthenticationTokensResponse struct {
	OperationToken string `json:"operation_token"`
	RefreshToken   string `json:"refresh_token"`
}

type RefreshOperationTokenResponse struct {
	OperationToken string `json:"operation_token"`
}

func (c *Client) CreateAuthenticationTokens(username, password string) (*CreateAuthenticationTokensResponse, error) {
	var response CreateAuthenticationTokensResponse
	err := c.opRequest(operation{
		Operation: OP_CREATE_AUTHENTICATION_TOKENS,
		Username:  username,
		Password:  password,
	}, &response)

	return &response, err
}

func (c *Client) RefreshOperationToken(refreshToken string) (*RefreshOperationTokenResponse, error) {
	var response RefreshOperationTokenResponse
	err := c.opRequest(operation{
		Operation:    OP_REFRESH_OPERATION_TOKEN,
		RefreshToken: refreshToken,
	}, &response)

	return &response, err
}
