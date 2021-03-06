package coreCoinsys

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	MClient *mongo.Client
}

type CryptoJSON struct {
	Response   string
	Type       int
	Aggregated bool
	Data       []CoinObject
	Timeto     int64
	Timefrom   int64
}

type CoinObject struct {
	Time       int64
	Close      float64
	High       float64
	Low        float64
	Open       float64
	Volumefrom float64
	Volumeto   float64
}

func Fetch(length string) {
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

	fetchCryptoAPI(c.SetupConfig.CryptoKey, c.CApi.PriceSingleSymbolSrice, length, c.SetupConfig.MongoDB)
}

func fetchCryptoAPI(apikey string, endpoint string, length string, port string) {
	switch length {
	case "all":
		resp, err := http.Get(endpoint + "&allData=true&api_key={" + apikey + "}")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var cryptojson CryptoJSON
		json.Unmarshal([]byte(body), &cryptojson)
		storeCryptoAPI(cryptojson, "all_days", port)
	default:
		url := fmt.Sprintf("&limit=%s&api_key={", length)
		resp, err := http.Get(endpoint + url + apikey + "}")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var cryptojson CryptoJSON
		json.Unmarshal([]byte(body), &cryptojson)
		storeCryptoAPI(cryptojson, fmt.Sprintf("%s_days", length), port)
	}
}

func storeCryptoAPI(cryptojson CryptoJSON, length string, port string) {
	mc := startMongodbClient(port)
	mc.clearMongodb("test", "BTC_Closing_Value_"+length)
	mc.clearMongodb("test", "BTC_Closing_Timestamp_"+length)
	for _, element := range cryptojson.Data {
		mc.insertMongodbFloat("test", "BTC_Closing_Value_"+length, element.Close)
		mc.insertMongodbInt("test", "BTC_Closing_Timestamp_"+length, element.Time)
	}
}

func startMongodbClient(port string) *MongoClient {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:" + port)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	cl := MongoClient{
		MClient: client,
	}
	return &cl
}

func (mc *MongoClient) insertMongodbFloat(dbName string, collection string, elementVal float64) {
	conn := mc.MClient.Database(dbName).Collection(collection)
	element := CoinDescFloat{elementVal}
	serializedElement := []interface{}{element}

	// insertion, err := conn.InsertMany(context.TODO(), serializedElement)
	_, err := conn.InsertMany(context.TODO(), serializedElement)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Inserted multiple documents: ", insertion.InsertedIDs)
}

func (mc *MongoClient) insertMongodbInt(dbName string, collection string, elementVal int64) {
	conn := mc.MClient.Database(dbName).Collection(collection)
	element := CoinDescInt{elementVal}
	serializedElement := []interface{}{element}

	// insertion, err := conn.InsertMany(context.TODO(), serializedElement)
	_, err := conn.InsertMany(context.TODO(), serializedElement)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Inserted multiple documents: ", insertion.InsertedIDs)
}

func (mc *MongoClient) clearMongodb(dbName string, collection string) {
	mc.MClient.Database(dbName).Collection(collection).Drop(context.TODO())
}
