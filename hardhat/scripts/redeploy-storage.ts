import { connectStorage, connectController } from "../utils/web3";
import { ethers, deployments } from "hardhat";
async function main() {
  const privateKey = process.env.WEB3_PRIVATE_KEY;

  if (!privateKey) {
    console.error(`WEB3_PRIVATE_KEY env variable is required`);
    process.exit(1);
  }

  const wallet = new ethers.Wallet(privateKey).connect(ethers.provider);

  const controller = await connectController();
  const initialControllerAddress = await controller.getAddress();
  const initalStorageAddr = await controller.getStorageAddress();

  const { deploy } = deployments;
  const storage = await deploy("CoopHiveStorage", {
    from: wallet.address,
    log: true,
  });

  const repointTx = await controller.setStorageAddress(storage.address);
  await repointTx.wait();
  const finalStorageAddr = await controller.getStorageAddress();
  const finalControllerAddress = await controller.getAddress();
  if (initialControllerAddress !== finalControllerAddress) {
    console.warn("Control address changed, something went very wrong");
    return;
  }

  if (initalStorageAddr == finalStorageAddr) {
    console.warn(
      "storage address unchanged, hardhat deploy wont redeploy if the contract is the same as prior"
    );
    return;
  }
  await deployments.execute(
    "CoopHiveStorage",
    {
      from: wallet.address,
      log: true,
    },
    "initialize"
  );

  await deployments.execute(
    "CoopHiveStorage",
    {
      from: wallet.address,
      log: true,
    },
    "setControllerAddress",
    initalStorageAddr
  );

  console.log("New Storage Address: ", finalStorageAddr);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
