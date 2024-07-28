#!/bin/sh

deploy --l1conn ws://ethereum:8546 --l1keystore /home/user/l1keystore --sequencerAddress 0xe2148eE53c0755215Df69b2616E552154EdC584f \
  --ownerAddress 0xe2148eE53c0755215Df69b2616E552154EdC584f --l1DeployAccount 0xe2148eE53c0755215Df69b2616E552154EdC584f --l1deployment /config/deployment.json \
  --authorizevalidators 10 --wasmrootpath /home/user/target/machines --l1chainid 1337 --l2chainconfig /config/l2_chain_config.json \
  --l2chainname arb-dev-test --l2chaininfo /config/deployed_chain_info.json
