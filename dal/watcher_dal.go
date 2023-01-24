package dal

import (
	"github.com/celer-network/goutils/eth/mon2"
	//"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/sqldb"
)

// implements WatchDAL interface methods
var _ mon2.DAL = &DB{}

func (dal *DB) GetMonitorBlock(key string) (blknum uint64, blkidx int64, found bool, err error) {
	q := `SELECT block_num, block_idx 
				FROM monitor_block 
				WHERE chid_addr = $1`
	err = dal.QueryRow(q, key).Scan(&blknum, &blkidx)
	found, err = sqldb.ChkQueryRow(err)
	return
}

func (dal *DB) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	q := `UPSERT INTO monitor_block (event, block_num, block_idx) 
				VALUES ($1, $2, $3)`
	_, err := dal.Exec(q, key, blockNum, blockIdx)
	return err
}
