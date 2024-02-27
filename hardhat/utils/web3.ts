import hre, { ethers } from 'hardhat'
import bluebird from 'bluebird'
import {
  BigNumberish,
  Signer,
  AddressLike
} from 'ethers'
import { Account } from './types'
import {
  ACCOUNTS,
  getAccount,
} from './accounts'
import {
  CoopHiveToken,
  CoopHivePayments,
  CoopHiveStorage,
  CoopHiveMediationRandom,
  CoopHiveController,
  CoopHiveOnChainJobCreator,
  CoopHiveUsers,
  ExampleClient,
} from '../typechain-types'

/*

  CONSTANTS

*/

// how much ether to send to each account
export const DEFAULT_ETHER_PER_ACCOUNT = ethers.parseEther('1000')

// a billion tokens in total
export const DEFAULT_TOKEN_SUPPLY = ethers.parseEther('1000000000')

// each service gets 100000 tokens
export const DEFAULT_TOKENS_PER_ACCOUNT = ethers.parseEther('100000')

/*

  WALLET UTILS

*/
export const getWallet = (name: string) => {
  const account = getAccount(name)
  return new ethers.Wallet(account.privateKey, ethers.provider)
}

export const getAddress = (name: string) => {
  const account = getAccount(name)
  return account.address
}

export const getRandomWallet = () => {
  return ethers.Wallet.createRandom()
}

export const transferEther = async (fromAccount: Account, toAccount: Account, amount: BigNumberish) => {
  const signer = new hre.ethers.Wallet(fromAccount.privateKey, hre.ethers.provider)
  const tx = await signer.sendTransaction({
    to: toAccount.address,
    value: amount,
  })
  await tx.wait()
  console.log(`Moved ${amount} ETHER from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}

export const transferTokens = async (fromAccount: Account, toAccount: Account, amount: BigNumberish) => {
  const token = await connectToken()
  const tx = await token
    .connect(getWallet(fromAccount.name))
    .transfer(toAccount.address, amount)
  await tx.wait()
  console.log(`Moved ${amount} TOKENS from ${fromAccount.name} (${fromAccount.address}) to ${toAccount.name} (${toAccount.address}) - ${tx.hash}.`)
}


/*

  DEPLOYMENT

*/
export async function deployContract<T extends any>(
  name: string,
  signer: Signer,
  args: any[] = [],
): Promise<T>{
  const factory = await ethers.getContractFactory(
    name,
    signer,
  )
  const contract = await factory.deploy(...args) as unknown as T
  return contract
}

/*

  TOKEN

*/
export async function fundTokens(
  tokenContract: CoopHiveToken,
  address: AddressLike,
  amount: BigNumberish = DEFAULT_TOKENS_PER_ACCOUNT,
) {
  await tokenContract
    .connect(getWallet('admin'))
    .transfer(address, amount)
}

export async function fundAccountsWithTokens(
  tokenContract: CoopHiveToken,
  amount: BigNumberish = DEFAULT_TOKENS_PER_ACCOUNT,
) {
  await bluebird.mapSeries(ACCOUNTS, async (account) => {
    await fundTokens(tokenContract, account.address, amount)
  })
}

/*

  HARDHAT DEPLOYMENTS

  these all require the deployments to have been done via hardhat
  and for there to be deployment.json files for us to look inside of

*/

export async function connectContract<T extends any>(
  name: string,
): Promise<T>{
  const deployment = await hre.deployments.get(name)
  const factory = await hre.ethers.getContractFactory(name)
  const contract = factory.attach(deployment.address) as unknown as T
  return contract
}

export async function getContractAddress(
  name: string,
): Promise<AddressLike>{
  const deployment = await hre.deployments.get(name)
  return deployment.address
}

/*

  storage

*/
export async function connectStorage() {
  return connectContract<CoopHiveStorage>('CoopHiveStorage')
}

export async function getStorageAddress() {
  return getContractAddress('CoopHiveStorage')
}

/*

  mediation

*/
export async function connectMediation() {
  return connectContract<CoopHiveMediationRandom>('CoopHiveMediationRandom')
}

export async function getMediationAddress() {
  return getContractAddress('CoopHiveMediationRandom')
}

/*

  token

*/
export async function connectToken() {
  return connectContract<CoopHiveToken>('CoopHiveToken')
}

export async function getTokenAddress() {
  return getContractAddress('CoopHiveToken')
}

/*

  payments

*/
export async function connectPayments() {
  return connectContract<CoopHivePayments>('CoopHivePayments')
}

export async function getPaymentsAddress() {
  return getContractAddress('CoopHivePayments')
}


/*

  job manager

*/
export async function connectJobManager() {
  return connectContract<CoopHiveOnChainJobCreator>('CoopHiveOnChainJobCreator')
}

export async function getJobManagerAddress() {
  return getContractAddress('CoopHiveOnChainJobCreator')
}


/*

  users

*/
export async function connectUsers() {
  return connectContract<CoopHiveUsers>('CoopHiveUsers')
}

export async function getUsersAddress() {
  return getContractAddress('CoopHiveUsers')
}

/*

  example client

*/
export async function connectExampleClient() {
  return connectContract<ExampleClient>('ExampleClient')
}

export async function getExampleClientAddress() {
  return getContractAddress('ExampleClient')
}

/*

  controller

*/
export async function connectController() {
  return connectContract<CoopHiveController>('CoopHiveController')
}

export async function getControllerAddress() {
  return getContractAddress('CoopHiveController')
}
