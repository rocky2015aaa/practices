package dbcontrol

import (
	. "n1qlServerTest/util"

	"github.com/couchbase/gocb"
)

const (
	CouchbaseUrl = "couchbase://localhost"
)

var cluster *gocb.Cluster

func init() {
	_ = connectDB()
}

func connectDB() error {
	if cluster == nil {
		cbCluster, err := gocb.Connect(CouchbaseUrl)
		if err != nil {
			LoggerError.Println("Failed to connect DB")
			return err
		}
		cluster = cbCluster
	}

	LoggerInfo.Println("Successfully connected couchbase db")
	return nil
}

func GetCluster() *gocb.Cluster {
	return cluster
}
