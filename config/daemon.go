package config

const (
	defaultUserName   = "psshd"
	defaultAdminGroup = "psshd"

	defaultPort            = 2222
	defaultMaxAuthAttempts = 3
	defaultDataDir         = "/srv/protectedssh"
)

// DaemonOpts is the configuration opts specific
// to the protectedssh daemon.
type DaemonOpts struct {
	userName   string `yaml:"user_name"`
	adminGroup string `yaml:"admin_group"`

	port            uint   `yaml:"port"`
	maxAuthAttempts int    `yaml:"max_auth_attempts"`
	dataDir         string `yaml:"data_dir"`
}

// GetUserName returns the system user name to use
// as specified by the config file else the default.
func (opts *DaemonOpts) GetUserName() string {
	if opts == nil {
		return defaultUserName
	}
	return opts.userName
}

// GetAdminGroup returns the name of the system group
// that has admin abilities over the daemon.
func (opts *DaemonOpts) GetAdminGroup() string {
	if opts == nil {
		return defaultAdminGroup
	}
	return opts.adminGroup
}

// GetPort returns the daemon's port to listen to
// as specified by the config file else the default.
func (opts *DaemonOpts) GetPort() uint {
	if opts == nil {
		return defaultPort
	}
	return opts.port
}

// GetDataDir returns the location of the daemon's
// folder for managing users. The default dir is
// returned if unspecified.
func (opts *DaemonOpts) GetDataDir() string {
	if opts == nil {
		return defaultDataDir
	}
	return opts.dataDir
}

// GetMaxAuthAttempts returns the max number of attempts
// before a user is denied. If unspecified, returns default.
func (opts *DaemonOpts) GetMaxAuthAttempts() int {
	if opts == nil {
		return defaultMaxAuthAttempts
	}
	return opts.maxAuthAttempts
}
