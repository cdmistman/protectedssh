package daemon

import (
	"net"
	"reflect"
)

type sshServerPortMessage struct {
	port int
	err  error
	conn *net.TCPConn
}

type sshServerPort struct {
	port int

	tcp *net.TCPListener
}

func (port *sshServerPort) Open() (err error) {
	addr := net.TCPAddr{
		Port: port.port,
	}
	port.tcp, err = net.ListenTCP("tcp", &addr)
	return
}

type sshServer struct {
	err        chan error
	fromDocker chan string

	ports []sshServerPort
}

func runSSHServer(opts *Opts) (errSend chan error) {
	var err error

	opts.sshComm = make(chan string)
	errSend = make(chan error)

	var ports []sshServerPort
	if err == nil {
		for _, port := range opts.ports {
			serverPort := sshServerPort{
				port: port,
			}
			err = serverPort.Open()
			ports = append(ports, serverPort)
			if err != nil {
				break
			}
		}
	}

	if err != nil {
		for _, port := range ports {
			_ = port.tcp.Close()
		}
	}

	server := sshServer{}
	server.err = errSend
	server.fromDocker = opts.dockerComm
	go _sshRun(server)

	return
}

func _sshRun(server sshServer) {
	chans := make([]reflect.SelectCase, len(server.ports))
	for ii, port := range server.ports {
		ch := make(chan *sshServerPortMessage)
		go func(ch chan *sshServerPortMessage, port *sshServerPort) {
			send := func(conn *net.TCPConn, err error) {
				msg := &sshServerPortMessage{
					port: port.port,
					err:  err,
					conn: conn,
				}
				ch <- msg
			}

			tcp := port.tcp
			for {
				next, err := tcp.AcceptTCP()
				if err != nil {
					send(nil, err)
					return
				}
				send(next, err)
			}
		}(ch, &port)

		chans[ii] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	for {
		chosen, value, ok := reflect.Select(chans)
		//TODO:
		if !ok {
			return
		}

		_ = chans[chosen]
		_ = value.Interface().(*sshServerPortMessage)
	}
}
