package model

import (
	. "n1qlServerTest/util"

	"github.com/couchbase/gocb"
)

func ExecuteQuery(query string, bucket *gocb.Bucket) (gocb.ViewResults, error) {
	result, err := bucket.ExecuteN1qlQuery(gocb.NewN1qlQuery(query), nil)
	if err != nil {
		LoggerError.Println("N1QL query error: ", err)
		return nil, err
	}

	return result, nil
}
