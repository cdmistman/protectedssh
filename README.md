# protectedssh

protectedssh is an SSH daemon for running SSH users in a Docker container.

## How it works

The administrator can create users easily using the CLI or with a configuration
file. The options for this user are then used for creating a Docker image that
the user will be given to once successfully authenticated.

## Config file

The config file right now is YAML, but honestly I prefer TOML so I'm going to
change it once I finish everything else.

The file has two main keys: `users` and `daemon`. The daemon key is for
configuring how the daemon works, and the users key is an array for maintaining
users. I'll write documentation on this eventually.

## Why Go?

I wanted to develop this as quickly as possible, which is an advantage Go has
over Rust. I also want to use the Docker SDK as the Rust bindings aren't quite
complete yet.

## License

I've MIT licensed this. Just follow it.
