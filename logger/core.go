package logger

import "github.com/spf13/viper"

// Config object that holds any necessary data needed by logger pkg.
type Config struct {
	Cnf *viper.Viper
	// NRApp used by newrelic when use newrelic as one of the log output as app name.
	NRApp string
	// NRLicense used by newrelic when use newrelic as one of the log output as app license.
	NRLicense string

	ConsoleLevel string // ConsoleLevel log level for console output.
	FileLevel    string // FileLevel log level for file output.
	NRLevel      string // NRLevel log level for newrelic output.
}
