package harperdb

type CustomFunctionStatusResponse struct {
	IsEnabled bool   `json:"is_enabled"`
	Port      int    `json:"port"`
	Directory string `json:"directory"`
}

type PackageCustomFunctionProjectResponse struct {
	Project string `json:"project"`
	Payload string `json:"payload"`
	File    string `json:"file"`
}

func (c *Client) CustomFunctionStatus() (*CustomFunctionStatusResponse, error) {
	var response CustomFunctionStatusResponse
	err := c.opRequest(operation{
		Operation: OP_CUSTOM_FUNCTIONS_STATUS,
	}, &response)

	return &response, err
}

func (c *Client) GetCustomFunctions() (map[string]interface{}, error) {
	var response map[string]interface{}
	err := c.opRequest(operation{
		Operation: OP_GET_CUSTOM_FUNCTIONS,
	}, &response)

	return response, err
}

func (c *Client) GetCustomFunction(project, type_, file string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_GET_CUSTOM_FUNCTION,
		Project:   project,
		Type:      type_,
		File:      file,
	}, &response)

	return &response, err
}

func (c *Client) SetCustomFunction(project, type_, file, functionContent string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation:       OP_SET_CUSTOM_FUNCTION,
		Project:         project,
		Type:            type_,
		File:            file,
		FunctionContent: functionContent,
	}, &response)

	return &response, err
}

func (c *Client) DropCustomFunction(project, type_, file string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_DROP_CUSTOM_FUNCTION,
		Project:   project,
		Type:      type_,
		File:      file,
	}, &response)

	return &response, err
}

func (c *Client) AddCustomFunctionProject(project string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_ADD_CUSTOM_FUNCTION_PROJECT,
		Project:   project,
	}, &response)

	return &response, err
}

func (c *Client) DropCustomFunctionProject(project string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_DROP_CUSTOM_FUNCTION_PROJECT,
		Project:   project,
	}, &response)

	return &response, err
}

func (c *Client) PackageCustomFunctionProject(project string, skipNodeModules bool) (*PackageCustomFunctionProjectResponse, error) {
	var response PackageCustomFunctionProjectResponse
	err := c.opRequest(operation{
		Operation:       OP_PACKAGE_CUSTOM_FUNCTION_PROJECT,
		Project:         project,
		SkipNodeModules: skipNodeModules,
	}, &response)

	return &response, err
}
