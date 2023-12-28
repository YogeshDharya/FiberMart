package secrets 
import( 
"github.com/spf13/viper"
"io"
)
  type MongoDB struct {
  MongoDBURI string 
  DbName string 
  UserCollection string 
  ProductCollection string 
  CartCollection string
}  
// type AppConfig struct{ What's the   purpose exactly  
//   mongo   MongoDB   
// }  
var Config *MongoDB  
func init(){  
  viper.SetConfigFile(".env")  
  viper.AutomaticEnv()  
  viper.SetDefault("MongoDB.URI","mongodb+srv://Dante:Dante$1998@myfirstcluster.zg5z4w8.mongodb.net/?retryWrites=true&w=majority")
viper.SetDefault("MongoDB.DefaultDB","MyKart")
  viper.SetDefault("MongoDB.UserCollection","Users") 
  viper.SetDefault("MongoDB.ProductCollection","Products")
  viper.SetDefault("MongoDB.CartCollection","Cart")
  if err := viper.ReadConfig(); err!= nil {
    panic(err)
  }
  Config = &MongoDB{}
  if err:= viper.Unmarshal(Config); err != nil{
    panic(err)
  }
}