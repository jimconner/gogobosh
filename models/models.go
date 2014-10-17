package models

// Director is a targeted BOSH director and login credentials
type Director struct {
	TargetURL string
	Username  string
	Password  string
}

// DirectorInfo contains the status of a target Director
type DirectorInfo struct {
	Name                         string
	URL                          string
	Version                      string
	User                         string
	UUID                         string
	CPI                          string
	DNSEnabled                   bool
	DNSDomainName                string
	CompiledPackageCacheEnabled  bool
	CompiledPackageCacheProvider string
	SnapshotsEnabled             bool
}

// Stemcell describes an available versioned stemcell
type Stemcell struct {
	Name    string
	Version string
	Cid     string
}

// Release describes a release and all available versions
type Release struct {
	Name     string
	Versions []ReleaseVersion
}

// ReleaseVersion describes an available versioned release
type ReleaseVersion struct {
	Version            string
	CommitHash         string
	UncommittedChanges bool
	CurrentlyDeployed  bool
}

// Deployment describes a running BOSH deployment and the
// Releases and Stemcells it is using.
type Deployment struct {
	Name      string
	Releases  []NameVersion
	Stemcells []NameVersion
}

// DeploymentManifest describes all the configuration for any BOSH deployment
type DeploymentManifest struct {
	Meta          map[string]interface{}
	Name          string
	DirectorUUID  string `yaml:"director_uuid"`
	Releases      []*NameVersion
	Compilation   *manifestCompilation
	Update        *manifestUpdate
	Networks      []*manifestNetwork
	ResourcePools []*manifestResourcePool `yaml:"resource_pools"`
	Jobs          []*manifestJob
	Properties    *map[string]interface{}
}

type manifestCompilation struct {
	Workers             int                     `yaml:"workers"`
	NetworkName         string                  `yaml:"network"`
	ReuseCompilationVMs bool                    `yaml:"reuse_compilation_vms"`
	CloudProperties     *map[string]interface{} `yaml:"cloud_properties"`
}

type manifestUpdate struct {
	Canaries        int
	MaxInFlight     int    `yaml:"max_in_flight"`
	CanaryWatchTime string `yaml:"canary_watch_time"`
	UpdateWatchTime string `yaml:"update_watch_time"`
	Serial          bool
}

type manifestNetwork struct {
	Name            string
	Type            string
	CloudProperties *map[string]interface{} `yaml:"cloud_properties"`
	Subnets         interface{}
}

type manifestResourcePool struct {
	Name            string
	NetworkName     string `yaml:"network"`
	Stemcell        string
	CloudProperties string `yaml:"cloud_properties"`
}

type manifestJob struct {
	Name             string
	Templates        []*manifestJobTemplate
	Instances        int             `yaml:"instances,omitempty"`
	ResourcePoolName string          `yaml:"resource_pool"`
	PersistentDisk   int             `yaml:"persistent_disk,omitempty"`
	Lifecycle        string          `yaml:"lifecycle,omitempty"`
	Update           *manifestUpdate `yaml:"update"`
	Networks         []*manifestJobNetwork
	Properties       *map[string]interface{} `yaml:"properties"`
}

type manifestJobTemplate struct {
	Name    string
	Release string
}

type manifestJobNetwork struct {
	Name      string
	Default   *[]string `yaml:"default"`
	StaticIPs *[]string `yaml:"static_ips"`
}

// DeploymentVM describes the association of a running server
// within a Deployment
type DeploymentVM struct {
	JobName string
	Index   int
	VMCid   string
	AgentID string
}

// NameVersion is a reusable structure for Name/Version information
type NameVersion struct {
	Name    string
	Version string
}

// TaskStatus summarizes the current status of a Task
type TaskStatus struct {
	ID          int
	State       string
	Description string
	TimeStamp   int
	Result      string
	User        string
}

// VMStatus summarizes the current status of a VM/server
// within a running deployment
type VMStatus struct {
	JobName               string
	Index                 int
	JobState              string
	VMCid                 string
	AgentID               string
	ResourcePool          string
	ResurrectionPaused    bool
	IPs                   []string
	DNSs                  []string
	CPUUser               float64
	CPUSys                float64
	CPUWait               float64
	MemoryPercent         float64
	MemoryKb              int
	SwapPercent           float64
	SwapKb                int
	DiskPersistentPercent float64
}
