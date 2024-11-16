package gcp

type Client struct {
	Project string
}

type StartInstance struct {
	Client         *Client
	Zone           string
	StackType      string
	MachineType    string
	Subnetwork     string
	ServiceAccount string
	DiskSize       int64
	Scopes         []string
	Tags           []string
	Lables         map[string]string
	Metadata       map[string]string
}

type StopInstance struct {
	Client *Client
	Zone   string
}

type GetInstance struct {
	Client *Client
	Zone   string
}
