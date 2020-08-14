package daemon

type dockerServer struct {
}

func runDockerServer(opts *Opts) (err chan error) {
	opts.dockerComm = make(chan string)

	_ = dockerServer{}

	return
}
