# dccs-to-ynab
A tool for importing transactions from a DCCS CardApp account into a given YNAB budgeted account.

## Installation

### Using Docker
Leveraging the official [docker image](https://hub.docker.com/r/mastermel/dccs-to-ynab) is the
recommended way of running dccs-to-ynab.

#### Configure via docker
With docker installed, navigate to the directory you'd like to run dccs-to-ynab from and
begin configuring it with the following command:
```
docker run -itv $(pwd)/config.yml:/dccs-to-ynab.yml \
  mastermel/dccs-to-ynab:latest \
  accounts create
```

See the rest of [the setup guide below](#setup-the-accounts-you-want-to-sync) for explanation on what's expected to create an account.

#### Run as a service using Docker Compose
If you'd like to have dccs-to-ynab run periodically and sync your accounts, managing it with
a Docker Compose is an easy way. You can use a compose file such as the following:
```
# docker-compose.yml

version: "3"

services:
  dccs-to-ynab:
    image: mastermel/dccs-to-ynab:latest
    container_name: dccs-to-ynab
    restart: unless-stopped

    # Uncomment the following if you're running on an ARM processor (eg. RaspberryPi)
    #platform: linux/arm64

    command: [
      "--config=/dccs-to-ynab.yml",
      "sync-loop"
    ]

    environment:
      - SYNC_INTERVAL="8h" # Change this to the interval you prefer

    volumes:
      - ${PWD}/config.yml:/dccs-to-ynab.yml
```

### Install manually via Go

#### Install Go 1.22 or later
You can find instructions for downloading and installing the latest version of Go
at [https://golang.org/doc/install](https://golang.org/doc/install)

#### Get the app
**Note:**
Right now there's something screwy in how `go get` resolves the `go.bmvs.io/ynab` module.
It somehow installs v2.0.0, where I can only get this app to to depend on v1.3.0. The later
version changes the signature for a method depended on, and thus is incompatible with this
module at the moment :(

```
git clone https://github.com/mastermel/dccs-to-ynab.git

cd dccs-to-ynab

go install github.com/mastermel/dccs-to-ynab
```

### Setup the accounts you want to sync
```
dccs-to-ynab accounts create
```

Or if you're running through Docker:
```
docker run -itv $(pwd)/config.yml:/dccs-to-ynab.yml \
  mastermel/dccs-to-ynab:latest \
  accounts create
```

You'll be asked for the following fields:

#### Name
A friendly name to identify this account by. Maybe use your own?

#### Enable Sync?
Whether you want to sync transactions for this account when `sync` is run. Use `false` if you
want to configure this account but not sync it yet.

#### Date to sync from (yyyy-mm-ddThh:mm:ss)
What is the earliest timestamp you want to find DCCS transactions for? The `sync` command will
ignore any transactions it finds earlier than this value. It will also update this value to the
current time whenever `sync` is run.

#### DCCS Username
The username used to login to the desired DCCS Card App account.

#### DCCS Password
The password used to login to the desired DCCS Card App account.

#### DCCS Pay Code
The pay code associated with the DCCS Card you want to pull transactions from. For example `61212a`

#### YNAB API Token
The complete API Token generated for the YNAB account you want to import transactions into.
See [https://api.youneedabudget.com](https://api.youneedabudget.com) for instructions on
generating an API token.

**Note:** if the API token you provide is not valid, you will be unable to complete the account configuration.

#### YNAB Budget to sync to?
Granted you provided a valid API Token in the previous step, you'll be prompted with a list
of budgets from the associated YNAB account. Select the one that you want to import
transactions into.

#### YNAB Budget account to sync to?
Here you should be prompted with a list of bank accounts within the YNAB budget you selected
displayed as `type: Name`. Select the one that you want to import transactions into.

### Run the sync!
```
dccs-to-ynab sync
```

Or with Docker:
```
docker run -v $(pwd)/config.yml:/dccs-to-ynab.yml \
  mastermel/dccs-to-ynab:latest \
  sync
```

You can run this sync as often as you like and it will only import DCCS transactions created
since the last time it was run.
