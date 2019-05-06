package psoft

type Config struct {
	PathInventoryFile string `config:"pathInventoryFile"`
	PathBlackoutFile  string `config:"pathBlackoutFile"`
	PathExclusionFile string `config:"pathExclusionFile"`
	AttribWebMetrics  string `config:"attribWebMetrics"`
	AttribAppMetrics  string `config:"attribAppMetrics"`
	AttribPrcMetrics  string `config:"attribPrcMetrics"`
	LogLevel          string `config:"logLevel"`
	ConcurrentWorkers int    `config:"concurrentWorkers"`
	NailgunServerConn string `config:"nailgunServerConn"`
	JavaPath          string `config:"javaPath"`
}

// DefaultConfig returns default module config
func DefaultConfig() Config {
	return Config{
		LogLevel:          "INFO",
		ConcurrentWorkers: 5,
		NailgunServerConn: "local:/tmp/psmetric.socket",
	}
}
