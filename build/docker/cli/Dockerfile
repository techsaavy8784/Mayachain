FROM registry.gitlab.com/mayachain/devops/node-launcher:bitcoin-daemon-24.1 as bitcoin
FROM registry.gitlab.com/mayachain/devops/node-launcher:bitcoin-cash-daemon-26.0.0 as bitcoin-cash
FROM registry.gitlab.com/mayachain/devops/node-launcher:dogecoin-daemon-1.14.6 as dogecoin
FROM registry.gitlab.com/mayachain/devops/node-launcher:litecoin-daemon-0.21.1rc1 as litecoin
FROM registry.gitlab.com/mayachain/devops/node-launcher:dash-daemon-20.0.0 as dash
FROM registry.gitlab.com/mayachain/devops/node-launcher:thornode-daemon-mocknet-1.126.0 as thorchain

FROM golang:1.22.3 AS build

RUN apt-get update && \
  apt-get install --no-install-recommends -y protobuf-compiler=3.12.4-1+deb11u1 jq=1.6-1 && \
  rm -rf /var/lib/apt/lists/*

# build geth since the image is alpine based
RUN go install github.com/ethereum/go-ethereum/cmd/geth@v1.11.5

# copy chain clis
COPY --from=bitcoin /usr/local/bin/bitcoin-cli /usr/local/bin/bitcoin-cli
COPY --from=bitcoin-cash /usr/local/bin/bitcoin-cli /usr/local/bin/bitcoin-cash-cli
COPY --from=dogecoin /usr/local/bin/dogecoin-cli /usr/local/bin/dogecoin-cli
COPY --from=litecoin /usr/local/bin/litecoin-cli /usr/local/bin/litecoin-cli
COPY --from=dash /usr/local/bin/dash-cli /usr/local/bin/dash-cli
COPY --from=thorchain /usr/bin/thornode /usr/local/bin/thornode

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TAG=mocknet
RUN make install
