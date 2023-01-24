package types

import (
	fmt "fmt"
	"log"
	"strings"

	"github.com/celer-network/im-executor/sgn-v2/common"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/near/borsh-go"
	"golang.org/x/crypto/sha3"
	"gopkg.in/yaml.v2"
)

func (m *Message) PrettyLog() string {
	var sigstr []string
	for _, sig := range m.Signatures {
		sigstr = append(sigstr, fmt.Sprintf("%s:%x", sig.Signer, sig.SigBytes))
	}
	return fmt.Sprintf("message: chain %d->%d addr %s->%s type %s status %s lastReqTime %s data %x sigs %s",
		m.SrcChainId, m.DstChainId, m.Sender, m.Receiver, m.TransferType, m.ExecutionStatus,
		common.TsSecToTime(uint64(m.LastReqTime)), m.Data, strings.Join(sigstr, ", "))
}

func (m *Message) PrettyPrint() {
	if m == nil {
		fmt.Println("nil message")
	}
	fmt.Printf("src_chain_id: %d\n", m.SrcChainId)
	fmt.Printf("dst_chain_id: %d\n", m.DstChainId)
	fmt.Printf("sender: %s\n", m.Sender)
	fmt.Printf("receiver: %s\n", m.Receiver)
	fmt.Printf("fee: %s\n", m.Fee)
	fmt.Printf("transfer_type: %s\n", m.TransferType.String())
	fmt.Printf("transfer_ref_id: %x\n", m.TransferRefId)
	fmt.Printf("src_tx_hash: %s\n", m.SrcTxHash)
	fmt.Printf("execution_status: %s\n", m.ExecutionStatus.String())
	fmt.Printf("last_req_time: %s\n", common.TsSecToTime(uint64(m.LastReqTime)))
	var signers []string
	for _, sig := range m.Signatures {
		signers = append(signers, sig.Signer)
	}
	fmt.Println("signers:", signers)
	fmt.Printf("data: 0x%x\n", m.Data)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

type InputDataRoute struct {
	sender     []byte
	receiver   []byte
	srcChainId uint64
	srcTxId    []byte
	dstChainId uint64
	msgType    []byte
}

type InputData struct {
	msgData []byte
	route   InputDataRoute
}

func (m *Message) ComputeProxyChainMessageIdNoTransfer() []byte {
	input := InputData{
		msgData: m.Data,
		route: InputDataRoute{
			sender:     eth.Hex2Bytes(m.Sender),
			receiver:   eth.Hex2Bytes(m.Receiver),
			srcChainId: m.SrcChainId,
			srcTxId:    []byte(m.SrcTxHash),
			dstChainId: m.DstChainId,
			msgType:    []byte("MessageOnly"), // only support message only now
		},
	}
	res, err := borsh.Serialize(input)
	if err != nil {
		log.Fatalln("borsh serialize failed")
	}
	return crypto.Keccak256(res)
}

func (m *Message) ComputeMessageIdNoTransfer() []byte {
	var data []byte
	data = append(data, eth.ToPadBytes(uint8(MsgType_MSG_TYPE_MESSAGE))...)
	if m.IsBytesAddr {
		data = append(data, eth.Hex2Bytes(m.Sender)...)
	} else {
		data = append(data, eth.Hex2Addr(m.Sender).Bytes()...)
	}
	data = append(data, eth.Hex2Addr(m.Receiver).Bytes()...)
	data = append(data, eth.ToPadBytes(m.SrcChainId)...)
	data = append(data, eth.Hex2Hash(m.SrcTxHash).Bytes()...)
	data = append(data, eth.ToPadBytes(m.DstChainId)...)
	data = append(data, m.Data...)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}
