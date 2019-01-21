package smpl

// Server represents the entire server configuration and lists the sites.
// Once smpl supports live config changes, this will need protection via mutex.
type Server struct {
	Sites []Site
}

// These are the environments available for certificates usage.
const (
	EnvDev   = "dev"
	EnvStage = "stage"
	EnvProd  = "prod"
)
