package psoft

import (
        "os"
)

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
	ConcatenateDomainWithHost bool `config:"concatDomainHost"`
	UseLastXCharactersOfHost int `config:"useLastXCharsOfHost"`
	LocalInventoryOnly bool `config:"localInventoryOnly"`

}

// DefaultConfig returns default module config
func DefaultConfig() Config {
	wd, _ := os.Getwd()
	return Config{
		LogLevel:          "INFO",
		ConcurrentWorkers: 5,
		NailgunServerConn: "local:" + wd + "/run/psmetric.socket",
		ConcatenateDomainWithHost: false,
		UseLastXCharactersOfHost: 0,
		LocalInventoryOnly: false,
		AttribWebMetrics: "web_metric.yaml",
		AttribAppMetrics: "app_metric.yaml",
		AttribPrcMetrics: "prc_metric.yaml",
	}
}
