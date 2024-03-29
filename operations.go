package harperdb

import "time"

const (
	OP_ADD_NODE               = "add_node"
	OP_ADD_ROLE               = "add_role"
	OP_ADD_USER               = "add_user"
	OP_ALTER_ROLE             = "alter_role"
	OP_ALTER_USER             = "alter_user"
	OP_CLUSTER_SET_ROUTES     = "cluster_set_routes"
	OP_CLUSTER_GET_ROUTES     = "cluster_get_routes"
	OP_CLUSTER_DELETE_ROUTES  = "cluster_delete_routes"
	OP_CLUSTER_STATUS         = "cluster_status"
	OP_CREATE_ATTRIBUTE       = "create_attribute"
	OP_CREATE_SCHEMA          = "create_schema"
	OP_CREATE_TABLE           = "create_table"
	OP_CSV_DATA_LOAD          = "csv_data_load"
	OP_CSV_FILE_LOAD          = "csv_file_load"
	OP_CSV_URL_LOAD           = "csv_url_load"
	OP_DELETE_FILES_BEFORE    = "delete_files_before"
	OP_DELETE_TRANSACTION_LOG = "delete_transaction_logs_before"
	OP_DESCRIBE_ALL           = "describe_all"
	OP_DESCRIBE_SCHEMA        = "describe_schema"
	OP_DESCRIBE_TABLE         = "describe_table"
	OP_DELETE                 = "delete"
	OP_DROP_ATTRIBUTE         = "drop_attribute"
	OP_DROP_ROLE              = "drop_role"
	OP_DROP_SCHEMA            = "drop_schema"
	OP_DROP_TABLE             = "drop_table"
	OP_DROP_USER              = "drop_user"
	OP_EXPORT_LOCAL           = "export_local"
	OP_EXPORT_TO_S3           = "export_to_s3"
	OP_GET_FINGERPRINT        = "get_fingerprint"
	OP_GET_JOB                = "get_job"
	OP_INSERT                 = "insert"
	OP_LIST_ROLES             = "list_roles"
	OP_LIST_USERS             = "list_users"
	OP_READ_LOG               = "read_log"
	OP_READ_TRANSACTION_LOG   = "read_transaction_log"
	OP_REGISTRATION_INFO      = "registration_info"
	OP_REMOVE_NODE            = "remove_node"
	OP_SEARCH_BY_HASH         = "search_by_hash"
	OP_SEARCH_BY_VALUE        = "search_by_value"
	OP_SEARCH_JOBS            = "search_jobs_by_start_date"
	OP_SET_LICENSE            = "set_license"
	OP_SQL                    = "sql"
	OP_SYSTEM_INFORMATION     = "system_information"
	OP_UPDATE                 = "update"
	OP_UPDATE_NODE            = "update_node"
	OP_USER_INFO              = "user_info"
)

type Operation interface {
	Prepare() interface{}
}

type OpAddRole struct {
	Permission Permission `json:"permission"`
	Role       string     `json:"role"`
}

func (o OpAddRole) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAddRole
	}
	return Return{
		Operation: OP_ADD_ROLE,
		OpAddRole: o,
	}
}

type OpAlterRole struct {
	ID         string     `json:"id"`
	Role       string     `json:"role"`
	Permission Permission `json:"permission"`
}

func (o OpAlterRole) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAlterRole
	}
	return Return{
		Operation:   OP_ALTER_ROLE,
		OpAlterRole: o,
	}
}

// Describe Schema
type OpDescribeSchema struct {
	Schema string `json:"schema"`
}

func (o OpDescribeSchema) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpDescribeSchema
	}
	return Return{
		Operation:        OP_DESCRIBE_SCHEMA,
		OpDescribeSchema: o,
	}
}

// Set Routes
type OpSetRoutes struct {
	Server string  `json:"server"` // Must be either "hub" or "leaf"
	Routes []Route `json:"routes"`
}

func (o OpSetRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpSetRoutes
	}

	return Return{
		Operation:   OP_CLUSTER_SET_ROUTES,
		OpSetRoutes: o,
	}
}

// Delete Routes
type OpDeleteRoutes struct {
	Routes []Route `json:"routes"`
}

func (o OpDeleteRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpDeleteRoutes
	}

	return Return{
		Operation:      OP_CLUSTER_DELETE_ROUTES,
		OpDeleteRoutes: o,
	}
}

// Get Routes
type OpGetRoutes struct{}

func (o OpGetRoutes) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
	}

	return Return{
		Operation: OP_CLUSTER_GET_ROUTES,
	}
}

// Add Node
type OpAddNode struct {
	NodeName      string         `json:"node_name"`
	Host          string         `json:"host"`
	Port          int            `json:"port"`
	Subscriptions []Subscription `json:"subscriptions"`
}

func (o OpAddNode) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpAddNode
	}

	return Return{
		Operation: OP_ADD_NODE,
		OpAddNode: o,
	}
}

// Remove Node
type OpRemoveNode struct {
	NodeName string `json:"node_name"`
}

func (o OpRemoveNode) Prepare() interface{} {
	type Return struct {
		Operation string `json:"operation"`
		OpRemoveNode
	}

	return Return{
		Operation:    OP_REMOVE_NODE,
		OpRemoveNode: o,
	}
}

type operation struct {
	Action          string          `json:"action,omitempty"`
	Active          *bool           `json:"active,omitempty"`
	Attribute       string          `json:"attribute,omitempty"`
	Company         string          `json:"company,omitempty"`
	CSVURL          string          `json:"csv_url,omitempty"`
	Data            string          `json:"data,omitempty"`
	Date            time.Time       `json:"date,omitempty"`
	FilePath        string          `json:"file_path,omitempty"`
	Format          string          `json:"format,omitempty"`
	From            string          `json:"from,omitempty"`
	FromDate        string          `json:"from_date,omitempty"`
	GetAttributes   AttributeList   `json:"get_attributes,omitempty"`
	HashAttribute   string          `json:"hash_attribute,omitempty"`
	HashValues      interface{}     `json:"hash_values,omitempty"`
	Host            string          `json:"host,omitempty"`
	ID              string          `json:"id,omitempty"`
	Key             string          `json:"key,omitempty"`
	Limit           int             `json:"limit,omitempty"`
	Name            string          `json:"name,omitempty"`
	Operation       string          `json:"operation"`
	Order           string          `json:"order,omitempty"`
	Password        string          `json:"password,omitempty"`
	Path            string          `json:"path,omitempty"`
	Permission      Permission      `json:"permission,omitempty"`
	Port            int             `json:"port,omitempty"`
	Records         interface{}     `json:"records,omitempty"`
	Role            string          `json:"role,omitempty"`
	S3              S3Credentials   `json:"s3,omitempty"`
	Schema          string          `json:"schema,omitempty"`
	SearchAttribute Attribute       `json:"search_attribute,omitempty"`
	SearchOperation SearchOperation `json:"search_operation,omitempty"`
	SearchType      string          `json:"search_type,omitempty"`
	SearchValue     interface{}     `json:"search_value,omitempty"`
	SearchValues    interface{}     `json:"search_values,omitempty"`
	Start           int             `json:"start,omitempty"`
	Subscriptions   []Subscription  `json:"subscriptions,omitempty"`
	SQL             string          `json:"sql,omitempty"`
	Table           string          `json:"table,omitempty"`
	Timestamp       int64           `json:"timestamp,omitempty"`
	ToDate          string          `json:"to_date,omitempty"`
	Until           string          `json:"until,omitempty"`
	Username        string          `json:"username,omitempty"`
}

func (o operation) Prepare() interface{} {
	return o
}
