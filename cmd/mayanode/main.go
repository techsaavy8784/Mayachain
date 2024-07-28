//go:build !regtest
// +build !regtest

package main

import (
	"os"

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

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome()); err != nil {
		os.Exit(1)
	}
}
