import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const deployController: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployments, getNamedAccounts } = hre
  const { deploy, execute } = deployments
  const {
    admin,
  } = await getNamedAccounts()
  await deploy("CoopHiveController", {
    from: admin,
    args: [],
    log: true,
  })
  
  const controllerContract = await deployments.get('CoopHiveController')
  const storageContract = await deployments.get('CoopHiveStorage')
  const usersContract = await deployments.get('CoopHiveUsers')
  const mediationContract = await deployments.get('CoopHiveMediationRandom')
  const paymentsContract = await deployments.get('CoopHivePayments')
  const jobCreatorContract = await deployments.get('CoopHiveOnChainJobCreator')

  await execute(
    'CoopHiveController',
    {
      from: admin,
      log: true,
    },
    'initialize',
    storageContract.address,
    usersContract.address,
    paymentsContract.address,
    mediationContract.address,
    jobCreatorContract.address
  )

  await execute(
    'CoopHiveStorage',
    {
      from: admin,
      log: true,
    },
    'setControllerAddress',
    controllerContract.address, 
  )

  await execute(
    'CoopHivePayments',
    {
      from: admin,
      log: true,
    },
    'setControllerAddress',
    controllerContract.address, 
  )

  await execute(
    'CoopHiveMediationRandom',
    {
      from: admin,
      log: true,
    },
    'setControllerAddress',
    controllerContract.address, 
  )

  return true
}

deployController.id = 'deployController'

export default deployController