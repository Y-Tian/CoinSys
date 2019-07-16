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

type CoinDesc struct {
	Value float64
}

type FindCoinDesc struct {
	// ID    primitive.ObjectID `bson:"_id"`
	Value float64 `bson:"value"`
}

type GraphingAxis struct {
	XAxis []int64
	YAxis []float64
}
