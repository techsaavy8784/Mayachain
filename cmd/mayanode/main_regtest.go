//go:build regtest
// +build regtest

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cosmos/cosmos-sdk/types"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"gitlab.com/mayachain/mayanode/app"
	prefix "gitlab.com/mayachain/mayanode/cmd"
	"gitlab.com/mayachain/mayanode/cmd/mayanode/cmd"
	"gitlab.com/mayachain/mayanode/common/cosmos"
)

func main() {
	config := cosmos.GetConfig()
	config.SetBech32PrefixForAccount(prefix.Bech32PrefixAccAddr, prefix.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(prefix.Bech32PrefixValAddr, prefix.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(prefix.Bech32PrefixConsAddr, prefix.Bech32PrefixConsPub)
	config.SetCoinType(prefix.BASEChainCoinType)
	config.SetPurpose(prefix.BASEChainCoinPurpose)
	config.Seal()
	types.SetCoinDenomRegex(func() string {
		return prefix.DenomRegex
	})

	// for coverage data we need to exit main without allowing the server to call os.Exit

	syn := make(chan error)
	go func() {
		rootCmd, _ := cmd.NewRootCmd()
		syn <- svrcmd.Execute(rootCmd, app.DefaultNodeHome())
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGUSR1)
	select {
	case <-sig:
	case err := <-syn:
		if err != nil {
			os.Exit(1)
		}
	}
}
