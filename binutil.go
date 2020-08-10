package main

import (
	"os"
	"strconv"

	osUser "os/user"

	"github.com/cdmistman/protectedssh/config"
)

func isDaemonUser(cfg *config.Config) bool {
	usr, err := osUser.Current()
	if err != nil {
		return false
	}

	return usr.Name == cfg.DaemonOpts.GetUserName()
}

func isAdmin(cfg *config.Config) bool {
	adminGroup, err := osUser.LookupGroup(cfg.DaemonOpts.GetAdminGroup())
	if err != nil {
		return false
	}

	adminGid, err := strconv.Atoi(adminGroup.Gid)
	if err != nil {
		return false
	}

	userGroups, err := os.Getgroups()
	if err != nil {
		return false
	}

	for _, group := range userGroups {
		if group == adminGid {
			return true
		}
	}

	return false
}

func isRootUser() bool {
	usr, err := osUser.Current()
	if err != nil {
		return false
	}

	uid, err := strconv.Atoi(usr.Uid)
	if err != nil {
		return false
	}

	return uid == 0
}
