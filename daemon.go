package main

// TODO: update to just use github.com/cdmistman/daemon for pssh.
// TODO: use github.com/kardianos/service for daemon management.

import (
	"errors"

	"github.com/cdmistman/protectedssh/config"
	"github.com/urfave/cli/v2"
)

// Installs the daemon on the current linux system.
func daemonInstall(ctx *cli.Context) (err error) {
	// ensure root
	if !isRootUser() {
		err = errors.New(ErrNotRoot)
		return
	}

	// create daemon user
	// set daemon user's home location
	// create daemon config
	// register startup

	return
}

// Runs the daemon. Only executable by the daemon user.
func daemonRun(ctx *cli.Context) (err error) {
	cfg := config.New()

	// ensure daemon user
	if !isDaemonUser(&cfg) {
		err = errors.New(ErrNotDaemon)
		return
	}

	// construct and run daemon
	// server := daemon.New(&cfg)
	// return server.Serve()
	return
}

// Tells the daemon server to start handling ssh requests.
func daemonStart(ctx *cli.Context) (err error) {
	cfg := config.New()

	// ensure admin user
	if !isDaemonUser(&cfg) {
		err = errors.New(ErrNotAdmin)
		return
	}

	//TODO

	return
}

// Tells the daemon server to stop handling ssh requests.
func daemonStop(ctx *cli.Context) (err error) {
	cfg := config.New()

	// ensure admin user
	if !isDaemonUser(&cfg) {
		err = errors.New(ErrNotAdmin)
		return
	}

	//TODO

	return
}

// Completely removes the daemon from the system.
func daemonUninstall(ctx *cli.Context) (err error) {
	if !isRootUser() {
		err = errors.New(ErrNotRoot)
		return
	}

	// tell daemon to stop
	// remove startup config
	// remove daemon user
	// remove daemon user home

	return
}
