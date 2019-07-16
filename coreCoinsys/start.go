package coreCoinsys

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type Dataset struct {
	ClosingPrices []int
}

func Start() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	viper.ReadInConfig()

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	var closingValuesFromDataset []float64
	// loadTemp := loadFromMongoClient("test", "BTC_Closing_Value_30_days", c.SetupConfig.MongoDB)
	loadTempClosingValues := loadFromMongoClient("test", "BTC_Closing_Value_All_Time", c.SetupConfig.MongoDB)
	for _, element := range loadTempClosingValues {
		closingValuesFromDataset = append(closingValuesFromDataset, element.Value)
	}

	var closingTimestampFromDataset []int64
	loadTempClosingTimestamp := loadFromMongoClient("test", "BTC_Closing_Timestamp_All_Time", c.SetupConfig.MongoDB)
	for _, element := range loadTempClosingTimestamp {
		closingTimestampFromDataset = append(closingTimestampFromDataset, element.Value)
	}

	var MACDSlice []float64
	MACDSlice := FindMACD(closingValuesFromDataset)
	var timestampSlice []int64
	timestampSlice := prepTimeAxis(closingTimestampFromDataset, len(MACDSlice))
}

func loadFromMongoClient(dbName string, collection string, port string) []FindCoinDesc {
	mc := startMongodbClient(port)

	conn := mc.MClient.Database(dbName).Collection(collection)

	// Here's an array in which you can store the decoded documents
	var results []FindCoinDesc

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := conn.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem FindCoinDesc
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func reverseArrayOrderFloat(input []float64) []float64 {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}

func reverseArrayOrderInt(input []int64) []int64 {
	for i := len(input)/2 - 1; i >= 0; i-- {
		opp := len(input) - 1 - i
		input[i], input[opp] = input[opp], input[i]
	}
	return input
}

func prepTimeAxis(input []int64, period int) []int64 {
	var xAxis []int64
	var temp []int64
	temp := reverseArrayOrderInt(input)
	for i := 0; i < period; i++ {
		xAxis = append(xAxis, temp[i])
	}
	return reverseArrayOrderInt(xAxis)
}
