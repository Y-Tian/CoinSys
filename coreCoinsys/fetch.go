package coreCoinsys

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

var apiKey string = "3c3efe2e831e7153df7f980cbc6149c8b61023d5d93c590763cfd7a68e6bc294"

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type CryptoAPI struct {
	PriceSingleSymbolSrice string `mapstructure:"price_single_symbol_price"`
}

type OutputConfig struct {
	File string
}

type Config struct {
	Db   DatabaseConfig `mapstructure:"database"`
	CApi CryptoAPI      `mapstructure:"cryptoendpoints"`
	Out  OutputConfig   `mapstructure:"output"`
}

func Init() {
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

	fetch(c.CApi.PriceSingleSymbolSrice)
}

func fetch(endpoint string) {
	resp, err := http.Get(endpoint + "&api_key={" + apiKey + "}")
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
