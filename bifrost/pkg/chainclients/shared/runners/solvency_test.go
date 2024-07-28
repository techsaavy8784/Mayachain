package runners

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	ckeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	. "gopkg.in/check.v1"

	"gitlab.com/mayachain/mayanode/bifrost/mayaclient"
	"gitlab.com/mayachain/mayanode/bifrost/metrics"
	"gitlab.com/mayachain/mayanode/cmd"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/config"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

func TestPackage(t *testing.T) { TestingT(t) }

type SolvencyTestSuite struct {
	sp   *DummySolvencyCheckProvider
	m    *metrics.Metrics
	cfg  config.BifrostClientConfiguration
	keys *mayaclient.Keys
}

var _ = Suite(&SolvencyTestSuite{})

func (s *SolvencyTestSuite) SetUpSuite(c *C) {
	sp := &DummySolvencyCheckProvider{}
	s.sp = sp

	m, _ := metrics.NewMetrics(config.BifrostMetricsConfiguration{
		Enabled:      false,
		ListenPort:   9090,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		Chains:       common.Chains{common.BNBChain},
	})
	s.m = m

	cfg := config.BifrostClientConfiguration{
		ChainID:         "thorchain",
		ChainHost:       "localhost",
		SignerName:      "bob",
		SignerPasswd:    "password",
		ChainHomeFolder: ".",
	}
	kb := ckeys.NewInMemory()
	_, _, err := kb.NewMnemonic(cfg.SignerName, ckeys.English, cmd.BASEChainHDPath, cfg.SignerPasswd, hd.Secp256k1)
	c.Assert(err, IsNil)
	s.cfg = cfg
	s.keys = mayaclient.NewKeysWithKeybase(kb, cfg.SignerName, cfg.SignerPasswd)

	c.Assert(err, IsNil)
}

func (s *SolvencyTestSuite) TestSolvencyCheck(c *C) {
	mimirMap := map[string]int{
		"HaltBNBChain":         0,
		"SolvencyHaltBNBChain": 0,
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Logf("================>:%s", r.RequestURI)
		if strings.HasPrefix(r.RequestURI, mayaclient.MimirEndpoint) {
			parts := strings.Split(r.RequestURI, "/key/")
			mimirKey := parts[1]

			mimirValue := 0
			if val, found := mimirMap[mimirKey]; found {
				mimirValue = val
			}

			if _, err := w.Write([]byte(strconv.Itoa(mimirValue))); err != nil {
				c.Error(err)
			}
		}
	})

	server := httptest.NewServer(h)
	defer server.Close()
	bridge, _ := mayaclient.NewMayachainBridge(config.BifrostClientConfiguration{
		ChainID:         "thorchain",
		ChainHost:       server.Listener.Addr().String(),
		ChainRPC:        server.Listener.Addr().String(),
		SignerName:      "bob",
		SignerPasswd:    "password",
		ChainHomeFolder: ".",
	}, s.m, s.keys)

	stopchan := make(chan struct{})
	wg := &sync.WaitGroup{}

	// Happy path, shouldn't check solvency if nothing halted (chain clients will report solvency)
	s.sp.ResetChecks()
	wg.Add(1)
	go SolvencyCheckRunner(common.BNBChain, s.sp, bridge, stopchan, wg, constants.MayachainBlockTime)
	time.Sleep(time.Second * 6)

	c.Assert(s.sp.ShouldReportSolvencyRan, Equals, false)
	c.Assert(s.sp.ReportSolvencyRun, Equals, false)

	// Admin halted, still don't check solvency
	mimirMap["HaltBNBChain"] = 1
	s.sp.ResetChecks()
	wg.Add(1)
	go SolvencyCheckRunner(common.BNBChain, s.sp, bridge, stopchan, wg, constants.MayachainBlockTime)
	time.Sleep(time.Second * 6)

	c.Assert(s.sp.ShouldReportSolvencyRan, Equals, false)
	c.Assert(s.sp.ReportSolvencyRun, Equals, false)

	// Double-spend check halted chain client, check solvency here
	mimirMap["HaltBNBChain"] = 10
	s.sp.ResetChecks()
	wg.Add(1)
	go SolvencyCheckRunner(common.BNBChain, s.sp, bridge, stopchan, wg, constants.MayachainBlockTime)
	time.Sleep(time.Second * 6)

	c.Assert(s.sp.ShouldReportSolvencyRan, Equals, true)
	c.Assert(s.sp.ReportSolvencyRun, Equals, true)
	mimirMap["HaltBNBChain"] = 0

	// Solvency halted chain, need to report solvency here as chain client is paused
	mimirMap["SolvencyHaltBNBChain"] = 1
	s.sp.ResetChecks()
	wg.Add(1)
	go SolvencyCheckRunner(common.BNBChain, s.sp, bridge, stopchan, wg, constants.MayachainBlockTime)
	time.Sleep(time.Second * 6)

	c.Assert(s.sp.ShouldReportSolvencyRan, Equals, true)
	c.Assert(s.sp.ReportSolvencyRun, Equals, true)
}

// Mock SolvencyCheckProvider
type DummySolvencyCheckProvider struct {
	ShouldReportSolvencyRan bool
	ReportSolvencyRun       bool
}

func (d *DummySolvencyCheckProvider) ResetChecks() {
	d.ShouldReportSolvencyRan = false
	d.ReportSolvencyRun = false
}

func (d *DummySolvencyCheckProvider) GetHeight() (int64, error) {
	return 0, nil
}

func (d *DummySolvencyCheckProvider) ShouldReportSolvency(height int64) bool {
	d.ShouldReportSolvencyRan = true
	return true
}

func (d *DummySolvencyCheckProvider) ReportSolvency(height int64) error {
	d.ReportSolvencyRun = true
	return nil
}

func (s *SolvencyTestSuite) TestIsVaultSolvent(c *C) {
	vault := types.Vault{
		BlockHeight: 1,
		PubKey:      types.GetRandomPubKey(),
		Coins: common.NewCoins(
			common.NewCoin(common.ETHAsset, cosmos.NewUint(102400000000)),
		),
		Type:   types.VaultType_AsgardVault,
		Status: types.VaultStatus_ActiveVault,
	}
	acct := common.Account{
		Sequence:      0,
		AccountNumber: 0,
		Coins:         common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(102400000000))),
	}
	c.Assert(IsVaultSolvent(acct, vault, cosmos.NewUint(0)), Equals, true)
	acct = common.Account{
		Sequence:      0,
		AccountNumber: 0,
		Coins:         common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(102305000000))),
	}
	c.Assert(IsVaultSolvent(acct, vault, cosmos.NewUint(80000*120)), Equals, true)
	acct = common.Account{
		Sequence:      0,
		AccountNumber: 0,
		Coins:         common.NewCoins(common.NewCoin(common.ETHAsset, cosmos.NewUint(102205000000))),
	}
	c.Assert(IsVaultSolvent(acct, vault, cosmos.NewUint(80000*120)), Equals, false)
}
