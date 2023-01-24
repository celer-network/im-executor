package types

import "fmt"

func (evl *EventLog) String() string {
	return fmt.Sprintf("EventLog: address %s, txhash %x, block number %d, log index %d, event %x, spare %x.",
		evl.Address, evl.TxHash, evl.BlockNumber, evl.Index, evl.Event, evl.Spare)
}
