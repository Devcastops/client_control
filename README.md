# Client_control

This application has been created to create and mange cloud vms through nomad. Its it is a cli application that creates and destroys instances, along with setting up cloudflare DNS.

## To use

To be able to run this application the terminal running it needs to be set up:

For packer the client_id and client_secret needs to be exported as env vars
```bash
HCP_CLIENT_ID=
HCP_CLIENT_SECRET=
```
For GCP your terminal needs to be authenticated with google, the easyest way to do this is with.
```bash
gcloud login
```

You then need to define a config. copy `config.json.example` as `config.json` and fill out the required infomation in it.

## commands

To get the most up to date commands run `client_control -h`

### start

Will start the desired Instance

Usage:
  client_control start [flags]

Flags:
  -c, --config string           location of config file (default "config.json")
  -h, --help                    help for start
  -m, --machine_type string     The machine_type to use (default "e2-standard-2")
  -n, --name string             name to give the instance (default "test")
  -l, --node_pool string        The node_pool to deploy the client in
  -i, --packer_channel string   packer channel to pull from (default "live")
  -p, --provider string         cloud provider to use (GCP) (default "GCP")

### stop

Will stop the desired Instance

Usage:
  client_control stop [flags]

Flags:
  -c, --config string     location of config file (default "config.json")
  -h, --help              help for stop
  -n, --name string       name to give the instance (default "test")
  -p, --provider string   cloud provider to use (GCP) (default "GCP")

### get

Will return infomation about the running instance

Usage:
  client_control get [flags]

Flags:
  -c, --config string     location of config file (default "config.json")
  -h, --help              help for get
  -n, --name string       name to give the instance (default "test")
  -p, --provider string   cloud provider to use (GCP) (default "GCP")

### autostop

Will run the instance untill it passes the given time (and/or) date. If the time is before the current time, it will be the next day it stops

Usage:
  client_control autostop [flags]

Flags:
  -d, --autostopdate string   date to stop the server
  -t, --autostoptime string   time to stop the server (default "00:00")
  -c, --config string         location of config file (default "config.json")
  -h, --help                  help for autostop
  -n, --name string           name to give the instance (default "test")
  -p, --provider string       cloud provider to use (GCP) (default "GCP")