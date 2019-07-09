package coreCoinsys

import (
	"context"
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

func TestInit() {
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

	fetchCryptoAPI(c.SetupConfig.CryptoKey, c.CApi.PriceSingleSymbolSrice)
	// log.Println(c.SetupConfig.CryptoKey)
	// startMongodbClient(c.SetupConfig.MongoDB)
	testMongodbClient(c.SetupConfig.MongoDB)
}

func fetchCryptoAPI(apikey string, endpoint string) {
	resp, err := http.Get(endpoint + "&api_key={" + apikey + "}")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func testMongodbClient(port string) {
	mc := startMongodbClient(port)
	mc.insertMongodb()
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

	fmt.Println("Connected to MongoDB!")

	cl := MongoClient{
		MClient: client,
	}
	return &cl
}

func (mc *MongoClient) insertMongodb() {
	collection := mc.MClient.Database("test").Collection("Bitcoin")
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	bitcn := CoinDesc{"BTC", 1600}

	// trainers := []interface{}{misty, brock}
	coins := []interface{}{bitcn}

	insertManyResult, err := collection.InsertMany(context.TODO(), coins)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}
