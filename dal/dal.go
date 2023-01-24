package dal

import (
	_ "embed"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/im-executor/types"
	"github.com/spf13/viper"
)

type DB struct {
	*sqldb.Db
}

//go:embed schema.sql
var schema string

var db *DB

func InitDB() *DB {
	log.Infoln("Creating DB connection")
	url := viper.GetString(types.KeyDbUrl)
	log.Infoln(url)
	_db, err := sqldb.NewDb("postgres", fmt.Sprintf("postgresql://root@%s/executor?sslmode=disable", url), 4)
	if err != nil {
		log.Fatalf("Failed to create db with url %s: %+v", url, err)
	}
	log.Infoln("Syncing DB schemas")
	_, err = _db.Exec(schema)
	if err != nil {
		log.Fatalln("failed to initialize tables", err)
	}
	db = &DB{_db}
	return db
}

func GetDB() *DB {
	if db == nil {
		log.Fatalf("DB uninitialized")
	}
	return db
}
