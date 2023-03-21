package harperdb

type Route struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type SetRouteResponse struct {
	Message string  `json:"message"`
	Set     []Route `json:"set"`
	Skipped []Route `json:"skipped"`
}

type GetRouteResponse struct {
	Hub  []Route `json:"hub"`
	Leaf []Route `json:"route"`
}

type DeleteRouteResponse struct {
	Message string  `json:"string"`
	Deleted []Route `json:"deleted"`
	Skipped []Route `json:"skipped"`
}

func (c *Client) SetRoutes(op OpSetRoutes) (response SetRouteResponse, err error) {
	err = c.opRequest(op, &response)
	return response, err
}

func (c *Client) GetRoutes() (response GetRouteResponse, err error) {
	err = c.opRequest(OpGetRoutes{}, &response)
	return response, err
}

func (c *Client) DeleteRoutes(routes []Route) (response DeleteRouteResponse, err error) {
	if len(routes) == 0 {
		return DeleteRouteResponse{}, &OperationError{Message: "must supply at least one route to delete"}
	}
	op := OpDeleteRoutes{
		Routes: routes,
	}
	err = c.opRequest(op, &response)
	return response, err
}
