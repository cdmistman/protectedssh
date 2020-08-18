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

## Wait, isn't this just another IaaS app?

Yes and no.

Yes, it is another IaaS app, as users don't have to interact with the machine setup, etc.
This means users can be added quickly, and you don't have to worry about them messing with your machine's settings.

However, this isn't a traditional IaaS app.
Typically, IaaS apps will load users in a dedicated VM that is almost entirely isolated from the host machine.
They'll have settings to expose ports, manage memory, etc.
With protectedssh, users still have the capacity to interact with the host machine (dependent on settings).
//TODO: this

## Why Go?

I wanted to develop this as quickly as possible, which is an advantage Go has
over Rust. I also want to use the Docker SDK as the Rust bindings aren't quite
complete yet.

## License

I've MIT licensed this. Just follow it.
