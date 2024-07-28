#!/bin/sh

set -o pipefail
format_1e8() {
  printf "%.2f\n" "$(jq -n "$1"/100000000 2>/dev/null)" 2>/dev/null | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta'
}

format_int() {
  printf "%.0f\n" "$1" 2>/dev/null | sed ':a;s/\B[0-9]\{3\}\>/,&/;ta'
}

calc_progress() {
  if [ "$1" = "$2" ]; then
    [ "$1" = "0" ] && echo "0.000%" || echo "100.000%"
  elif [ -n "$3" ]; then
    progress="$(echo "scale=6; $3 * 100" | bc 2>/dev/null)" 2>/dev/null && printf "%.3f%%" "$progress" || echo "Error"
  else
    progress="$(echo "scale=6; $1/$2 * 100" | bc 2>/dev/null)" 2>/dev/null && printf "%.3f%%" "$progress" || echo "Error"
  fi
}

API=http://mayanode:1317
MAYANODE_PORT="${MAYANODE_SERVICE_PORT_RPC:-27147}"

THORNODE_ENDPOINT="${THOR_HOST:-thornode-daemon:${THORNODE_SERVICE_PORT_RPC:-27147}}"
BINANCE_ENDPOINT="${BINANCE_HOST:-binance-daemon:${BINANCE_DAEMON_SERVICE_PORT_RPC:-26657}}"
BITCOIN_ENDPOINT="${BTC_HOST:-bitcoin-daemon:${BITCOIN_DAEMON_SERVICE_PORT_RPC:-8332}}"
LITECOIN_ENDPOINT="${LTC_HOST:-litecoin-daemon:${LITECOIN_DAEMON_SERVICE_PORT_RPC:-9332}}"
BITCOIN_CASH_ENDPOINT="${BCH_HOST:-bitcoin-cash-daemon:${BITCOIN_CASH_DAEMON_SERVICE_PORT_RPC:-8332}}"
DOGECOIN_ENDPOINT="${DOGE_HOST:-dogecoin-daemon:${DOGECOIN_DAEMON_SERVICE_PORT_RPC:-22555}}"
DASH_ENDPOINT="${DASH_HOST:-dash-daemon:${DASH_DAEMON_SERVICE_PORT_RPC:-9998}}"
ETHEREUM_ENDPOINT="${ETH_HOST:-http://ethereum-daemon:${ETHEREUM_DAEMON_SERVICE_PORT_RPC:-8545}}"
ETHEREUM_BEACON_ENDPOINT=$(echo "$ETHEREUM_ENDPOINT" | sed 's/:[0-9]*$/:3500/g')
GAIA_ENDPOINT="${GAIA_HOST:-http://gaia-daemon:26657}"
KUJI_ENDPOINT="${KUJI_HOST:-http://kuji-daemon:26657}"
AVALANCHE_ENDPOINT="${AVAX_HOST:-http://avalanche-daemon:9650/ext/bc/C/rpc}"
ARBITRUM_ENDPOINT="${ARB_HOST:-http://arbitrum-daemon:${ARBITRUM_DAEMON_SERVICE_PORT_RPC:-8547}}"

ADDRESS=$(echo "$SIGNER_PASSWD" | mayanode keys show "$SIGNER_NAME" -a --keyring-backend file)
JSON=$(curl -sL --fail -m 10 "$API/mayachain/node/$ADDRESS")

IP=$(echo "$JSON" | jq -r ".ip_address")
VERSION=$(echo "$JSON" | jq -r ".version")
BOND=$(echo "$JSON" | jq -r ".bond")
REWARDS=$(echo "$JSON" | jq -r ".current_award")
SLASH=$(echo "$JSON" | jq -r ".slash_points")
STATUS=$(echo "$JSON" | jq -r ".status")
PREFLIGHT=$(echo "$JSON" | jq -r ".preflight_status")
[ "$VALIDATOR" = "false" ] && IP=$EXTERNAL_IP

