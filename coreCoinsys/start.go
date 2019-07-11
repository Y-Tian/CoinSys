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

	var valuesFromDataset []float64
	loadTemp := loadFromMongoClient("test", "BTC_Closing_Value_30_days", c.SetupConfig.MongoDB)
	for _, element := range loadTemp {
		valuesFromDataset = append(valuesFromDataset, element.Value)
	}

	findSMA(valuesFromDataset)
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

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}

func findSMA(values []float64) {
	log.Println(values)
}
