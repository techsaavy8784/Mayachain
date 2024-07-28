#!/bin/sh

set -o pipefail

add_account() {
  ADDRS=$(jq --arg ADDRESS "$1" '.app_state.auth.accounts[] | select(.address == $ADDRESS) .address' <~/.mayanode/config/genesis.json)

  if [ -z "$ADDRS" ]; then
    #If account doesn't exist, create account with asset
    jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.auth.accounts += [{
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": $ADDRESS,
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
      }]' <~/.mayanode/config/genesis.json >/tmp/genesis.json
    # "coins": [ { "denom": $ASSET, "amount": $AMOUNT } ],
    mv /tmp/genesis.json ~/.mayanode/config/genesis.json

    jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.bank.balances += [{
          "address": $ADDRESS,
          "coins": [ { "denom": $ASSET, "amount": $AMOUNT } ],
      }]' <~/.mayanode/config/genesis.json >/tmp/genesis.json
    mv /tmp/genesis.json ~/.mayanode/config/genesis.json
  else
    #If account exist, add balance
    PREV_AMOUNT=$(jq --arg ADDRESS "$1" --arg ASSET "$2" '.app_state.bank.balances[] | select(.address == $ADDRESS) .coins[] | select(.denom == $ASSET) .amount' <~/.mayanode/config/genesis.json)
    if [ -z "$PREV_AMOUNT" ]; then
      # Add new balance to address from non-exiting asset
      jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '.app_state.bank.balances = [(
        .app_state.bank.balances[] | select(.address == $ADDRESS) .coins += [{
        "denom": $ASSET,
        "amount": $AMOUNT
        }])]' <~/.mayanode/config/genesis.json >/tmp/genesis.json
      mv /tmp/genesis.json ~/.mayanode/config/genesis.json
    else
      # Add balance to address from existing asset
      jq --arg ADDRESS "$1" --arg ASSET "$2" --arg AMOUNT "$3" '(.app_state.bank.balances[] | select(.address == $ADDRESS)).coins = [
        .app_state.bank.balances[] | select(.address == $ADDRESS).coins[] | select(.denom == $ASSET).amount += $AMOUNT
        ]' <~/.mayanode/config/genesis.json >/tmp/genesis.json
      mv /tmp/genesis.json ~/.mayanode/config/genesis.json
    fi
  fi
}

deploy_evm_contracts() {
  for CHAIN in ETH AVAX ARB; do
    if [ "$CHAIN" = "ARB" ]; then
      wait_arbitrum
    fi
    # deploy contract and get address from output
    echo "Deploying $CHAIN contracts"
    if ! python3 scripts/evm/evm-tool.py --chain "$CHAIN" --rpc "$(eval echo "\$${CHAIN}_HOST")" --action deploy >/tmp/evm-tool.log 2>&1; then
      cat /tmp/evm-tool.log && exit 1
    fi
    cat /tmp/evm-tool.log
    CONTRACT=$(grep </tmp/evm-tool.log "Router Contract Address" | awk '{print $NF}')

    # add contract address to genesis
    echo "$CHAIN Contract Address: $CONTRACT"
    jq --arg CHAIN "$CHAIN" --arg CONTRACT "$CONTRACT" \
      '.app_state.mayachain.chain_contracts += [{"chain": $CHAIN, "router": $CONTRACT}]' \
      ~/.mayanode/config/genesis.json >/tmp/genesis.json
    mv /tmp/genesis.json ~/.mayanode/config/genesis.json
  done
}

gen_bnb_address() {
  if [ ! -f ~/.bond/private_key.txt ]; then
    echo "Generating BNB address"
    mkdir -p ~/.bond
    # because the generate command can get API rate limited, THORNode may need to retry
    n=0
    until [ $n -ge 60 ]; do
      generate >/tmp/bnb && break
      n=$((n + 1))
      sleep 1
    done
    ADDRESS=$(grep </tmp/bnb MASTER= | awk -F= '{print $NF}')
    echo "$ADDRESS" >~/.bond/address.txt
    BINANCE_PRIVATE_KEY=$(grep </tmp/bnb MASTER_KEY= | awk -F= '{print $NF}')
    echo "$BINANCE_PRIVATE_KEY" >/root/.bond/private_key.txt
    PUBKEY=$(grep </tmp/bnb MASTER_PUBKEY= | awk -F= '{print $NF}')
    echo "$PUBKEY" >/root/.bond/pubkey.txt
    MNEMONIC=$(grep </tmp/bnb MASTER_MNEMONIC= | awk -F= '{print $NF}')
    echo "$MNEMONIC" >/root/.bond/mnemonic.txt
  fi
}

wait_arbitrum() {
  echo "Waiting for Arbitrum..."
  while true; do
    nc -z arbitrum 8547 && break
    sleep 5
  done
  echo "Arbitrum initialized"
}
