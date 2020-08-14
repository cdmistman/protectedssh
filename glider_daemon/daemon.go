package daemon

import (
	"fmt"
	"sync"

	"github.com/cdmistman/protectedssh/config"
	"github.com/gliderlabs/ssh"
)

// Daemon is the current daemon process.
type Daemon struct {
	mux sync.Mutex

	cfg *config.Config
}

// New returns a new Daemon.
func New(cfg *config.Config) Daemon {
	return Daemon{
		cfg: cfg,
	}
}

// Serve starts listening for incoming ssh connections
// and handles them appropriately.
func (daemon *Daemon) Serve() error {
	server := ssh.Server{}
	defer server.Close()

	// init the ssh server
	server.Addr = fmt.Sprintf(":%v", daemon.cfg.DaemonOpts.GetPort())
	server.Handler = userHandler
	server.HostSigners = daemon.getHostSigners()
	server.PublicKeyHandler = publicKeyHandler
	server.ServerConfigCallback = serverConfigCallback

	// this should be in a loop to allow stopping and starting
	return server.ListenAndServe()
}

//TODO:
func (daemon *Daemon) getHostSigners() []ssh.Signer {
	res := make([]ssh.Signer, 0)
	return res
}
