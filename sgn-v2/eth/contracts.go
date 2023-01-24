package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type StakingContract struct {
	*Staking
	Address Addr
}

func NewStakingContract(address Addr, client *ethclient.Client) (*StakingContract, error) {
	staking, err := NewStaking(address, client)
	if err != nil {
		return nil, err
	}
	return &StakingContract{
		Staking: staking,
		Address: address,
	}, nil
}

func (c *StakingContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *StakingContract) GetABI() string {
	return StakingABI
}

type SgnContract struct {
	*SGN
	Address Addr
}

func NewSgnContract(address Addr, client *ethclient.Client) (*SgnContract, error) {
	sgn, err := NewSGN(address, client)
	if err != nil {
		return nil, err
	}
	return &SgnContract{
		SGN:     sgn,
		Address: address,
	}, nil
}

func (c *SgnContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *SgnContract) GetABI() string {
	return SGNABI
}

type StakingRewardContract struct {
	*StakingReward
	Address Addr
}

func NewStakingRewardContract(address Addr, client *ethclient.Client) (*StakingRewardContract, error) {
	stakingReward, err := NewStakingReward(address, client)
	if err != nil {
		return nil, err
	}
	return &StakingRewardContract{
		StakingReward: stakingReward,
		Address:       address,
	}, nil
}

func (c *StakingRewardContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *StakingRewardContract) GetABI() string {
	return StakingRewardABI
}

type FarmingRewardsContract struct {
	*FarmingRewards
	Address Addr
}

func NewFarmingRewardsContract(address Addr, client *ethclient.Client) (*FarmingRewardsContract, error) {
	farmingRewards, err := NewFarmingRewards(address, client)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsContract{
		FarmingRewards: farmingRewards,
		Address:        address,
	}, nil
}

func (c *FarmingRewardsContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *FarmingRewardsContract) GetABI() string {
	return FarmingRewardsABI
}

type GovernContract struct {
	*Govern
	Address Addr
}

func NewGovernContract(address Addr, client *ethclient.Client) (*GovernContract, error) {
	govern, err := NewGovern(address, client)
	if err != nil {
		return nil, err
	}
	return &GovernContract{
		Govern:  govern,
		Address: address,
	}, nil
}

func (c *GovernContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *GovernContract) GetABI() string {
	return GovernABI
}

type ViewerContract struct {
	*Viewer
	Address Addr
}

func NewViewerContract(address Addr, client *ethclient.Client) (*ViewerContract, error) {
	Viewer, err := NewViewer(address, client)
	if err != nil {
		return nil, err
	}
	return &ViewerContract{
		Viewer:  Viewer,
		Address: address,
	}, nil
}

func (c *ViewerContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *ViewerContract) GetABI() string {
	return ViewerABI
}

type BridgeContract struct {
	*Bridge
	Address Addr
}

func NewBridgeContract(address Addr, client *ethclient.Client) (*BridgeContract, error) {
	bridge, err := NewBridge(address, client)
	if err != nil {
		return nil, err
	}
	return &BridgeContract{
		Bridge:  bridge,
		Address: address,
	}, nil
}

func (c *BridgeContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *BridgeContract) GetABI() string {
	return BridgeABI
}

type PegBridgeContract struct {
	*PeggedTokenBridge
	Address Addr
}

func NewPegBridgeContract(address Addr, client *ethclient.Client) (*PegBridgeContract, error) {
	ptb, err := NewPeggedTokenBridge(address, client)
	if err != nil {
		return nil, err
	}
	return &PegBridgeContract{
		PeggedTokenBridge: ptb,
		Address:           address,
	}, nil
}

func (c *PegBridgeContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *PegBridgeContract) GetABI() string {
	return PeggedTokenBridgeABI
}

type PegBridgeV2Contract struct {
	*PeggedTokenBridgeV2
	Address Addr
}

func NewPegBridgeV2Contract(address Addr, client *ethclient.Client) (*PegBridgeV2Contract, error) {
	ptb, err := NewPeggedTokenBridgeV2(address, client)
	if err != nil {
		return nil, err
	}
	return &PegBridgeV2Contract{
		PeggedTokenBridgeV2: ptb,
		Address:             address,
	}, nil
}

