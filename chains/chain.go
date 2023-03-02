package chains

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Chain struct {
	*OneChainConfig
	*Contracts

	EthClient   *ethclient.Client
	monitor2    *mon2.Monitor
	balAlertCtl chan bool
}

type Chains struct {
	// chainid => Chain
	chains map[uint64]*Chain
	lock   sync.RWMutex
	initWg sync.WaitGroup
}

var chains *Chains

func InitChains() *Chains {
	chains = NewChains()
	return chains
}

func NewChains() *Chains {
	log.Infoln("Initializing chains")
	var configs []*OneChainConfig
	err := viper.UnmarshalKey(FlagMultiChain, &configs)
	if err != nil {
		log.Fatalf("failed to load multichain configs:%v", err)
		return nil
	}
	cs := &Chains{chains: make(map[uint64]*Chain)}
	for _, config := range configs {
		cs.initWg.Add(1)
		go cs.addChain(config)
	}
	cs.initWg.Wait()
	log.Infoln("Finished initializing all chains")
	return cs
}

func (c *Chains) addChain(config *OneChainConfig) {
	ec := newEthClient(config)

	// init monitor
	chainConfig := mon2.PerChainCfg{
		BlkIntv:         time.Duration(config.BlkInterval) * time.Second,
		BlkDelay:        config.BlkDelay,
		MaxBlkDelta:     config.MaxBlkDelta,
		ForwardBlkDelay: config.ForwardBlkDelay,
	}
	db := dal.GetDB()
	mon, err := mon2.NewMonitor(ec, db, chainConfig)
	if err != nil {
		log.Fatalln("failed to create monitor: ", err)
	}

	chain := &Chain{
		OneChainConfig: config,
		Contracts:      NewContracts(ec, config),
		EthClient:      ec,
		monitor2:       mon,
		balAlertCtl:    make(chan bool),
	}
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		c.initWg.Done()
	}()
	c.chains[chain.ChainID] = chain
}

func (c *Chains) GetChain(chid uint64) (*Chain, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	chain, ok := c.chains[chid]
	return chain, ok
}

func GetChain(chid uint64) (*Chain, bool) {
	return chains.GetChain(chid)
}

func GetChainMustExist(chid uint64) *Chain {
	chain, ok := chains.GetChain(chid)
	if !ok {
		log.Panicf("chain %d not found", chid)
	}
	return chain
}

func GetChainIDs() []uint64 {
	var ids []uint64;
	for _, chain := range chains.chains {
		ids = append(ids, chain.ChainID)
	}
	return ids
}

func StartMonitoring(filters contracts.ReceiverContracts, signers map[uint64][]eth.Addr) {
	for _, chain := range chains.chains {
		go chain.startMonitoringEvents(filters)
		go chain.startMonitoringBalance(signers[chain.ChainID])
	}
}

func StopMonitoring() {
	for _, chain := range chains.chains {
		chain.monitor2.Close()
		close(chain.balAlertCtl)
	}
}

func ReloadConfig() error {
	cs := NewChains()
	chains.lock.Lock()
	defer chains.lock.Unlock()
	chains.chains = cs.chains
	return nil
}

func newEthClient(config *OneChainConfig) *ethclient.Client {
	// init eth client
	log.Infof("Dialing eth client for chain %d at %s", config.ChainID, config.Gateway)
	var ec *ethclient.Client
	var err error
	if config.ProxyPort > 0 {
		proxyPort := config.ProxyPort + 600
		if err = endpointproxy.StartProxy(config.Gateway, config.ChainID, proxyPort); err != nil {
			log.Fatalln("can not start proxy for chain:", config.ChainID, "gateway:", config.Gateway, "port:", proxyPort, "err:", err)
		}
		ec, err = ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", proxyPort))
		if err != nil {
			log.Fatalln("dial", config.Gateway, "err:", err)
		}
	} else {
		ec, err = ethclient.Dial(config.Gateway)
		if err != nil {
			log.Fatalln("dial", config.Gateway, "err:", err)
		}
	}
	checkChainID(ec, config.ChainID)
	return ec
}

func checkChainID(ec *ethclient.Client, chainID uint64) {
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", chainID, err)
	}
	if chid.Uint64() != chainID {
		log.Fatalf("chainid mismatch! chainConf has %d but onchain has %d", chainID, chid.Uint64())
	}
}
