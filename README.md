# generic distributed compute network

Generic DCN is a project that aims to facilitate the creation of a distributed compute network.

It will use an EVM based blockchain to manage agreed job state and payment and use [bacalhau](https://www.bacalhau.org/) to manage the compute nodes.

It attempts to bring various tools without a strog opinion on how they should be used together.  For example, it is possible to deploy the smart contracts on any EVM compatible blockchain and implement an interface to use any other tool than bacalhau to manage the compute nodes.

## architecture

<put images here and describe how it all fits together>

## local development

This section will demonstrate how to run the stack on your local machine.

### pre-requisites

You will need the following tools:

 * go (>= v1.20)
 * docker
 * node.js (>= v16)

### initial setup

These steps only need to be done once.

#### install bacalhau

We are currently pinned to bacalhau v1.0.3 - to install this version run the following commands:

```bash
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
# Extract the downloaded archive and move the `bacalhau` binary to `/usr/local/bin`
tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
mv bacalhau /usr/local/bin
# Set the IPFS data path by exporting the `BACALHAU_SERVE_IPFS_PATH` variable to your desired location
export BACALHAU_SERVE_IPFS_PATH=/tmp/lilypad/data/ipfs
```

#### clone faucet repo

The [faucet](https://github.com/bacalhau-project/eth-faucet) allows us to mint tokens for testing purposes.

We first need to clone the repo:

```bash
# run this command at the same level as the generic-dcn repo
git clone git@github.com:bacalhau-project/eth-faucet.git
```

#### install

Run the following script which will:

 * build the faucet docker image from the locally cloned repo
 * download the go modules
 * install the node modules for hardhat
 * install the node modules for the frontend

```bash
cd generic-dcn
./stack install
```