func (c *PegBridgeV2Contract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *PegBridgeV2Contract) GetABI() string {
	return PeggedTokenBridgeV2ABI
}

type PegVaultContract struct {
	*OriginalTokenVault
	Address Addr
}

func NewPegVaultContract(address Addr, client *ethclient.Client) (*PegVaultContract, error) {
	otv, err := NewOriginalTokenVault(address, client)
	if err != nil {
		return nil, err
	}
	return &PegVaultContract{
		OriginalTokenVault: otv,
		Address:            address,
	}, nil
}

func (c *PegVaultContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *PegVaultContract) GetABI() string {
	return OriginalTokenVaultABI
}

type PegVaultV2Contract struct {
	*OriginalTokenVaultV2
	Address Addr
}

func NewPegVaultV2Contract(address Addr, client *ethclient.Client) (*PegVaultV2Contract, error) {
	otv, err := NewOriginalTokenVaultV2(address, client)
	if err != nil {
		return nil, err
	}
	return &PegVaultV2Contract{
		OriginalTokenVaultV2: otv,
		Address:              address,
	}, nil
}

func (c *PegVaultV2Contract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *PegVaultV2Contract) GetABI() string {
	return OriginalTokenVaultV2ABI
}

type WdInboxContract struct {
	*WithdrawInbox
	Address Addr
}

func NewWdInboxContract(address Addr, client *ethclient.Client) (*WdInboxContract, error) {
	wdi, err := NewWithdrawInbox(address, client)
	if err != nil {
		return nil, err
	}
	return &WdInboxContract{
		WithdrawInbox: wdi,
		Address:       address,
	}, nil
}

func (c *WdInboxContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *WdInboxContract) GetABI() string {
	return WithdrawInboxABI
}

type CLPContract struct {
	*ContractAsLP
	Address Addr
}

func NewContractAsLPContract(address Addr, client *ethclient.Client) (*CLPContract, error) {
	cLP, err := NewContractAsLP(address, client)
	if err != nil {
		return nil, err
	}
	return &CLPContract{
		ContractAsLP: cLP,
		Address:      address,
	}, nil
}

func (c *CLPContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *CLPContract) GetABI() string {
	return ContractAsLPABI
}

type MsgBusContract struct {
	*MessageBus
	Address Addr
}

func NewMsgBusContract(address Addr, client *ethclient.Client) (*MsgBusContract, error) {
	bus, err := NewMessageBus(address, client)
	if err != nil {
		return nil, err
	}
	return &MsgBusContract{
		MessageBus: bus,
		Address:    address,
	}, nil
}

func (c *MsgBusContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *MessageBus) GetABI() string {
	return MessageBusABI
}

type XferAgtContract struct {
	*TransferAgent
	Address Addr
}

func NewXferAgtContract(address Addr, client *ethclient.Client) (*XferAgtContract, error) {
	agent, err := NewTransferAgent(address, client)
	if err != nil {
		return nil, err
	}
	return &XferAgtContract{
		TransferAgent: agent,
		Address:       address,
	}, nil
}

func (c *XferAgtContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *XferAgtContract) GetABI() string {
	return TransferAgentABI
}

type VotingEscrowTokenContract struct {
	*VotingEscrowToken
	Address Addr
}

func NewVotingEscrowTokenContract(address Addr, client *ethclient.Client) (*VotingEscrowTokenContract, error) {
	votingEscrowToken, err := NewVotingEscrowToken(address, client)
	if err != nil {
		return nil, err
	}
	return &VotingEscrowTokenContract{
		VotingEscrowToken: votingEscrowToken,
		Address:           address,
	}, nil
}

func (c *VotingEscrowTokenContract) GetAddr() Addr {
	if c == nil {
		return ZeroAddr
	}
	return c.Address
}

func (c *VotingEscrowTokenContract) GetABI() string {
	return VotingEscrowTokenABI
}
