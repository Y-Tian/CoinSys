package coreCoinsys

import (
	"context"
	_ "expvar"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

type Dataset struct {
	ClosingPrices []int
}

func Start(length string) {
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
	loadTempClosingValues := loadFromMongoClientFloat("test", fmt.Sprintf("BTC_Closing_Value_%s_days", length), c.SetupConfig.MongoDB)
	for _, element := range loadTempClosingValues {
		closingValuesFromDataset = append(closingValuesFromDataset, element.Value)
	}

	var closingTimestampFromDataset []int64
	loadTempClosingTimestamp := loadFromMongoClientInt("test", fmt.Sprintf("BTC_Closing_Timestamp_%s_days", length), c.SetupConfig.MongoDB)
	for _, element := range loadTempClosingTimestamp {
		closingTimestampFromDataset = append(closingTimestampFromDataset, element.Value)
	}

	var TimestampSlice []int64
	var MACDSlice []float64
	var SignalSlice []float64
	var HistogramSlice []float64
	MACDSlice = FindMACD(closingValuesFromDataset)
	TimestampSlice = prepTimeAxis(closingTimestampFromDataset, len(MACDSlice))
	SignalSlice = FindSignalLine(MACDSlice)
	HistogramSlice = FindHistogram(MACDSlice, SignalSlice)
	sliceOffset := len(MACDSlice) - len(HistogramSlice)
	MACDSlice = MACDSlice[sliceOffset:]
	TimestampSlice = TimestampSlice[sliceOffset:]

	fmt.Println("Please visit localhost:8080 to see a visualization.")
	fmt.Println("Reminder: your visualization will be delayed 30 days due to the nature of the MACD indicator." +
		"Please fetch your data range accordingly.")
	RunGraph(TimestampSlice, MACDSlice, SignalSlice, HistogramSlice)
}

func loadFromMongoClientFloat(dbName string, collection string, port string) []FindCoinDescFloat {
	mc := startMongodbClient(port)
	conn := mc.MClient.Database(dbName).Collection(collection)
	var results []FindCoinDescFloat

	cur, err := conn.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem FindCoinDescFloat
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results
}

func loadFromMongoClientInt(dbName string, collection string, port string) []FindCoinDescInt {
	mc := startMongodbClient(port)
	conn := mc.MClient.Database(dbName).Collection(collection)
	var results []FindCoinDescInt

	cur, err := conn.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem FindCoinDescInt
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
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
	temp = reverseArrayOrderInt(input)
	for i := 0; i < period; i++ {
		xAxis = append(xAxis, temp[i])
	}
	return reverseArrayOrderInt(xAxis)
}
