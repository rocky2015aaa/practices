package controller

import (
	"github.com/couchbase/gocb"
)

var bucket *gocb.Bucket

func init() {
	cluster, _ := gocb.Connect("couchbase://127.0.0.1")
	bucket, _ = cluster.OpenBucket("welcomeAboard", "")
}

func GetBucket() *gocb.Bucket {
	return bucket
}
