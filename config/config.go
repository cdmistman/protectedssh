package config

import (
	"errors"
	"os"
	"sync"

	"github.com/cdmistman/protectedssh/user"
	"gopkg.in/yaml.v2"
)

const (
	// ErrCreateEncoder indicates that the file encoder was
	// returned nil.
	ErrCreateEncoder = "File encoder not created succesfully."

	// ErrCreateDecoder indicates that the file decoder was
	// returned nil.
	ErrCreateDecoder = "File decoder not created succesfully."
)

const (
	// ConfigFileLocation is the directory that stores all of the config
	ConfigFileLocation = "/etc/protectedssh.yaml"
)

// Config is the configuration file for protectedssh.
type Config struct {
	mux sync.Mutex `yaml:"-"`

	// daemon configuration opts
	DaemonOpts *DaemonOpts `yaml:"daemon"`

	// user configurations
	users []*user.User `yaml:"users"`
}

// New returns a Config that represents the file for
// configuring protectedssh. If file already exists,
// loads the saved information. If the file doesn't
// exist, just returns the default Config.
func New() (config Config) {
	// default stuff
	config.DaemonOpts = nil
	config.users = make([]*user.User, 0)

	// load any options already in the file
	config.Load()

	return
}

// Load loads all data in the configuration file that
// have been specified. If any value in the current
// Config hasn't been saved yet, it is overwritten with
// the value specified in the file, if it exists.
func (config *Config) Load() (err error) {
	config.mux.Lock()
	defer config.mux.Unlock()

	// if the config file doesn't exist, just return defaults
	if _, err = os.Stat(ConfigFileLocation); err != nil {
		// assume that the file doesn't exist, there's an issue
		// getting the stats for it anyways.
		return
	}

	// load the opts from the config file
	file, err := os.OpenFile(ConfigFileLocation, os.O_RDWR|os.O_TRUNC, 0666)
	if file == nil {
		return
	}
	defer file.Close()
	if err != nil {
		return
	}

	// decode into this
	decoder := yaml.NewDecoder(file)
	if decoder == nil {
		err = errors.New(ErrCreateDecoder)
		return
	}
	decoder.Decode(config)

	return
}

// Save saves the current Config in the file location that
// is specified by the fileLocation field set during New.
// Note that this method does not get called internally and
// changes will not be saved until this method gets called.
func (config *Config) Save() (err error) {
	config.mux.Lock()
	defer config.mux.Unlock()

	file, err := os.OpenFile(ConfigFileLocation, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if file == nil {
		return
	}
	defer file.Close()
	if err != nil {
		return
	}

	encoder := yaml.NewEncoder(file)
	if encoder == nil {
		err = errors.New(ErrCreateEncoder)
		return
	}
	defer encoder.Close()

	err = encoder.Encode(config)
	return
}
