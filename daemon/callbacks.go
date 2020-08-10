package daemon

import (
	"github.com/cdmistman/protectedssh/config"
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

const (
	// KeyGetConfig is the key in the ssh Context
	KeyGetConfig = "protectedssh_ctxkey_options"
)

// serverConfigCallback generates a ServerConfig for the crypto/ssh
// Server configuration. This is in alternative to generating an empty
// base config.
func serverConfigCallback(ctx ssh.Context) (res *gossh.ServerConfig) {
	opts := ctx.Value(KeyGetConfig).(*config.Config)

	res.MaxAuthTries = opts.DaemonOpts.GetMaxAuthAttempts()

	//TODO:
	// - AuthLogCallback for logging auth attempts
	// - BannerCallback for showing banners
	// ? GSSAPIWithMICCallback
	// ? RekeyThreshhold
	// ? KeyExchanges
	// ? Ciphers
	// ? MACs

	return
}

// publicKeyHandler checks to see if the incoming user has a matching
// private key on file.
func publicKeyHandler(ctx ssh.Context, key ssh.PublicKey) bool {
	opts := ctx.Value(KeyGetConfig).(*config.Config)

	localUser := opts.GetUser(ctx.User())
	if localUser == nil {
		return false
	}

	keys, err := localUser.Keys(opts.DaemonOpts.GetDataDir())
	if err != nil {
		return false
	}

	for _, unparsedKey := range keys {
		userKey, err := ssh.ParsePublicKey([]byte(unparsedKey))
		if err != nil {
			continue
		}

		if userKey == key {
			return true
		}
	}
	return false
}
