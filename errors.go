package main

const (
	// ErrNotRoot indicates that this command requires root
	// user access.
	ErrNotRoot = "Please run as root."

	// ErrNotDaemon indicates that this command requires being
	// run as the protectedssh daemon user.
	ErrNotDaemon = "Please run as the daemon user."

	// ErrNotAdmin indicates that the user is not in the daemon
	// admin group.
	ErrNotAdmin = "You do not have administrative privileges for protectedssh."
)

func getExitCode(err error) int {
	return 1
}
