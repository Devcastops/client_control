package config

type Config struct {
	Packer  Packer  `json:"packer"`
	GCP     GCP     `json:"gcp"`
	Nomad   Nomad   `json:"nomad"`
	Webhook Webhook `json:"webhook"`
}

type Packer struct {
	OrganizationID string `json:"organizationID"`
	ProjectID      string `json:"projectID"`
	BucketName     string `json:"bucketName"`
}

type GCP struct {
	Project string     `json:"project"`
	Compute GCPCompute `json:"compute"`
}
type GCPCompute struct {
	Zone           string `json:"zone"`
	ServiceAccount string `json:"serviceAccount"`
	Subnetwork     string `json:"subnetwork"`
	DiskSize       int    `json:"diskSize"`
}
type Nomad struct {
	ServerIPs []string `json:"serverIPs"`
}
type Webhook struct{}
