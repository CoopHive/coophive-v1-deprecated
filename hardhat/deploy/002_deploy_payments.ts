import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployPayments: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  await deploy("LilypadPayments", {
    from: admin,
    args: [],
    log: true,
  })

  const tokenContract = await deployments.get('LilypadToken')
  const paymentsContract = await deployments.get('LilypadPayments')

  await execute(
    'LilypadPayments',
    {
      from: admin,
      log: true,
    },
    'initialize',
    tokenContract.address
  )

  await execute(
    'LilypadToken',
    {
      from: admin,
      log: true,
    },
    'setControllerAddress',
    paymentsContract.address
  )

  return true
}

deployPayments.id = 'deployPayments'

export default deployPayments