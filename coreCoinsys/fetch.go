package coreCoinsys

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type CryptoKey struct {
	CryptoKey string `mapstructure:"api_key"`
}

type CryptoAPI struct {
	PriceSingleSymbolSrice string `mapstructure:"price_single_symbol_price"`
}

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type Config struct {
	CKey CryptoKey      `mapstructure:"cryptokey"`
	CApi CryptoAPI      `mapstructure:"cryptoendpoints"`
	Db   DatabaseConfig `mapstructure:"database"`
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

	fetch(c.CKey.CryptoKey, c.CApi.PriceSingleSymbolSrice)
}

func fetch(apikey string, endpoint string) {
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
