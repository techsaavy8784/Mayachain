//go:build !testnet && !mocknet
// +build !testnet,!mocknet

package bitcoincash

import (
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
)

func (s *BitcoinCashSuite) TestGetAddress(c *C) {
	pubkey := common.PubKey("mayapub1addwnpepqt7qug8vk9r3saw8n4r803ydj2g3dqwx0mvq5akhnze86fc536xcyvg4c5e")
	addr := s.client.GetAddress(pubkey)
	c.Assert(addr, Equals, "qpfztpuwwujkvvenjm7mg9d6mzqkmqwshv07z34njm")
}
