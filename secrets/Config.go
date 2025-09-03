package secrets

import (
	"github.com/spf13/viper"
)

type MongoDB struct {
	MongoDBURI        string `mapstructure:"MONGODB_URI"`
	DbName            string `mapstructure:"DB"`
	UserCollection    string `mapstructure:"USER_COLLECTION"`
	ProductCollection string `mapstructure:"PRODUCT_COLLECTION"`
	CartCollection    string `mapstructure:"CART_COLLECTION"`
	RedisAddr         string
	RedisPassword     string
	RedisDB           int
	RedisChannel      string
}

var Config *MongoDB

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("MONGODB_URI", "mongodb://localhost:27017/")
	viper.SetDefault("DB", "fiber_mart")
	viper.SetDefault("USER_COLLECTION", "users")
	viper.SetDefault("PRODUCT_COLLECTION", "products")
	viper.SetDefault("CART_COLLECTION", "cart")
	viper.SetDefault("Redis.Addr", "localhost:6379")
	viper.SetDefault("Redis.Password", "")
	viper.SetDefault("Redis.DB", 0)
	viper.SetDefault("Redis.Channel", "fiber_audit")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	Config = &MongoDB{}
	if err := viper.Unmarshal(Config); err != nil {
		panic(err)
	}
}
