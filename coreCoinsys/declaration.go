package coreCoinsys

type Setup struct {
	CryptoKey string `mapstructure:"api_key"`
	MongoDB   string `mapstructure:"mongodb_port"`
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
	SetupConfig Setup          `mapstructure:"config"`
	CApi        CryptoAPI      `mapstructure:"cryptoendpoints"`
	Db          DatabaseConfig `mapstructure:"database"`
}

type Trainer struct {
	Name string
	Age  int
	City string
}

type CoinDesc struct {
	Name  string
	Value float64
}
