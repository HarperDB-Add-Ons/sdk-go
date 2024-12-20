package harperdb

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const (
	SearchBySQL   = "sql"
	SearchByHash  = "search_by_hash"
	SearchByValue = "search_by_value"
)

type ImportFromS3Response struct {
	MessageResponse
	JobId string `json:"job_id"`
}

type DeleteRecordsBeforeResponse struct {
	MessageResponse
	JobId string `json:"job_id"`
}

type SearchOperation struct {
	Operation string `json:"operation,omitempty"`
	SQL       string `json:"sql,omitempty"`
}

type S3Credentials struct {
	AWSAccessKeyID     string `json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey string `json:"aws_secret_access_key,omitempty"`
	Bucket             string `json:"bucket,omitempty"`
	Key                string `json:"filename,omitempty"`
	Region             string `json:"region,omitempty"`
}

type SysInfo struct {
	System struct {
		Platform    string `json:"platform,omitempty"`
		Distro      string `json:"distro,omitempty"`
		Release     string `json:"release,omitempty"`
		Codename    string `json:"codename,omitempty"`
		Kernel      string `json:"kernel,omitempty"`
		Arch        string `json:"arch,omitempty"`
		Hostname    string `json:"hostname,omitempty"`
		NodeVersion string `json:"node_version,omitempty"`
		NPMVersion  string `json:"npm_version,omitempty"`
	} `json:"system,omitempty"`
	Time struct {
		Current      Timestamp `json:"current,omitempty"`
		Uptime       float64   `json:"uptime,omitempty"`
		Timezone     string    `json:"timezone,omitempty"`
		TimezoneName string    `json:"timezoneName,omitempty"`
	} `json:"time,omitempty"`
	CPU struct {
		Manufacturer  string  `json:"manufacturer,omitempty"`
		Brand         string  `json:"brand,omitempty"`
		Vendor        string  `json:"vendor,omitempty"`
		Speed         float64 `json:"speed,omitempty"`
		Cores         int     `json:"cores,omitempty"`
		PhysicalCores int     `json:"physicalCores,omitempty"`
		Processors    int     `json:"processors,omitempty"`
		CPUSpeed      struct {
			Min   float64   `json:"min,omitempty"`
			Max   float64   `json:"max,omitempty"`
			Avg   float64   `json:"avg,omitempty"`
			Cores []float64 `json:"cores,omitempty"`
		} `json:"cpu_speed,omitempty"`
		CurrentLoad struct {
			AvgLoad           float64 `json:"avgload,omitempty"`
			CurrentLoad       float64 `json:"currentload,omitempty"`
			CurrentLoadUser   float64 `json:"currentload_user,omitempty"`
			CurrentLoadSystem float64 `json:"currentload_system,omitempty"`
			CurrentLoadNice   float64 `json:"currentload_nice,omitempty"`
			CurrentLoadIdle   float64 `json:"currentload_idle,omitempty"`
			CurrentLoadIRQ    float64 `json:"currentload_irq,omitempty"`
		} `json:"current_load,omitempty"`
		CPUs []CPULoad `json:"cpus,omitempty"`
	} `json:"cpu,omitempty"`
	Memory struct {
		Total     int64 `json:"total,omitempty"`
		Free      int64 `json:"free,omitempty"`
		Used      int64 `json:"used,omitempty"`
		Active    int64 `json:"active,omitempty"`
		Available int64 `json:"available,omitempty"`
		SwapTotal int64 `json:"swaptotal,omitempty"`
		SwapUsed  int64 `json:"swapused,omitempty"`
		SwapFree  int64 `json:"swapfree,omitempty"`
	} `json:"memory,omitempty"`
	Disk struct {
		IO struct {
			RIO int64 `json:"rIO,omitempty"`
			WIO int64 `json:"wIO,omitempty"`
			TIO int64 `json:"tIO,omitempty"`
		} `json:"io,omitempty"`
		ReadWrite struct {
			RX int64 `json:"rx,omitempty"`
			WX int64 `json:"wx,omitempty"`
			TX int64 `json:"tx,omitempty"`
			MS int64 `json:"ms,omitempty"`
		} `json:"read_write,omitempty"`
		Size []DiskSize `json:"size,omitempty"`
	} `json:"disk,omitempty"`
	Network struct {
		DefaultInterface string `json:"default_interface,omitempty"`
		Latency          struct {
			URL    string `json:"url,omitempty"`
			Ok     bool   `json:"ok,omitempty"`
			Status int64  `json:"status,omitempty"`
			MS     int64  `json:"ms,omitempty"`
		} `json:"latency,omitempty"`
		Interfaces  []NetworkInterface  `json:"interfaces,omitempty"`
		Stats       []NetworkStats      `json:"stats,omitempty"`
		Connections []NetworkConnection `json:"connections,omitempty"`
	} `json:"network,omitempty"`
	HarperDBProcesses struct {
		Core       []HDBProcess `json:"core,omitempty"`
		Clustering []HDBProcess `json:"clustering,omitempty"`
	} `json:"harperdb_processes,omitempty"`
	TableSize   []TableSize      `json:"table_size,omitempty"`
	Replication []NATSStreamInfo `json:"replication,omitempty"`
	Threads     []Thread         `json:"threads,omitempty"`
}

type CPULoad struct {
	Load       float64 `json:"load,omitempty"`
	LoadUser   float64 `json:"load_user,omitempty"`
	LoadSystem float64 `json:"load_system,omitempty"`
	LoadNice   float64 `json:"load_nice,omitempty"`
	LoadIdle   float64 `json:"load_idle,omitempty"`
	LoadIRQ    float64 `json:"load_irq,omitempty"`
}

type DiskSize struct {
	FS    string  `json:"fs,omitempty"`
	Type  string  `json:"overlay,omitempty"`
	Size  int64   `json:"size,omitempty"`
	Used  int64   `json:"used,omitempty"`
	Use   float64 `json:"use,omitempty"`
	Mount string  `json:"mount,omitempty"`
}

type NetworkInterface struct {
	Iface          string  `json:"iface,omitempty"`
	IfaceName      string  `json:"ifaceName,omitempty"`
	IP4            string  `json:"ip4,omitempty"`
	IP6            string  `json:"ip6,omitempty"`
	Mac            string  `json:"mac,omitempty"`
	OperState      string  `json:"operstate,omitempty"`
	Type           string  `json:"virtual,omitempty"`
	Duplex         string  `json:"duplex,omitempty"`
	Speed          float64 `json:"speed,omitempty"`
	CarrierChanges int64   `json:"carrierChanges,omitempty"`
}

type NetworkStats struct {
	Iface     string `json:"iface,omitempty"`
	OperState string `json:"operstate,omitempty"`
	RxBytes   int64  `json:"rx_bytes,omitempty"`
	RxDropped int64  `json:"rx_dropped,omitempty"`
	RxErrors  int64  `json:"rx_errors,omitempty"`
	TxBytes   int64  `json:"tx_bytes,omitempty"`
	TxDropped int64  `json:"tx_dropped,omitempty"`
	TxErrors  int64  `json:"tx_errors,omitempty"`
}

type NetworkConnection struct {
	Protocol     string `json:"protocol,omitempty"`
	LocalAddress string `json:"localaddress,omitempty"`
	LocalPort    string `json:"localport,omitempty"`
	PeerAddress  string `json:"peeraddress,omitempty"`
	PeerPort     string `json:"peerport,omitempty"`
	State        string `json:"state,omitempty"`
	PID          int64  `json:"pid,omitempty"`
	Process      string `json:"node,omitempty"`
}

type HDBProcess struct {
	PID       int64          `json:"pid,omitempty"`
	ParentPID int64          `json:"parentPid,omitempty"`
	Name      string         `json:"name,omitempty"`
	CPU       float64        `json:"cpu,omitempty"`
	CPUUser   float64        `json:"cpuu,omitempty"`
	CPUSystem float64        `json:"cpus,omitempty"`
	Memory    float64        `json:"mem,omitempty"`
	Priority  int64          `json:"priority,omitempty"`
	MemVsz    int64          `json:"memVsz,omitempty"`
	MemRSS    int64          `json:"memRss,omitempty"`
	Nice      int64          `json:"nice,omitempty"`
	Started   ProcessStarted `json:"started,omitempty"`
	State     string         `json:"state,omitempty"`
	TTY       string         `json:"tty,omitempty"`
	User      string         `json:"user,omitempty"`
	Command   string         `json:"command,omitempty"`
	Params    string         `json:"params,omitempty"`
	Path      string         `json:"path,omitempty"`
}

type ProcessStarted time.Time

// UnmarshalJSON for ProcessStarted values parses date-times that look like
// "YYYY-MM-DD HH:mm:ss" into time.Time values
func (ps *ProcessStarted) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*ps = ProcessStarted(t)
	return nil
}

type TableSize struct {
	Schema                    string `json:"schema,omitempty"`
	Table                     string `json:"table,omitempty"`
	TableSize                 int64  `json:"table_size,omitempty"`
	RecordCount               int64  `json:"record_count,omitempty"`
	TransactionLogSize        int64  `json:"transaction_log_size,omitempty"`
	TransactionLogRecordCount int64  `json:"transaction_log_record_count,omitempty"`
}

type NATSStreamInfo struct {
	StreamName string     `json:"stream_name,omitempty"`
	Database   string     `json:"database,omitempty"`
	Table      string     `json:"table,omitempty"`
	State      string     `json:"state,omitempty"`
	Consumers  []Consumer `json:"consumers,omitempty"`
}

type Consumer struct {
	Name           string    `json:"name,omitempty"`
	Created        time.Time `json:"created,omitempty"`
	NumAckPending  int64     `json:"num_ack_pending,omitempty"`
	NumRedelivered int64     `json:"num_redelivered,omitempty"`
	NumWaiting     int64     `json:"num_waiting,omitempty"`
	NumPending     int64     `json:"num_pending,omitempty"`
}

type Thread struct {
	ThreadID        int64   `json:"threadId,omitempty"`
	Name            string  `json:"name,omitempty"`
	HeapTotal       int64   `json:"heapTotal,omitempty"`
	HeapUsed        int64   `json:"heapUsed,omitempty"`
	ExternalMemory  int64   `json:"externalMemory,omitempty"`
	ArrayBuffers    int64   `json:"arrayBuffers,omitempty"`
	SinceLastUpdate int64   `json:"sinceLastUpdate,omitempty"`
	Idle            float64 `json:"idle,omitempty"`
	Active          float64 `json:"active,omitempty"`
	Utilization     float64 `json:"utilization,omitempty"`
}

func (c *Client) ExportLocal(format, path string, searchOperation SearchOperation) error {
	return c.opRequest(operation{
		Operation:       OP_EXPORT_TO_S3,
		Format:          format,
		Path:            path,
		SearchOperation: &searchOperation,
	}, nil)
}

func (c *Client) ExportToS3(format string, s3creds S3Credentials, searchOperation SearchOperation) error {
	return c.opRequest(operation{
		Operation:       OP_EXPORT_TO_S3,
		Format:          format,
		S3:              &s3creds,
		SearchOperation: &searchOperation,
	}, nil)
}

func (c *Client) ImportFromS3(action, database, table string, s3creds S3Credentials) (*ImportFromS3Response, error) {
	var result ImportFromS3Response
	err := c.opRequest(operation{
		Operation: OP_IMPORT_FROM_S3,
		Action:    action,
		Database:  database,
		Table:     table,
		S3:        &s3creds,
	}, &result)
	return &result, err
}

func (c *Client) SystemInformationAll() (*SysInfo, error) {
	var sysInfo SysInfo
	err := c.opRequest(operation{
		Operation: OP_SYSTEM_INFORMATION,
	}, &sysInfo)
	return &sysInfo, err
}

func (c *Client) SystemInformation(attrs []string) (*SysInfo, error) {
	var sysInfo SysInfo
	err := c.opRequest(operation{
		Operation:  OP_SYSTEM_INFORMATION,
		Attributes: attrs,
	}, &sysInfo)
	return &sysInfo, err
}

func (c *Client) Restart() (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_RESTART,
	}, &response)

	return &response, err
}

func (c *Client) RestartService(service string) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_RESTART_SERVICE,
		Service:   service,
	}, &response)

	return &response, err
}

func (c *Client) DeleteRecordsBefore(date time.Time, schema, table string) (*DeleteRecordsBeforeResponse, error) {
	var response DeleteRecordsBeforeResponse

	err := c.opRequest(operation{
		Operation: OP_DELETE_RECORDS_BEFORE,
		Date:      date.UTC().Format(time.RFC3339),
		Schema:    schema,
		Table:     table,
	}, &response)

	return &response, err
}

func (c *Client) InstallNodeModules(projects []string, dryRun bool) (*MessageResponse, error) {
	var response MessageResponse
	err := c.opRequest(operation{
		Operation: OP_INSTALL_NODE_MODULES,
		Projects:  projects,
		DryRun:    dryRun,
	}, &response)

	return &response, err
}

func (c *Client) GetConfiguration() (map[string]interface{}, error) {
	var response map[string]interface{}
	err := c.opRequest(operation{
		Operation: OP_GET_CONFIGURATION,
	}, &response)

	return response, err
}

func (c *Client) SetConfiguration(configuration interface{}) (*MessageResponse, error) {
	var response MessageResponse
	data, err := json.Marshal(configuration)
	if err != nil {
		return &MessageResponse{"unable to marshal struct into json"}, errors.New("unable to marshal struct into json")
	}
	v2 := map[string]interface{}{}
	if err := json.Unmarshal(data, &v2); err != nil {
		return &MessageResponse{"unable to unmarhsal struct into map"}, errors.New("unable to unmarhsal struct into map")
	}
	v2["operation"] = "set_configuration"
	err = c.SetConfigurationRequest(v2, &response)

	return &response, err
}
