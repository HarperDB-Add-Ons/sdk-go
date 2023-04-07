package harperdb

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	endpoint   string
	HttpClient *resty.Client
}

func NewClient(endpoint string, username string, password string) *Client {
	httpClient := resty.
		New().
		SetDisableWarn(true).
		SetBasicAuth(username, password)

	return &Client{
		endpoint:   endpoint,
		HttpClient: httpClient,
	}
}

// RawRequest allows raw requests to be made against the client.
// The recommended route for making calls is via the specific function endpoints.
func (c *Client) RawRequest(op Operation, result interface{}) error {
	return c.opRequest(op, result)
}

func (c *Client) opRequest(op Operation, result interface{}) error {
	e := ErrorResponse{}

	req := c.HttpClient.
		NewRequest().
		SetBody(op.Prepare()).
		SetError(&e)

	if result != nil {
		req.SetResult(result)
	}

	resp, err := req.Post(c.endpoint)
	if err != nil {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    err.Error()}
	}
	if resp.StatusCode() > 399 {
		return &OperationError{
			StatusCode: resp.StatusCode(),
			Message:    string(resp.Body())}
	}

	return nil
}
