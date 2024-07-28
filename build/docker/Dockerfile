#
# BASEChain
#

#
# Build

FROM registry.gitlab.com/mayachain/mayanode:builder-v5@sha256:b3a025996892ec307a62a0e0da054856fa81aa4747e78a3a7e58253bc4228b7a as build
ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TAG=mainnet
ARG NET=mainnet
RUN make install

#
# Main
#
FROM registry.gitlab.com/mayachain/mayanode:runner-base-v1@sha256:8fc433efa3e853b59cafe920654705e37c94228ae99cba6892933a61594c11f3
# Copy the compiled binaries over.
COPY --from=build /go/bin/generate /go/bin/mayanode /go/bin/bifrost /usr/bin/

COPY build/scripts /scripts

# default to mainnet
ARG TAG=mainnet
ENV NET=$TAG

# default to fullnode
CMD ["/scripts/fullnode.sh"]
