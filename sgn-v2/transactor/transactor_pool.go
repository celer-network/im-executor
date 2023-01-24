package transactor

import (
	"sync/atomic"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type TransactorPool struct {
	transactors       []*Transactor
	index             uint64
	homeDir           string
	chainID           string
	legacyAmino       *codec.LegacyAmino
	cdc               codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry
}

func NewTransactorPool(homeDir, chainID string, legacyAmino *codec.LegacyAmino, cdc codec.Codec, interfaceRegistry codectypes.InterfaceRegistry) *TransactorPool {
	return &TransactorPool{
		transactors:       []*Transactor{},
		index:             0,
		homeDir:           homeDir,
		chainID:           chainID,
		legacyAmino:       legacyAmino,
		cdc:               cdc,
		interfaceRegistry: interfaceRegistry,
	}
}

// Add transactors to the pool
func (tp *TransactorPool) AddTransactor(transactor *Transactor) {
	tp.transactors = append(tp.transactors, transactor)
}

// Add transactors to the pool
func (tp *TransactorPool) AddTransactors(nodeURI, passphrase string, addrs []string) error {
	return tp.AddTransactorsWithGrpc(nodeURI, passphrase, "", addrs)
}

// Add transactors to the pool
func (tp *TransactorPool) AddTransactorsWithGrpc(nodeURI, passphrase, grpcUrl string, addrs []string) error {
	var transactors []*Transactor
	for _, addr := range addrs {
		transactor, err := NewTransactorWithGrpc(tp.homeDir, tp.chainID, nodeURI, addr, passphrase, grpcUrl, tp.legacyAmino, tp.cdc, tp.interfaceRegistry)
		if err != nil {
			return err
		}
		// transactor.Run(0)

		transactors = append(transactors, transactor)
	}

	tp.transactors = append(tp.transactors, transactors...)
	return nil
}

// Get a transactor from the pool
func (tp *TransactorPool) GetTransactor() *Transactor {
	if len(tp.transactors) == 0 {
		return nil
	}

	transactor := tp.transactors[tp.index%uint64(len(tp.transactors))]
	atomic.AddUint64(&tp.index, 1)
	return transactor
}
