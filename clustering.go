package harperdb

type ClusterNetworkResponse struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Name          string   `json:"name"`
	ResponseTime  int      `json:"response_time"`
	ConnectedNoes []string `json:"connected_nodes"`
	Routes        []Route  `json:"routes"`
}

type ConfigureClusterConnection struct {
	NodeName      string         `json:"node_name"`
	Subscriptions []Subscription `json:"subscriptions"`
}

type PurgeStreamOptions struct {
	Keep string `json:"keep,omitempty"`
	Seq  string `json:"seq,omitempty"`
}

func (c *Client) SetNodeReplication(nodeName string, subscriptions []Subscription) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation:     OP_SET_NODE_REPLICATION,
		NodeName:      nodeName,
		Subscriptions: subscriptions,
	}, &response)

	return &response, err
}

func (c *Client) ClusterNetwork() (*ClusterNetworkResponse, error) {
	var response ClusterNetworkResponse
	err := c.opRequest(operation{
		Operation: OP_CLUSTER_NETWORK,
	}, &response)

	return &response, err
}

func (c *Client) ConfigureCluster(connections []ConfigureClusterConnection) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation:   OP_CONFIGURE_CLUSTER,
		Connections: connections,
	}, &response)

	return &response, err
}

func (c *Client) PurgeStream(database, table string, options PurgeStreamOptions) error {
	return c.opRequest(operation{
		Operation: OP_PURGE_STREAM,
		Database:  database,
		Table:     table,
		Options:   &options,
	}, nil)
}
