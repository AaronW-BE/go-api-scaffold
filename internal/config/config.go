package config

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
	Redis struct {
		Addr     string
		Password string
	}
}

var Conf Config

func bindFlag(prefix string, m map[string]interface{}) {
	for k, val := range m {
		key := fmt.Sprintf("%s.%s", prefix, k)
		switch v := val.(type) {
		case map[string]interface{}:
			bindFlag(key, v)
		default:
			// 如果命令行提供参数，则覆盖
			if pflag.Lookup(strings.ReplaceAll(key, ".", "_")) != nil {
				pflag.Lookup(strings.ReplaceAll(key, ".", "_")).Value.Set(fmt.Sprintf("%v", val))
			}
		}
	}
}

func LoadConfig() {
	configFile := pflag.String("config", "config/config.yaml", "config file path")
	pflag.Parse()

	v := viper.New()
	v.SetConfigFile(*configFile)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Use config file:", v.ConfigFileUsed())
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(err)
	}

	v.WatchConfig()
	fmt.Printf("Current configuration: %+v\n", Conf)
}
