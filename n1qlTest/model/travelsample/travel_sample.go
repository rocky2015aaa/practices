package travelsample

import (
	"n1qlServerTest/dbcontrol"
	"n1qlServerTest/model"
	. "n1qlServerTest/util"

	"github.com/couchbase/gocb"
)

const (
	TravelSampleBucketName = "travel-sample"
)

var cluster *gocb.Cluster
var travelSampleBucket *gocb.Bucket

type Airlines struct {
	Airlines []*Airline `json:"airlines"`
}

type Airline struct {
	Id       int    `json:"id"`
	Type     string `json:"string"`
	Name     string `json:"string"`
	Iata     string `json:"iata"`
	Icao     string `json:"icao"`
	Callsign string `json:"callsign"`
	Country  string `json:"country"`
}

func init() {
	cluster = dbcontrol.GetCluster()
	_ = openTravelSampleBucket()
}

func openTravelSampleBucket() error {
	travelSampleB, err := cluster.OpenBucket(TravelSampleBucketName, "")
	if err != nil {
		LoggerError.Println("Failed to get ", TravelSampleBucketName, "bucket")
		return err
	}
	travelSampleBucket = travelSampleB
	return err
}

func GetTravelSampleBucket() *gocb.Bucket {
	return travelSampleBucket
}

func GetAllAirlines() (airlines Airlines, err error) {
	query := "select `" + TravelSampleBucketName + "`.* from `" + TravelSampleBucketName + "` where type=\"airline\""
	queryResult, err := model.ExecuteQuery(query, travelSampleBucket)
	if err != nil {
		LoggerError.Println("Failed to get query result")
		return airlines, err
	}
	return getResult(queryResult)
}

func getResult(queryResult gocb.ViewResults) (airlines Airlines, err error) {
	results := []*Airline{}
	var result Airline
	for queryResult.Next(&result) {
		results = append(results, &result)
	}
	if err = queryResult.Close(); err != nil {
		LoggerError.Println("N1ql query error: ", err)
		return airlines, err
	}
	airlines.Airlines = results
	return airlines, nil
}
