package svc

import (
	"github.com/celer-network/im-executor/accounts"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn"
)

type ExecutorService struct {
	db       *dal.DB
	sgn      *sgn.SgnClient
	chains   *chains.Chains
	accounts *accounts.Accounts
}

func NewExecutorService(db *dal.DB, sgnClient *sgn.SgnClient, chains *chains.Chains, accs *accounts.Accounts) *ExecutorService {
	return &ExecutorService{
		db:       db,
		sgn:      sgnClient,
		chains:   chains,
		accounts: accs,
	}
}
