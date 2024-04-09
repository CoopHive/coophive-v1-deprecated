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
  console.log("current storage", await controller.getStorageAddress());

  const { deploy } = deployments;
  const storage = await deploy("CoopHiveStorage", {
    from: wallet.address,
    log: true,
  });

  const repointTx = await controller.setStorageAddress(storage.address);
  console.log("new storage", await controller.getStorageAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
