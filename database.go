package harperdb

type Database struct {
	Record
	Schema string `json:"schema"`
}

type GetBackupOptions struct {
	Table  string   `json:"table,omitempty"`
	Tables []string `json:"tables,omitempty"`
}

func (c *Client) CreateDatabase(database string) error {
	return c.opRequest(operation{
		Operation: OP_CREATE_DATABSE,
		Database:  database,
	}, nil)
}

func (c *Client) DropDatabase(database string) error {
	return c.opRequest(operation{
		Operation: OP_DROP_DATABASE,
		Database:  database,
	}, nil)
}

// func (c *Client) GetBackup(database string, options GetBackupOptions) {

// }
