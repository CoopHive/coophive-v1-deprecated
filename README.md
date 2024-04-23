# CoopHive

CoopHive is a two-sided marketplace for computational resources. It enables clients to run computational workloads in a permissionless protocol, where anyone can get paid to connect their compute nodes to the network and run jobs. It uses an EVM-compatible blockchain to manage payments and storage, and uses [bacalhau](https://www.bacalhau.org/) to manage the compute nodes.

# Getting started

## Halcyon Testnet

The testnet has a base curency of ETH and you will also get HIVE to pay for jobs (and nodes to stake).

Metamask:

```
Network name: CoopHive Halcyon testnet
New RPC URL: http://halcyon.co-ophive.network:8545
Chain ID: 1337
Currency symbol: ETH
Block explorer URL: (leave blank)
```

### Fund your Wallet with ETH and HIVE

To obtain funds, go to [http://halcyon-faucet.co-ophive.network:8085](http://halcyon-faucet.co-ophive.network:8085)

The faucet will give you both ETH (to pay for gas) and HIVE (to stake and pay for jobs).

## Install CLI

Download the latest release of `hive` for your platform. Both the amd64/x86_64 and arm64 variants of macOS and Linux are supported. (If you are on Apple Silicon, you'll want arm64). 

Nb:  to check your version use ```which hive``` - if an old version run ```rm <path>``` to remove that path then reinstall newest version

The commands below will automatically detect your OS and processor architecture and download the correct build for your machine.

```
# Detect your machine's architecture and set it as $OSARCH
OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && export OSARCH
# Detect your operating system and set it as $OSNAME
OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && export OSNAME
```
Then Download & Install
```
# Download the latest production build
curl -sSL -o hive https://github.com/CoopHive/coophive/releases/download/v0.1.0-5e6070b/hive-$OSNAME-$OSARCH
# Make executable and install it
chmod +x hive
sudo mv hive /usr/local/bin/hive
```

You can also, at your option, choose to compile `hive` using Go and install it that way on any machine that supports the Go toolchain.


## Run a Job

First, make sure your Web3 private key is in the environment.

```
export WEB3_PRIVATE_KEY=<your private key>
```

Alternatively, arrange for the key to be in your environment in a more secure way that doesn't get written to your shell history.


### Cows

```
hive run cowsay:v0.0.1 -i Message="hello world"
```


### SDXL

```
hive run sdxl:v0.2.9 -i PromptEnv="PROMPT=a new hexagonal beehive in the halcyon fields of springtime"
```

Not working?
Try ```rm -rf /tmp/coophive/data/repos```. Then reinstall from the start.


## Run a Node, Earn HIVE


```
hive resource-provider --offer-cpu 1  --offer-ram 1024 --offer-gpu 0 --offer-count 1
```

systemd units & more details [here](https://github.com/CoopHive/coophive/blob/main/ARCHITECTURE.md)


# Available Modules

Check the github releases page for each module or just use the git hash as the tag.

* [sdxl](https://github.com/CoopHive/coophive-module-sdxl)
* [cowsay](https://github.com/CoopHive/coophive-module-cowsay)

More coming soon!


# Write a Module



A module is just a git repo.

Module versions are just git tags.

In your repo, create a file called `module.coophive`

See [cowsay](https://github.com/CoopHive/coophive-module-cowsay) for example

This is a json template with Go text/template style `{{.Message}}` sections which will be replaced by hive with json encoded inputs to modules. You can also do fancy things with go templates like setting defaults, see cowsay for example. While developing a module, you can use the git hash to test it.

Pass inputs as:

```
hive run github.com/username/repo:tag -i Message=moo
```

Inputs are a map of strings to strings.

### Writing Advanced Modules

#### `subt`

The `subt` function allows for substitutions in your template.

This function is a workaround for the lack of direct substitution support in the module. It implements the [printf](https://pkg.go.dev/text/template#Template.Funcs) function under the hood, which allows you to format strings with placeholders.

##### Usage   

The `subt` function can be used in the same way as the `printf` function in Go. You pass in a format string, followed by values that correspond to the placeholders in the format string.
```
const templateText = {{ subt "Hello %s" .name }}
```

