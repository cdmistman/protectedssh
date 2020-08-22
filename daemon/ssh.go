package daemon

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net"
	"reflect"

	"golang.org/x/crypto/ssh"
)

const (
	// PSSHSoftwareVersion is the softwareversion field of RFC 4253 section 4.2.
	PSSHSoftwareVersion string = "PROTECTEDSSH0.1"
)

// sshServerPortMessage is a message from an sshServerPort.
type sshServerPortMessage struct {
	port int
	err  error
	conn *net.TCPConn
}

// sshServerPort represents a port being listened to by pssh.
type sshServerPort struct {
	port int

	tcp *net.TCPListener
}

// Open opens the tcp port specified.
func (port *sshServerPort) Open() (err error) {
	addr := net.TCPAddr{
		Port: port.port,
	}
	port.tcp, err = net.ListenTCP("tcp", &addr)
	return
}

type sshServer struct {
	// send errors to the runner.
	err chan error
	// messages from docker goroutines
	fromDocker chan string

	ports []sshServerPort

	serverConfig ssh.ServerConfig
}

// runSSHServer runs the SSH server from the pssh opts.
// Errors are sent through the channel;
// it's up to the runner to decide if it's fatal.
func runSSHServer(opts *Opts) (errSend chan error) {
	var err error

	opts.sshComm = make(chan string)
	errSend = make(chan error)

	// for every port that pssh is configured to run on,
	// open the port for listening to incoming clients.
	var ports []sshServerPort
	if err == nil {
		for port := range opts.ports {
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

	// create the server and run it.
	server := sshServer{}
	server.ports = ports
	server.err = errSend
	server.fromDocker = opts.dockerComm
	server.serverConfig = _sshConfig(opts)
	go _sshRun(server)

	return
}

// _sshConfig creates the ssh.ServerConfig from the pssh options.
func _sshConfig(opts *Opts) (res ssh.ServerConfig) {
	// Config TODO: decide if it's worth tampering with these config options
	// NoClientAuth
	res.NoClientAuth = false

	// MaxAuthTries
	res.MaxAuthTries = opts.maxAuthTries

	// PasswordCallback
	res.PasswordCallback = func(conn ssh.ConnMetadata, password []byte) (perm *ssh.Permissions, err error) {
		perm = &ssh.Permissions{}
		reject := func() {
			// golang.org/x/crypto/ssh doesn't actually read the error
			err = errors.New("")
		}

		username := conn.User()
		user, ok := opts.users[username]
		if !ok {
			reject()
			return
		}

		hashedPass := sha256.Sum256(password)
		if hashedPass != user.passHash {
			reject()
			return
		}

		return
	}

	// PublicKeyCallback
	// KeyboardInteractiveCallback
	// AuthLogCallback

	// ServerVersion
	comments := "" // TODO: maybe comments? RFC 4253 section 4.2
	res.ServerVersion = fmt.Sprintf("SSH-2.0-%s%s\r\n", PSSHSoftwareVersion, comments)

	// BannerCallback
	// GSSAPIWithMICCallback
	// private keys

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
		_, value, ok := reflect.Select(chans)
		msg := value.Interface().(*sshServerPortMessage)
		if !ok {
			// TODO: log and handle
			msg.conn.Close()
			continue
		}

		if msg.err != nil {
			// TODO: log and handle
			msg.conn.Close()
			continue
		}

		go _handleIncoming(&server, msg)
	}
}

func _handleIncoming(server *sshServer, msg *sshServerPortMessage) {
	conn, chans, reqs, err := ssh.NewServerConn(msg.conn, &server.serverConfig)
	if err != nil {
		// TODO: log and handle
		return
	}

	for newChan := range chans {
		switch newChan.ChannelType() {
		case "session":
			go ssh.DiscardRequests(reqs)
			err = _handleSession(newChan, conn)

		// TODO:
		case "x11":
			fallthrough

		// TODO:
		case "forwarded-tcpip":
			fallthrough

		// TODO:
		case "direct-tcpip":
			fallthrough

		default:
			go ssh.DiscardRequests(reqs)
			// TODO: log and handle
		}
	}

	if err != nil {
		// TODO: log and handle
	}
}

func _handleSession(nc ssh.NewChannel, conn *ssh.ServerConn) (err error) {
	_, _, err = nc.Accept()
	if err != nil {
		return
	}

	// TODO: create/load docker for user
	// TODO: handle requests

	return
}
