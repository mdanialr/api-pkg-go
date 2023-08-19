package logger

import "github.com/spf13/viper"

// Config object that holds any necessary data needed by logger pkg.
type Config struct {
	Cnf *viper.Viper
	// NRApp used by newrelic when use newrelic as one of the log output as app name.
	NRApp string
	// NRLicense used by newrelic when use newrelic as one of the log output as app license.
	NRLicense string
}
