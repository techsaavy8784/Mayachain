package tss

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/thorchain/tss/go-tss/keygen"
	"gitlab.com/thorchain/tss/go-tss/tss"

	"gitlab.com/mayachain/mayanode/bifrost/mayaclient"
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
	"gitlab.com/mayachain/mayanode/x/mayachain/types"
)

// KeyGen is
type KeyGen struct {
	keys           *mayaclient.Keys
	logger         zerolog.Logger
	client         *http.Client
	server         *tss.TssServer
	bridge         mayaclient.MayachainBridge
	currentVersion semver.Version
	lastCheck      time.Time
}

// NewTssKeyGen create a new instance of TssKeyGen which will look after TSS key stuff
func NewTssKeyGen(keys *mayaclient.Keys, server *tss.TssServer, bridge mayaclient.MayachainBridge) (*KeyGen, error) {
	if keys == nil {
		return nil, fmt.Errorf("keys is nil")
	}
	return &KeyGen{
		keys:   keys,
		logger: log.With().Str("module", "tss_keygen").Logger(),
		client: &http.Client{
			Timeout: time.Second * 130,
		},
		server: server,
		bridge: bridge,
	}, nil
}

func (kg *KeyGen) getVersion() semver.Version {
	requestTime := time.Now()
	if !kg.currentVersion.Equals(semver.Version{}) && requestTime.Sub(kg.lastCheck).Seconds() < constants.MayachainBlockTime.Seconds() {
		return kg.currentVersion
	}
	version, err := kg.bridge.GetMayachainVersion()
	if err != nil {
		kg.logger.Err(err).Msg("fail to get current mayachain version")
		return kg.currentVersion
	}
	kg.currentVersion = version
	kg.lastCheck = requestTime
	return kg.currentVersion
}

func (kg *KeyGen) GenerateNewKey(keygenBlockHeight int64, pKeys common.PubKeys) (pk common.PubKeySet, blame types.Blame, err error) {
	// No need to do key gen
	if len(pKeys) == 0 {
		return common.EmptyPubKeySet, types.Blame{}, nil
	}

	// add some logging
	defer func() {
		if blame.IsEmpty() {
			kg.logger.Info().Int64("height", keygenBlockHeight).Str("pubkey", pk.String()).Msg("tss keygen results success")
		} else {
			blames := make([]string, len(blame.BlameNodes))
			for i := range blame.BlameNodes {
				var pk common.PubKey
				pk, err = common.NewPubKey(blame.BlameNodes[i].Pubkey)
				if err != nil {
					kg.logger.Error().Err(err).Int64("height", keygenBlockHeight).Str("pubkey", blame.BlameNodes[i].Pubkey).Msg("tss keygen results error")
					continue
				}
				var acc cosmos.AccAddress
				acc, err = pk.GetThorAddress()
				if err != nil {
					kg.logger.Error().Err(err).Int64("height", keygenBlockHeight).Str("pubkey", pk.String()).Msg("tss keygen results error")
					continue
				}
				blames[i] = acc.String()
			}
			sort.Strings(blames)
			kg.logger.Info().Int64("height", keygenBlockHeight).Str("pubkey", pk.String()).Str("round", blame.Round).Str("blames", strings.Join(blames, ", ")).Str("reason", blame.FailReason).Msg("tss keygen results blame")
		}
	}()

	var keys []string
	for _, item := range pKeys {
		keys = append(keys, item.String())
	}
	keyGenReq := keygen.Request{
		Keys: keys,
	}
	currentVersion := kg.getVersion()
	keyGenReq.Version = currentVersion.String()

	// Use the churn try's block to choose the same leader for every node in an Asgard,
	// since a successful keygen requires every node in the Asgard to take part.
	keyGenReq.BlockHeight = keygenBlockHeight

	ch := make(chan bool, 1)
	defer close(ch)
	timer := time.NewTimer(30 * time.Minute)
	defer timer.Stop()

	var resp keygen.Response
	go func() {
		resp, err = kg.server.Keygen(keyGenReq)
		ch <- true
	}()

	select {
	case <-ch:
		// do nothing
	case <-timer.C:
		panic("tss keygen timeout")
	}

	// copy blame to our own struct
	blame = types.Blame{
		FailReason: resp.Blame.FailReason,
		IsUnicast:  resp.Blame.IsUnicast,
		Round:      resp.Blame.Round,
		BlameNodes: make([]types.Node, len(resp.Blame.BlameNodes)),
	}
	for i, n := range resp.Blame.BlameNodes {
		blame.BlameNodes[i].Pubkey = n.Pubkey
		blame.BlameNodes[i].BlameData = n.BlameData
		blame.BlameNodes[i].BlameSignature = n.BlameSignature
	}

	if err != nil {
		// the resp from kg.server.Keygen will not be nil
		if blame.IsEmpty() {
			blame.FailReason = err.Error()
		}
		return common.EmptyPubKeySet, blame, fmt.Errorf("fail to keygen,err:%w", err)
	}

	cpk, err := common.NewPubKey(resp.PubKey)
	if err != nil {
		return common.EmptyPubKeySet, blame, fmt.Errorf("fail to create common.PubKey,%w", err)
	}

	// TODO later on THORNode need to have both secp256k1 key and ed25519
	return common.NewPubKeySet(cpk, cpk), blame, nil
}
