import { connectStorage, connectController } from "../utils/web3";
import { ethers, upgrades, deployments } from "hardhat";
async function main() {
  const privateKey = process.env.WEB3_PRIVATE_KEY;

  if (!privateKey) {
    console.error(`WEB3_PRIVATE_KEY env variable is required`);
    process.exit(1);
  }

  const wallet = new ethers.Wallet(privateKey).connect(ethers.provider);

  const storageProxy = await connectStorage();
  const deployment = await deployments.get("CoopHiveStorage");
  const factory = await ethers.getContractFactory("CoopHiveStorage");
  console.log("factory", factory);

  const controllerProxy = await connectController();

  const CoopHiveStorage = await ethers.getContractFactory("CoopHiveStorage");
  const storage = await upgrades.upgradeProxy(storageProxy, CoopHiveStorage, {
    redeployImplementation: "onchange",
  });

  const repointTx = await controllerProxy.setStorageAddress(
    await storage.getAddress()
  );
  //const newStorageImpl = await CoopHiveStorage.deploy();
  //const upgradetx = await controllerProxy.reinitialize(newStorageImpl.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
