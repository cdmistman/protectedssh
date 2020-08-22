package daemon

// User is the configuration for a psshd user.
type User struct {
	passHash [32]byte
}
