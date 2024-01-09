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

#### install stack

```bash
cd generic-dcn
./stack install
```

This script will:

 * build the faucet docker image from the locally cloned repo
 * download the go modules
 * install the node modules for hardhat
 * install the node modules for the frontend
 * compile the solidity contracts and generate the typechain bindings
 * generate the dev `.env` file with insecure private keys

After you've run the install script - you can look inside of `.env` to see the core service private keys and addresses that are used in the local dev stack.

### run web3 stack

These steps boot geth, deploy our contracts and ensure that the various services named in `.env` are funded with ether and tokens.

```bash
cd generic-dcn
./stack boot
```

This script will:

 * start geth as a docker container
 * fund the admin account with ether
 * fund the various services with ether
 * compile and deploy the solidity contracts
 * fund the various services with tokens
 * print the balances of the various accounts in `.env` 

### run services

Run the following commands in separate terminal windows:

```bash
./stack solver
```

Wait for the solver to start when `🟡 SOL solver registered` is logged, and then, in another terminal window, run:

```bash
./stack mediator
```

In another terminal window run:

```bash
./stack jobcreator
```

In another terminal window run:

```bash
./stack bacalhau-serve
```

If you have a GPU, run the following command in a separate terminal window:

```bash
./stack resource-provider --offer-gpu 1
```

Otherwise, if you don't have a GPU:

```bash
./stack resource-provider
```

### run faucet

To run the faucet container so you can test with other user accounts:

```bash
./stack faucet
```

Once the faucet is running, you can access it using http://localhost:8085

**NOTE**: if you want a different logo or otherwise a different design for the faucet, fork the [repo](https://github.com/bacalhau-project/eth-faucet) and use that as your basis for the faucet container.

You can find the frontend code in the `web` directory and the images are in the `web/public` directory.

### run jobs

Now you can run jobs on the stack as follows:

```bash
./stack run cowsay:v0.0.1 -i Message="moo"
```

If you have a GPU node - you can run SDXL (which needs a GPU):

```bash
./stack runsdxl sdxl:v0.9-lilypad1 PROMPT="beautiful view of iceland with a record player"
```

To demonstrate triggering jobs being run from on-chain smart contracts:

```bash
./stack run-cowsay-onchain
```

### stop stack

To stop the various services you have started in the numerous terminal windows, `ctrl+c` will suffice.

To stop geth:

```bash
./stack geth-stop
```

To stop the faucet:

```bash
./stack faucet-stop
```

To reset Geth data, effectively performing a complete restart, use the following command:

```bash
./stack clean
```

Please note that after running `clean`, you will need to re-run the `fund-admin` and `fund-services` commands.

### Unit Tests

Run the smart contract unit tests with the following command:

```bash
./stack unit-tests
```

### Regenerating Go Bindings

Whenever you make changes to the smart contracts, regenerate the Go bindings in `pkg/contract/bindings/contracts` by running:

```bash
./stack compile-contracts
```