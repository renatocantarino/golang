package configs

//https://github.com/spf13/viper
//https://github.com/go-chi/jwtauth
import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        string `mapstructure:"DB_PORT"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPass        string `mapstructure:"DB_PASS"`
	DbName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEBSERVER_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	JwtTokenAuth  *jwtauth.JWTAuth
}

var cfg *config

func Load(pathConfiguration string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(pathConfiguration)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.JwtTokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return cfg, err

}
