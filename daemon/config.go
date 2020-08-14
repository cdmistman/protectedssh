package daemon

// ExitMode is the mode for exiting.
// Only the explicit constants should be used.
type ExitMode int8

const (
	// ExitRude indicates that the daemon should exit as soon as it receives this notification.
	// Nothing should be expected to be logged and there should be no cleanup.
	ExitRude ExitMode = 1

	// ExitErr is the same as ExitRude but should log the message for exiting.
	ExitErr ExitMode = -1

	// ExitClean indicates that the daemon should exit cleanly when it's safe to do so.
	// Only the cleanup will be logged.
	ExitClean ExitMode = 2

	// ExitWithReason is the same as ExitClean but should log the reason for exiting.
	ExitWithReason ExitMode = -2
)

// Message is some message being sent to the daemon that should be logged.
type Message string

// Opts is the options that are used to run the daemon.
type Opts struct {
	exit chan ExitMode

	sshComm    chan string
	dockerComm chan string

	ports []int

	users []User // The list of user configurations.
}
