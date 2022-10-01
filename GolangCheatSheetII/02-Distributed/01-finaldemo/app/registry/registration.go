package registry

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	HeartbeatURL     string
	ServiceUpdateURL string
	RequiredServices []ServiceName
}
