package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

const (
	// ErrNoFile indicates that the specified file
	// doesn't exist for the current user.
	ErrNoFile = "Specified file not found: %v"
)

const (
	// First value is the data directory, second is the
	// user's name.
	keyFileFormat = "%v/%v/keys"
)

// User represents a permitted user in protectedssh.
type User struct {
	mux sync.Mutex `yaml:"-"`

	Name           string   `yaml:"name"`
	shareLocations []string `yaml:"share"`
	keysFile       string   `yaml:"keys_file"`
}

// NewUser returns a User with the default configuration
// options.
func NewUser(name string) (user User) {
	user.shareLocations = []string{}
	return
}

// GetKeys returns all keys for the user as located in the
// keys file.
func (user *User) GetKeys() []string {
	res := make([]string, 0)

	if user.keysFile == "" {
		user.keysFile = fmt.Sprintf("")
	}

	return res
}

// UpdateDocker updates the docker image for the user and
// updates the lock file with the current settings. If
// there is no Docker image for the user already, creates
// one.
func (user *User) UpdateDocker() {
}

// Keys returns all of the keys available in the user's
// keys file.
func (user *User) Keys(dataDir string) (keys []string, err error) {
	fileLocation := fmt.Sprintf(keyFileFormat, dataDir, user.Name)
	if _, err = os.Stat(fileLocation); err != nil {
		err = errors.New(ErrNoFile)
		return
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = text[:len(text)-1]
		keys = append(keys, text)
	}
	err = scanner.Err()

	return
}

// AddKey adds a key to the current user's keys file.
func (user *User) AddKey(dataDir, key string) error {
	fileLocation := fmt.Sprintf(keyFileFormat, dataDir, user.Name)

	file, err := os.OpenFile(fileLocation, os.O_APPEND|os.O_CREATE|os.SEEK_END, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	keyBytes := []byte(key)
	bw, err := file.Write(keyBytes)
	if err != nil {
		return err
	}
	if bw != len(keyBytes) {
		return errors.New("Partial key write")
	}

	return nil
}