if [ "$VALIDATOR" = "true" ]; then
  # calculate BNB chain sync progress
  if [ -z "$BNB_DISABLED" ]; then
    NET=$(curl -sL --fail -m 10 "$BINANCE_ENDPOINT/status" | jq -r ".result.node_info.network")
    if [ "$NET" = "mainnet" ] || [ "$NET" = "stagenet" ]; then # Seeds from https://docs.binance.org/smart-chain/developer/rpc.html
      BNB_PEERS='https://dataseed1.binance.org https://dataseed2.binance.org https://dataseed3.binance.org https://dataseed4.binance.org'
    else
      BNB_PEERS='http://data-seed-pre-0-s1.binance.org http://data-seed-pre-1-s1.binance.org http://data-seed-pre-2-s1.binance.org http://data-seed-pre-0-s3.binance.org http://data-seed-pre-1-s3.binance.org'
    fi
    for BNB_PEER in ${BNB_PEERS}; do
      BNB_HEIGHT=$(curl -sL --fail -m 10 "$BNB_PEER"/status | jq -e -r ".result.sync_info.latest_block_height") || continue
      if [ -z "$BNB_HEIGHT" ]; then continue; fi # Continue if empty height (malformed/bad json reply?)
      break
    done
    BNB_SYNC_HEIGHT=$(curl -sL --fail -m 10 "$BINANCE_ENDPOINT/status" | jq -r ".result.sync_info.index_height")
    BNB_PROGRESS=$(calc_progress "$BNB_SYNC_HEIGHT" "$BNB_HEIGHT")
  fi

  # calculate BTC chain sync progress
  BTC_RESULT=$(curl -sL --fail -m 10 --data-binary '{"jsonrpc": "1.0", "id": "node-status", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://mayachain:password@"$BITCOIN_ENDPOINT")
  BTC_HEIGHT=$(echo "$BTC_RESULT" | jq -r ".result.headers")
  BTC_SYNC_HEIGHT=$(echo "$BTC_RESULT" | jq -r ".result.blocks")
  BTC_PROGRESS=$(echo "$BTC_RESULT" | jq -r ".result.verificationprogress")
  BTC_PROGRESS=$(calc_progress "$BTC_SYNC_HEIGHT" "$BTC_HEIGHT" "$BTC_PROGRESS")

  # calculate LTC chain sync progress
  if [ -z "$LTC_DISABLED" ]; then
    LTC_RESULT=$(curl -sL --fail -m 10 --data-binary '{"jsonrpc": "1.0", "id": "node-status", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://mayachain:password@"$LITECOIN_ENDPOINT")
    LTC_HEIGHT=$(echo "$LTC_RESULT" | jq -r ".result.headers")
    LTC_SYNC_HEIGHT=$(echo "$LTC_RESULT" | jq -r ".result.blocks")
    LTC_PROGRESS=$(echo "$LTC_RESULT" | jq -r ".result.verificationprogress")
    LTC_PROGRESS=$(calc_progress "$LTC_SYNC_HEIGHT" "$LTC_HEIGHT" "$LTC_PROGRESS")
  fi

  ETH_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' -H 'content-type: application/json' "$ETHEREUM_ENDPOINT")
  if [ "$ETH_RESULT" = '{"jsonrpc":"2.0","id":1,"result":false}' ]; then
    ETH_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H 'content-type: application/json' "$ETHEREUM_ENDPOINT")
    ETH_HEIGHT=$(printf "%.0f" "$(echo "$ETH_RESULT" | jq -r ".result")")
    ETH_SYNC_HEIGHT=$ETH_HEIGHT
    ETH_PROGRESS=$(calc_progress "$ETH_SYNC_HEIGHT" "$ETH_HEIGHT")
  elif [ -n "$ETH_RESULT" ]; then
    ETH_HEIGHT=$(printf "%.0f" "$(echo "$ETH_RESULT" | jq -r ".result.highestBlock")")
    ETH_SYNC_HEIGHT=$(printf "%.0f" "$(echo "$ETH_RESULT" | jq -r ".result.currentBlock")")
  else
    ETH_PROGRESS=Error
  fi

  # calculate ETH chain sync progress
  ETH_BEACON_RESULT=$(curl -sL --fail -m 10 "$ETHEREUM_BEACON_ENDPOINT/eth/v1/node/syncing")
  if [ -n "$ETH_BEACON_RESULT" ]; then
    ETH_BEACON_HEIGHT=$(echo "$ETH_BEACON_RESULT" | jq -r "(.data.head_slot|tonumber)+(.data.sync_distance|tonumber)")
    ETH_BEACON_SYNC_HEIGHT=$(echo "$ETH_BEACON_RESULT" | jq -r ".data.head_slot|tonumber")
    ETH_BEACON_PROGRESS=$(calc_progress "$ETH_BEACON_SYNC_HEIGHT" "$ETH_BEACON_HEIGHT")
  else
    ETH_BEACON_PROGRESS=Error
  fi

  # calculate BCH chain sync progress
  if [ -z "$BCH_DISABLED" ]; then
    BCH_RESULT=$(curl -sL --fail -m 10 --data-binary '{"jsonrpc": "1.0", "id": "node-status", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://mayachain:password@"$BITCOIN_CASH_ENDPOINT")
    BCH_HEIGHT=$(echo "$BCH_RESULT" | jq -r ".result.headers")
    BCH_SYNC_HEIGHT=$(echo "$BCH_RESULT" | jq -r ".result.blocks")
    BCH_PROGRESS=$(echo "$BCH_RESULT" | jq -r ".result.verificationprogress")
    BCH_PROGRESS=$(calc_progress "$BCH_SYNC_HEIGHT" "$BCH_HEIGHT" "$BCH_PROGRESS")
  fi

  # calculate DOGE chain sync progress
  DOGE_RESULT=$(curl -sL --fail -m 10 --data-binary '{"jsonrpc": "1.0", "id": "node-status", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://mayachain:password@"$DOGECOIN_ENDPOINT")
  DOGE_HEIGHT=$(echo "$DOGE_RESULT" | jq -r ".result.headers")
  DOGE_SYNC_HEIGHT=$(echo "$DOGE_RESULT" | jq -r ".result.blocks")
  DOGE_PROGRESS=$(echo "$DOGE_RESULT" | jq -r ".result.verificationprogress")
  DOGE_PROGRESS=$(calc_progress "$DOGE_SYNC_HEIGHT" "$DOGE_HEIGHT" "$DOGE_PROGRESS")

  # calculate DASH chain sync progress
  if [ -z "$DASH_DISABLED" ]; then
    DASH_RESULT=$(curl -sL --fail -m 10 --data-binary '{"jsonrpc": "1.0", "id": "node-status", "method": "getblockchaininfo", "params": []}' -H 'content-type: text/plain;' http://mayachain:password@"$DASH_ENDPOINT")
    DASH_HEIGHT=$(echo "$DASH_RESULT" | jq -r ".result.headers")
    DASH_SYNC_HEIGHT=$(echo "$DASH_RESULT" | jq -r ".result.blocks")
    DASH_PROGRESS=$(echo "$DASH_RESULT" | jq -r ".result.verificationprogress")
    DASH_PROGRESS=$(calc_progress "$DASH_SYNC_HEIGHT" "$DASH_HEIGHT" "$DASH_PROGRESS")
  fi

  # calculate Gaia chain sync progress
  GAIA_HEIGHT=$(curl -sL --fail -m 10 https://gaia.ninerealms.com/status | jq -r ".result.sync_info.latest_block_height")
  GAIA_SYNC_HEIGHT=$(curl -sL --fail -m 10 "$GAIA_ENDPOINT/status" | jq -r ".result.sync_info.latest_block_height")
  GAIA_PROGRESS=$(calc_progress "$GAIA_SYNC_HEIGHT" "$GAIA_HEIGHT")

  # calculate Kuji chain sync progress
  if [ -z "$KUJI_DISABLED" ]; then
    KUJI_HEIGHT=$(curl -sL --fail -m 10 https://kuji.mayachain.info/status | jq -r ".result.sync_info.latest_block_height")
    KUJI_SYNC_HEIGHT=$(curl -sL --fail -m 10 "$KUJI_ENDPOINT/status" | jq -r ".result.sync_info.latest_block_height")
    KUJI_PROGRESS=$(calc_progress "$KUJI_SYNC_HEIGHT" "$KUJI_HEIGHT")
  fi

  # calculate AVAX chain sync progress
  AVAX_HEIGHT_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H 'content-type: application/json' "https://api.avax.network/ext/bc/C/rpc")
  AVAX_SYNC_HEIGHT_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H 'content-type: application/json' "$AVALANCHE_ENDPOINT")
  AVAX_HEIGHT=$(printf "%.0f" "$(echo "$AVAX_HEIGHT_RESULT" | jq -r ".result")")
  if [ -n "$AVAX_SYNC_HEIGHT_RESULT" ]; then
    AVAX_SYNC_HEIGHT=$(printf "%.0f" "$(echo "$AVAX_SYNC_HEIGHT_RESULT" | jq -r ".result")")
  else
    AVAX_SYNC_HEIGHT=0
  fi
  AVAX_PROGRESS=$(calc_progress "$AVAX_SYNC_HEIGHT" "$AVAX_HEIGHT")

  # calculate ARB chain sync progress
  if [ -z "$ARB_DISABLED" ]; then
    ARB_HEIGHT_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H 'content-type: application/json' "https://arb1.arbitrum.io/rpc")
    ARB_SYNC_HEIGHT_RESULT=$(curl -X POST -sL --fail -m 10 --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' -H 'content-type: application/json' "$ARBITRUM_ENDPOINT")
    ARB_HEIGHT=$(printf "%d" "$(echo "$ARB_HEIGHT_RESULT" | jq -r ".result")")
    if [ -n "$ARB_SYNC_HEIGHT_RESULT" ]; then
      ARB_SYNC_HEIGHT=$(printf "%.0f" "$(echo "$ARB_SYNC_HEIGHT_RESULT" | jq -r ".result")")
    else
      ARB_SYNC_HEIGHT=0
    fi
    ARB_PROGRESS=$(calc_progress "$ARB_SYNC_HEIGHT" "$ARB_HEIGHT")
  fi

  # calculate THOR chain sync progress
  if [ -z "$THORNODE_DISABLED" ]; then
    # calculate THOR chain sync progress
    THOR_PEER="$THOR_HOST"
    if [ "$NET" = "mainnet" ] || [ "$NET" = "stagenet" ]; then
      THOR_PEER="https://rpc.ninerealms.com"
    fi
    THOR_SYNC_HEIGHT=$(curl -sL --fail -m 10 "$THORNODE_ENDPOINT/status" | jq -r ".result.sync_info.latest_block_height")
    THOR_HEIGHT=$(curl -sL --fail -m 10 "$THOR_PEER/status" | jq -r ".result.sync_info.latest_block_height")
    THOR_PROGRESS=$(printf "%.3f%%" "$(jq -n "$THOR_SYNC_HEIGHT"/"$THOR_HEIGHT"*100 2>/dev/null)" 2>/dev/null) || THOR_PROGRESS=Error
  fi
fi

# calculate MAYA chain sync progress
MAYA_SYNC_HEIGHT=$(curl -sL --fail -m 10 "$CHAIN_RPC/status" | jq -r ".result.sync_info.latest_block_height")
if [ "$PEER" != "" ]; then
  MAYA_HEIGHT=$(curl -sL --fail -m 10 "$PEER:$MAYANODE_PORT/status" | jq -r ".result.sync_info.latest_block_height")
elif [ "$SEEDS" != "" ]; then
  OLD_IFS=$IFS
  IFS=","
  for PEER in $SEEDS; do
    MAYA_HEIGHT=$(curl -sL --fail -m 10 "$PEER:$MAYANODE_PORT/status" | jq -r ".result.sync_info.latest_block_height") || continue
    break
  done
  IFS=$OLD_IFS
else
  MAYA_HEIGHT=$(curl -sL --fail -m 10 "https://mayanode.mayachain.info/blocks/latest" | jq -r ".block.header.height")
fi
MAYA_PROGRESS=$(printf "%.3f%%" "$(jq -n "$MAYA_SYNC_HEIGHT"/"$MAYA_HEIGHT"*100 2>/dev/null)" 2>/dev/null) || MAYA_PROGRESS=Error

cat <<EOF
   __  ________  _____   _  __        __   
  /  |/  /   \ \/ /   | / |/ /__  ___/ /__ 
 / /|_/ / /| |\  / /| |/    / __\/ __ / -_)
/_/  /_/_/ |_|/_/_/ |_/_/|_/\___/\_,_/\__/ 
EOF
echo

if [ "$VALIDATOR" = "true" ]; then
  echo "ADDRESS     $ADDRESS"
  echo "IP          $IP"
  echo "VERSION     $VERSION"
  echo "STATUS      $STATUS"
  echo "BOND        $(format_1e8 "$BOND")"
  echo "REWARDS     $(format_1e8 "$REWARDS")"
  echo "SLASH       $(format_int "$SLASH")"
  echo "PREFLIGHT   $PREFLIGHT"
fi

echo
echo "API         http://$IP:1317/mayachain/doc/"
echo "RPC         http://$IP:$MAYANODE_PORT"
echo "MIDGARD     http://$IP:8080/v2/doc"

# set defaults to avoid failures in math below
MAYA_HEIGHT=${MAYA_HEIGHT:=0}
MAYA_SYNC_HEIGHT=${MAYA_SYNC_HEIGHT:=0}
THOR_HEIGHT=${THOR_HEIGHT:=0}
THOR_SYNC_HEIGHT=${THOR_SYNC_HEIGHT:=0}
BNB_HEIGHT=${BNB_HEIGHT:=0}
BNB_SYNC_HEIGHT=${BNB_SYNC_HEIGHT:=0}
BTC_HEIGHT=${BTC_HEIGHT:=0}
BTC_SYNC_HEIGHT=${BTC_SYNC_HEIGHT:=0}
ETH_HEIGHT=${ETH_HEIGHT:=0}
ETH_SYNC_HEIGHT=${ETH_SYNC_HEIGHT:=0}
LTC_HEIGHT=${LTC_HEIGHT:=0}
LTC_SYNC_HEIGHT=${LTC_SYNC_HEIGHT:=0}
BCH_HEIGHT=${BCH_HEIGHT:=0}
BCH_SYNC_HEIGHT=${BCH_SYNC_HEIGHT:=0}
DOGE_HEIGHT=${DOGE_HEIGHT:=0}
DOGE_SYNC_HEIGHT=${DOGE_SYNC_HEIGHT:=0}
DASH_HEIGHT=${DASH_HEIGHT:=0}
DASH_SYNC_HEIGHT=${DASH_SYNC_HEIGHT:=0}
GAIA_HEIGHT=${GAIA_HEIGHT:=0}
GAIA_SYNC_HEIGHT=${GAIA_SYNC_HEIGHT:=0}
KUJI_HEIGHT=${KUJI_HEIGHT:=0}
KUJI_SYNC_HEIGHT=${KUJI_SYNC_HEIGHT:=0}
ARB_HEIGHT=${ARB_HEIGHT:=0}
ARB_SYNC_HEIGHT=${ARB_SYNC_HEIGHT:=0}

echo
printf "%-10s %-10s %-14s %-10s\n" CHAIN SYNC BEHIND TIP
printf "%-10s %-10s %-14s %-10s\n" MAYA "$MAYA_PROGRESS" "$(format_int $((MAYA_SYNC_HEIGHT - MAYA_HEIGHT)))" "$(format_int "$MAYA_HEIGHT")"
[ "$VALIDATOR" = "true" ] && printf "%-10s %-10s %-14s %-10s\n" THOR "$THOR_PROGRESS" "$(format_int $((THOR_SYNC_HEIGHT - THOR_HEIGHT)))" "$(format_int "$THOR_HEIGHT")"
[ "$VALIDATOR" = "true" ] && printf "%-10s %-10s %-14s %-10s\n" ETH "$ETH_PROGRESS" "$(format_int $((ETH_SYNC_HEIGHT - ETH_HEIGHT)))" "$(format_int "$ETH_HEIGHT")"
[ "$VALIDATOR" = "true" ] && printf "%-18s %-10s %-14s %-10s\n" "ETH (beacon slot)" "$ETH_BEACON_PROGRESS" "$(format_int $((ETH_BEACON_SYNC_HEIGHT - ETH_BEACON_HEIGHT)))" "$(format_int "$ETH_BEACON_HEIGHT")"
[ "$VALIDATOR" = "true" ] && printf "%-10s %-10s %-14s %-10s\n" BTC "$BTC_PROGRESS" "$(format_int $((BTC_SYNC_HEIGHT - BTC_HEIGHT)))" "$(format_int "$BTC_HEIGHT")"
if [ "$VALIDATOR" = "true" ] && [ -z "$BNB_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" BNB "$BNB_PROGRESS" "$(format_int $((BNB_SYNC_HEIGHT - BNB_HEIGHT)))" "$(format_int "$BNB_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$LTC_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" LTC "$LTC_PROGRESS" "$(format_int $((LTC_SYNC_HEIGHT - LTC_HEIGHT)))" "$(format_int "$LTC_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$BCH_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" BCH "$BCH_PROGRESS" "$(format_int $((BCH_SYNC_HEIGHT - BCH_HEIGHT)))" "$(format_int "$BCH_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$DOGE_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" DOGE "$DOGE_PROGRESS" "$(format_int $((DOGE_SYNC_HEIGHT - DOGE_HEIGHT)))" "$(format_int "$DOGE_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$DASH_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" DASH "$DASH_PROGRESS" "$(format_int $((DASH_SYNC_HEIGHT - DASH_HEIGHT)))" "$(format_int "$DASH_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$GAIA_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" GAIA "$GAIA_PROGRESS" "$(format_int $((GAIA_SYNC_HEIGHT - GAIA_HEIGHT)))" "$(format_int "$GAIA_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$KUJI_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" KUJI "$KUJI_PROGRESS" "$(format_int $((KUJI_SYNC_HEIGHT - KUJI_HEIGHT)))" "$(format_int "$KUJI_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$AVAX_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" AVAX "$AVAX_PROGRESS" "$(format_int $((AVAX_SYNC_HEIGHT - AVAX_HEIGHT)))" "$(format_int "$AVAX_HEIGHT")"
fi
if [ "$VALIDATOR" = "true" ] && [ -z "$ARB_DISABLED" ]; then
  printf "%-10s %-10s %-14s %-10s\n" ARB "$ARB_PROGRESS" "$(format_int $((ARB_SYNC_HEIGHT - ARB_HEIGHT)))" "$(format_int "$ARB_HEIGHT")"
fi
exit 0
