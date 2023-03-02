package accounts

import (
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/sgn-v2/eth"
)

type Account struct {
	*AccountConfig
	Address     eth.Addr
	transactors map[uint64]*ethutils.Transactor // chainID to transactor
	signers     map[uint64]ethutils.Signer
}

func newAccount(conf *AccountConfig) *Account {
	log.Infof("initializing account '%s'", conf.ID)
	transactors := make(map[uint64]*ethutils.Transactor)
	signers := make(map[uint64]ethutils.Signer)
	address := eth.ZeroAddr
	for _, id := range chains.GetChainIDs() {
		chain := chains.GetChainMustExist(id)
		if _, ok := transactors[chain.ChainID]; !ok {
			transactor, signer := newTransactor(chain, conf.Keystore, conf.Passphrase)
			transactors[chain.ChainID] = transactor
			signers[chain.ChainID] = signer
			address = transactor.Address()
		}
	}

	err := conf.SenderGroups.Validate()
	if err != nil {
		log.Fatalf("cannot initialize account: %s", err.Error())
	}
	log.Infof("account '%s' sender groups:\n%s", conf.ID, conf.SenderGroups)
	err = conf.ReceiverContracts.Validate()
	if err != nil {
		log.Fatalf("cannot initialize account: %s", err.Error())
	}
	log.Infof("account %s receiver contracts:\n%s", conf.ID, conf.ReceiverContracts)
	log.Infof("initialized account '%s'", conf.ID)
	return &Account{
		AccountConfig: conf,
		Address:       address,
		transactors:   transactors,
		signers:       signers,
	}
}

func (a *Account) Transactor(chid uint64) (*ethutils.Transactor, bool) {
	transactor, ok := a.transactors[chid]
	return transactor, ok
}

func (a *Account) Signer(chid uint64) (ethutils.Signer, bool) {
	signer, ok := a.signers[chid]
	return signer, ok
}

func (a *Account) ReceiverContract(receiver *contracts.ContractAddress) (*contracts.ReceiverContract, bool) {
	for _, c := range a.ReceiverContracts {
		if c.ContractAddress().Equal(receiver) {
			return c, true
		}
	}
	return nil, false
}

func (a *Account) IsSenderAllowed(sender, receiver *contracts.ContractAddress) bool {
	if sender.GetAddress() == eth.ZeroAddr {
		// zero addr sender means the it's a refund msg
		return true
	}
	recvContract, ok := a.ReceiverContract(receiver)
	if !ok {
		log.Warnf("unable to check sender, receiver contract %s not found", receiver)
		return true
	}
	// if no allowed sender groups are configured, it is defaulted to allowed
	if len(recvContract.AllowSenderGroups) == 0 {
		return true
	}
	for _, group := range recvContract.AllowSenderGroups {
		allowedSenders := a.SenderGroups.AllowedSenders(group)
		if allowedSenders.Has(sender) {
			return true
		}
	}
	return false
}

func newTransactor(chain *chains.Chain, signerKey, signerPass string) (*ethutils.Transactor, ethutils.Signer) {
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, new(big.Int).SetUint64(chain.ChainID))
	if err != nil {
		log.Fatalln("CreateSigner err:", err)
	}
	return ethutils.NewTransactorByExternalSigner(
		addr,
		signer,
		chain.EthClient,
		big.NewInt(int64(chain.ChainID)),
		ethutils.WithBlockDelay(chain.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(chain.BlkInterval)*time.Second),
		ethutils.WithAddGasEstimateRatio(chain.AddGasEstimateRatio),
		ethutils.WithGasLimit(chain.GasLimit),
		ethutils.WithAddGasGwei(chain.AddGasGwei),
		ethutils.WithMaxGasGwei(chain.MaxGasGwei),
		ethutils.WithMinGasGwei(chain.MinGasGwei),
		ethutils.WithMaxFeePerGasGwei(chain.MaxFeePerGasGwei),
		ethutils.WithMaxPriorityFeePerGasGwei(chain.MaxPriorityFeePerGasGwei),
	), signer
}
