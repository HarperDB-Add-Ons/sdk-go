package harperdb

type PackageComponentResponse struct {
	Project string `json:"project"`
	Payload string `json:"payload"`
}

type GetComponentsResponse struct {
	Name    string        `json:"name"`
	Entries []interface{} `json:"entries"`
}

type DeployComponentOptions struct {
	Payload string
	Package string
}

func (c *Client) AddComponent(project string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_ADD_COMPONENT,
		Project:   project,
	}, &response)

	return &response, err
}

func (c *Client) DeployComponent(project string, options DeployComponentOptions) (*MessageResponse, error) {
	var response MessageResponse
	op := operation{
		Operation: OP_DEPLOY_COMPONENT,
		Project:   project,
	}

	if len(options.Payload) > 0 {
		op.Payload = options.Package
	}
	if len(options.Package) > 0 {
		op.Package = options.Package
	}
	err := c.opRequest(op, &response)

	return &response, err
}

func (c *Client) PackageComponent(project string, skipNodeModules bool) (*PackageComponentResponse, error) {
	var response PackageComponentResponse
	err := c.opRequest(operation{
		Operation:       OP_PACKAGE_COMPONENT,
		Project:         project,
		SkipNodeModules: skipNodeModules,
	}, &response)

	return &response, err
}

func (c *Client) DropComponent(project, file string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_DROP_COMPONENT,
		Project:   project,
		File:      file,
	}, &response)

	return &response, err
}

func (c *Client) GetComponents() (*GetComponentsResponse, error) {
	var response GetComponentsResponse
	err := c.opRequest(operation{
		Operation: OP_GET_COMPONENTS,
	}, &response)

	return &response, err
}

func (c *Client) GetComponentFile(project, file string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_GET_COMPONENT_FILE,
		Project:   project,
		File:      file,
	}, &response)

	return &response, err
}

func (c *Client) SetComponentFile(project, file, payload string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_SET_COMPONENT_FILE,
		Project:   project,
		File:      file,
		Payload:   payload,
	}, &response)

	return &response, err
}
