package config

import (
	"github.com/cdmistman/protectedssh/user"
)

// AddUser adds a new user to the permitted users section of
// the protectedssh config file.
func (config *Config) AddUser(newUser *user.User) {
	config.users = append(config.users, newUser)
}

// RemoveUser removes a user from the permitted users section
// of the protectedssh config file.
func (config *Config) RemoveUser(name string) {
	if index := config.findUser(name); index >= 0 {
		// remove the user at index
		config.users = append(config.users[:index], config.users[index+1:]...)
	}
}

// GetUser returns a pointer to the user defined in the config
// file. Again, modifications to the User won't be saved unless
// Save is called.
func (config *Config) GetUser(name string) (user *user.User) {
	index := config.findUser(name)
	if index >= 0 {
		user = config.users[index]
	} else {
		user = nil
	}
	return
}

// findUser returns the index of the user that has the specified
// name. Is negative if no such user exists.
func (config *Config) findUser(name string) int {
	for ii, declaredUser := range config.users {
		if declaredUser.Name == name {
			return ii
		}
	}
	return -1
}
