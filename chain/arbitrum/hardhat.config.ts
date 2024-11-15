import { task } from "hardhat/config";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { BigNumber } from "ethers";
import "@nomiclabs/hardhat-waffle";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-etherscan";
import "@typechain/hardhat";
import "hardhat-contract-sizer";
import "hardhat-gas-reporter";
import "hardhat-deploy";
import "solidity-coverage";
// import dotenv to get environment variables
import dotenv from "dotenv";

dotenv.config();
const FORK_SEPOLIA = false;
const FORK_MAINNET = true;
const forkingData = FORK_SEPOLIA
  ? {
      url: "https://sepolia-rollup.arbitrum.io/rpc",
      blockNumber: 16343472,
    }
  : FORK_MAINNET
    ? {
        url: `https://arb-mainnet.g.alchemy.com/v2/${process.env.ALCHEMY_API_KEY}`,
        blockNumber: 195245031,
      }
    : undefined;

task(
  "accounts",
  "Prints the list of accounts",
  async (args, hre): Promise<void> => {
    const accounts: SignerWithAddress[] = await hre.ethers.getSigners();
    accounts.forEach((account: SignerWithAddress): void => {
      console.log(account.address);
    });
  },
);

task(
  "balances",
  "Prints the list of ARB account balances",
  async (args, hre): Promise<void> => {
    const accounts: SignerWithAddress[] = await hre.ethers.getSigners();
    for (const account of accounts) {
      const balance: BigNumber = await hre.ethers.provider.getBalance(
        account.address,
      );
      console.log(`${account.address} has balance ${balance.toString()}`);
    }
  },
);

export default {
  solidity: {
    version: "0.8.9",
    settings: {
      optimizer: {
        enabled: true,
        runs: 100000,
      },
    },
  },
  networks: {
    hardhat: {
      gasPrice: 225000000000,
      chainId: !forkingData ? 43112 : 43114,
      forking: forkingData,
    },
    local: {
      url: "http://127.0.0.1:8545/",
      gasPrice: 225000000000,
      chainId: 412346,
    },
    sepolia: {
      url: "https://sepolia-rollup.arbitrum.io/rpc",
      gasPrice: 225000000000,
      chainId: 421614,
      accounts: [],
    },
    mainnet: {
      url: "https://arb1.arbitrum.io/rpc",
      chainId: 42161,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
  },
  paths: {
    deploy: "./scripts/deploy",
    deployments: "./deployments",
    sources: "./src/contracts",
  },
  namedAccounts: {
    admin: {
      default: 0,
    },
    wallet1: {
      default: 1,
    },
    wallet2: {
      default: 2,
    },
    wallet3: {
      default: 3,
    },
    asgard1: {
      default: 4,
    },
    asgard2: {
      default: 5,
    },
    asgard3: {
      default: 6,
    },
    yggdrasil: {
      default: 7,
    },
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
  },
  gasReporter: {
    enabled: false,
    currency: "USD",
  },
  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
  mocha: {
    timeout: 130000,
  },
};
