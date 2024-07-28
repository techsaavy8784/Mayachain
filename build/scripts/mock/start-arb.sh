#!/bin/sh

check_balance() {
  curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x5E1497dD1f08C87b2d8FE23e9AAB6c1De833D927", "latest"],"id":1}' http://ethereum:8545 | jq -r '.result'
}

# Initial balance check
BALANCE=$(check_balance)
echo "Checking balance..."

# Loop until BALANCE is different from "0x0"
while [ "$BALANCE" = "0x0" ] || [ "$BALANCE" = "" ]; do
  BALANCE=$(check_balance)
  echo "Waiting for balance..."
  if [ "$BALANCE" != "0x0" ] && [ "$BALANCE" != "" ]; then
    break
  fi
  sleep 5
done
echo "Balance changed: $BALANCE"

# Deploy contract in ethereum
deploy --l1conn ws://ethereum:8546 --l1keystore /home/user/l1keystore --sequencerAddress 0xe2148eE53c0755215Df69b2616E552154EdC584f \
  --ownerAddress 0x5E1497dD1f08C87b2d8FE23e9AAB6c1De833D927 --l1DeployAccount 0x5E1497dD1f08C87b2d8FE23e9AAB6c1De833D927 --l1deployment /config/deployment.json \
  --authorizevalidators 10 --wasmrootpath /home/user/target/machines --l1chainid 1337 --l2chainconfig /config/l2_chain_config.json \
  --l2chainname arb-dev-test --l2chaininfo /config/deployed_chain_info.json

# Save the deployed chain info
jq '[.[]]' /config/deployed_chain_info.json >/config/l2_chain_info.json

# Start
nitro \
  --conf.file /config/sequencer_config.json \
  --node.feed.output.enable \
  --node.feed.output.port 9642 \
  --http.api net,web3,eth,txpool,debug,personal \
  --node.seq-coordinator.my-url ws://arbitrum:8548
