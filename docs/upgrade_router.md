<!-- markdownlint-disable MD024 -->

# What is router

On EVM based chain, bifrost rely on Router to emit correct events to determine what had happened, all inbound/outbound transactions go through a smart contract, we call it Router. Current Router is on V4. Each connected EVM-based ETH forks have a router (currently Ethereum and Arbitrum).

Router contract hold all ERC20 assets, but not the native asset (e.g. `ETH` or `ARB`). The native assets will be sent to asgard address directly.

# Where is Router code?

https://gitlab.com/mayachain/ethereum/eth-router , if you need to make changes to this router, please raise a PR in [this repository](https://gitlab.com/mayachain/ethereum/eth-router)

# How to upgrade Router?

**Note:** Newer version router needs to be compatible with old router.

## What you can do?

- You can add new functions, new events

## What you can't do?

- Don't change existing function signature , Don't add parameter , don't remove parameter , don't change return value etc.
- Don't change events , don't add new fields , don't remove fields

## Router upgrade procedure

Before router upgrade , make sure you already make relevant changes in mayanode repo. Replace `<chain>` in the below variables with the lowercase, shortened chain identifier (e.g. `eth`, `arb`).

- New router has been deployed , and the router address has been updated. `<chain>OldRouter` is your current router address, `<chain>NewRouter` is your new router address
  - [Mocknet](https://gitlab.com/mayachain/mayanode/-/blob/develop/x/mayachain/router_upgrade_info_mocknet.go)
  - [Stagenet](https://gitlab.com/mayachain/mayanode/-/blob/develop/x/mayachain/router_upgrade_info_stagenet.go)
  - [Mainnet](https://gitlab.com/mayachain/mayanode/-/blob/develop/x/mayachain/router_upgrade_info.go)

Before upgrade , make sure the network is healthy , all active nodes / standby nodes are online. If some nodes are not healthy , bifrost are not online it will cause the node's vault in a bad state

## Detail upgrade procedure

Replace `<CHAIN>` in each Mimir key with capitalized, shortened chain identifier (e.g. `ETH`, `ARB`)

1. Set admin mimir `ChurnInterval` -> `432000` to stop churn
2. Set admin mimir `StopSolvencyCheck<CHAIN>` -> `1` to stop Solvency checker on `CHAIN`, this will make sure the migration fund will not cause solvency checker to halt the chain
3. Set admin mimir `MimirUpgradeContract<CHAIN>` -> `1` to update the router
4. Set admin mimir `ChurnInterval` -> `43200`
5. Wait a churn to kick off , and make sure funds have been migrated from older router to new router. And vault retired successfully
6. Set admin mimir `StopSolvencyCheck<CHAIN>` -> `0` to resume solvency checker on `CHAIN`
