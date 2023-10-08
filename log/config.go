package log

// NewConfig return new Config after applying given options.
func NewConfig(opts ...ConfigOpt) *Config {
	var c Config
	// apply options
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

type (
	// Config required object that holds any necessary data used by each log output implementation
	Config struct {
		NR   NRConfig
		File FileConfig
	}
	// NRConfig specific config for new relic as the log output
	NRConfig struct {
		Name    string
		License string
	}
	// FileConfig specific config for file as the log output
	FileConfig struct {
		Path string
		Size int
		Age  int
		Num  int
	}
)

type ConfigOpt func(*Config)

// Output define currently supported target output for logging.
type Output int8

const (
	CONSOLE  Output = iota // CONSOLE target log output to console/terminal
	NEWRELIC               // NEWRELIC target log output directly to new relic via their client sdk
	FILE                   // FILE target log output to local file
)

// WithNRAppName set new relic application name.
func WithNRAppName(name string) ConfigOpt {
	return func(c *Config) {
		c.NR.Name = name
	}
}

// WithNRLicense set new relic license.
func WithNRLicense(lic string) ConfigOpt {
	return func(c *Config) {
		c.NR.License = lic
	}
}

// WithFilePath set target readable directory + local file which the log data
// will be written.
func WithFilePath(p string) ConfigOpt {
	return func(c *Config) {
		c.File.Path = p
	}
}

// WithFileSize set maximum file log size in megabytes before got rotated.
func WithFileSize(size int) ConfigOpt {
	return func(c *Config) {
		c.File.Size = size
	}
}

// WithFileAge set maximum number of days to retain old log files based on the
// timestamp encoded in their filename
func WithFileAge(age int) ConfigOpt {
	return func(c *Config) {
		c.File.Age = age
	}
}

// WithFileMaxBackup set maximum number of old log files to retain before got
// removed.
func WithFileMaxBackup(max int) ConfigOpt {
	return func(c *Config) {
		c.File.Num = max
	}
}
