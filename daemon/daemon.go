package daemon

import (
	"fmt"
	"os"
)

// Daemon represents a running pssh daemon.
type daemon struct {
}

// Run runs the daemon using the given DaemonOpts object.
func Run(opts Opts, exit chan ExitMode, log chan Message) (err error) {
	opts.exit = make(chan ExitMode)

	sshErr := runSSHServer(&opts)
	if err != nil {

		return
	}

	dockerErr := runDockerServer(&opts)
	if err != nil {
		return
	}

loop:
	for {
		select {
		case msg := <-log:
			fmt.Println(msg)
		case code := <-exit:
			if code == 0 {
				os.Exit(1)
			} else {
				break loop
			}
		case err = <-sshErr:
			break loop
		case err = <-dockerErr:
			break loop
		}
	}

	return
}
