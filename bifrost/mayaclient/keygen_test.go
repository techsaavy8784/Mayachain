package mayaclient

import (
	"net/http"
	"net/http/httptest"
	"strings"

	ckeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/config"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

type KeygenSuite struct {
	server  *httptest.Server
	bridge  *mayachainBridge
	cfg     config.BifrostClientConfiguration
	fixture string
	kb      ckeys.Keyring
}

var _ = Suite(&KeygenSuite{})

func (s *KeygenSuite) SetUpSuite(c *C) {
	s.server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.RequestURI, KeygenEndpoint) {
			httpTestHandler(c, rw, s.fixture)
		}
	}))

	s.cfg, _, s.kb = SetupMayachainForTest(c)
	s.cfg.ChainHost = s.server.Listener.Addr().String()
	var err error
	bridge, err := NewMayachainBridge(s.cfg, GetMetricForTest(c), NewKeysWithKeybase(s.kb, s.cfg.SignerName, s.cfg.SignerPasswd))
	var ok bool
	s.bridge, ok = bridge.(*mayachainBridge)
	c.Assert(ok, Equals, true)
	s.bridge.httpClient.RetryMax = 1
	c.Assert(err, IsNil)
	c.Assert(s.bridge, NotNil)
}

func (s *KeygenSuite) TearDownSuite(c *C) {
	s.server.Close()
}

func (s *KeygenSuite) TestGetKeygen(c *C) {
	s.fixture = "../../test/fixtures/endpoints/keygen/template.json"

	// GENERATE SIGNATURE
	// block := types.NewKeygenBlock(1718)
	// members := []string{
	// 	"tmayapub1addwnpepq2kdyjkm6y9aa3kxl8wfaverka6pvkek2ygrmhx6sj3ec6h0fegwsuewymv",
	// 	"tmayapub1addwnpepqt7qug8vk9r3saw8n4r803ydj2g3dqwx0mvq5akhnze86fc536xcy7cau6l",
	// }
	// keygen, err := types.NewKeygen(1718, members, types.KeygenType_AsgardKeygen)
	// keygen.ID = common.TxID("FEDA8BEDB84117C3EF6BEDA1A2639C11D73724AD0E85268E86CADEA13089E400")
	// keygen.Members = members
	// c.Assert(err, IsNil)
	// block.Keygens = append(block.Keygens, keygen)
	// buf, err := json.Marshal(block)
	// c.Assert(err, IsNil)
	// sig, _, err := s.kb.Sign("mayachain", buf)
	// c.Assert(err, IsNil)
	// fmt.Printf("Sig: %s\n", base64.StdEncoding.EncodeToString(sig))
	// fmt.Printf("KEYGEN1: %+v\n", block)

	pk := types.GetRandomPubKey()
	expectedKey, err := common.NewPubKey("tmayapub1addwnpepq2kdyjkm6y9aa3kxl8wfaverka6pvkek2ygrmhx6sj3ec6h0fegwsuewymv")
	c.Assert(err, IsNil)
	keygenBlock, err := s.bridge.GetKeygenBlock(1718, pk.String())
	c.Assert(err, IsNil)
	c.Assert(keygenBlock, NotNil)
	c.Assert(keygenBlock.Height, Equals, int64(1718))
	c.Assert(keygenBlock.Keygens[0].GetMembers()[0], Equals, expectedKey)
}
