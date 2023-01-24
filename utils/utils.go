package utils

import (
	"fmt"

	"github.com/celer-network/goutils/log"
)

func CheckErrf(err error, msg string, args ...interface{}) {
	if err != nil {
		msg = fmt.Sprintf(msg, args...)
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("fatal error occurred: %s", err.Error())
	}
}
