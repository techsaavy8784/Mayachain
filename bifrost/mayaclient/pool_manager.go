package mayaclient

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	stypes "gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// PoolManager provide all the functionalities need to deal with pool
type PoolManager interface {
	GetValue(source, target common.Asset, amount cosmos.Uint) (cosmos.Uint, error)
}

// PoolMgr implement PoolManager interface
type PoolMgr struct {
	bridge    MayachainBridge
	logger    zerolog.Logger
	lastCheck time.Time
	lock      *sync.Mutex
	pools     stypes.Pools
}

// NewPoolMgr create a new instance of PoolMgr
func NewPoolMgr(bridge MayachainBridge) *PoolMgr {
	return &PoolMgr{
		bridge: bridge,
		logger: log.With().Str("module", "pool_mgr").Logger(),
		lock:   &sync.Mutex{},
	}
}

func (pm *PoolMgr) updatePool() {
	pm.lock.Lock()
	defer pm.lock.Unlock()
	pools, err := pm.bridge.GetPools()
	if err != nil {
		pm.logger.Err(err).Msgf("fail to get pool: %s", err)
		return
	}
	pm.pools = pools
}

func (pm *PoolMgr) getPool(asset common.Asset) stypes.Pool {
	duration := time.Since(pm.lastCheck)
	if duration > constants.MayachainBlockTime {
		pm.updatePool()
		pm.lastCheck = time.Now()
	}

	for _, p := range pm.pools {
		if p.Asset.Equals(asset) {
			return p
		}
	}
	return stypes.Pool{}
}

// GetValue is to convert source asset into target , and return the value
// for example, we could ask PoolManager, how much TKN asset in ETH?
func (pm *PoolMgr) GetValue(source, target common.Asset, amount cosmos.Uint) (cosmos.Uint, error) {
	sourcePool := pm.getPool(source)
	if sourcePool.IsEmpty() {
		return cosmos.ZeroUint(), fmt.Errorf("pool:%s doesn't exist", source)
	}
	runeValue := sourcePool.AssetValueInRune(amount)
	destPool := pm.getPool(target)
	if destPool.IsEmpty() {
		return cosmos.ZeroUint(), fmt.Errorf("pool:%s doesn't exist", destPool)
	}
	return destPool.RuneValueInAsset(runeValue), nil
}
