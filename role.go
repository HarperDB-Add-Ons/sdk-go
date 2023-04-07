package harperdb

type SchemaPermission struct {
	Tables map[string]TablePermission `json:"tables"`
}

type AttributePermissions struct {
	AttributeName string `json:"attribute_name"`
	Read          bool   `json:"read"`
	Insert        bool   `json:"insert"`
	Update        bool   `json:"update"`
}

type TablePermission struct {
	Read                 bool                   `json:"read"`
	Insert               bool                   `json:"insert"`
	Update               bool                   `json:"update"`
	Delete               bool                   `json:"delete"`
	AttributePermissions []AttributePermissions `json:"attribute_permissions"`
}

type Permission map[string]interface{}

func (p Permission) SetSuperUser(su bool) {
	p["super_user"] = su
}

func (p Permission) SetClusterUser(su bool) {
	p["cluster_user"] = su
}

func (p Permission) AddSchemaPermission(schema string, sp SchemaPermission) {
	p[schema] = sp
}

func (p SchemaPermission) AddTablePermission(table string, tp TablePermission) {
	p.Tables[table] = tp
}

type Role struct {
	Record
	Role       string     `json:"role"`
	ID         string     `json:"id"`
	Permission Permission `json:"permission"`
}

// TODO Correct JSON modelling needs to be verified
func (c *Client) ListRoles() ([]Role, error) {
	roles := []Role{}
	err := c.opRequest(operation{
		Operation: OP_LIST_ROLES,
	}, &roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (c *Client) AddRole(role string, perm Permission) (*Role, error) {
	var newRole Role
	req := OpAddRole{
		Role:       role,
		Permission: perm,
	}
	err := c.opRequest(req, &newRole)
	if err != nil {
		return nil, err
	}

	return &newRole, nil
}

func (c *Client) DropRole(id string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_ROLE,
		ID:        id,
	}, nil)
}

func (c *Client) AlterRole(id string, role string, perm Permission) (*Role, error) {
	var newRole Role
	req := OpAlterRole{
		ID:         id,
		Role:       role,
		Permission: perm,
	}
	err := c.opRequest(req, &newRole)
	if err != nil {
		return nil, err
	}

	return &newRole, nil
}
