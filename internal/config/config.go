package config

import (
	"time"
	mysql "wikivin/pkg/database/MySQL"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	HTTP HTTPConfig
	MySQL mysql.MySQLConfig
}

type HTTPConfig struct {
	Addr        string `mapstructure:"port"`
	ReadTimeout time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderBytes int `mapstructure:"maxHeaderBytes"`
}

func Init(path string) (*Config, error){
	if err:= parseConfigFile(path); err != nil{
		return nil, err
	}
	var cfg Config
	if err:= unmarshal(&cfg); err != nil{
		return nil, err
	}
	if err:= setFromEnv(&cfg); err != nil{
		return nil, err
	}
	return &cfg, nil
}

func unmarshal(config *Config) error{
	if err:= viper.UnmarshalKey("http", &config.HTTP); err != nil{
		return err
	}
	return nil
}

func setFromEnv(config *Config) error{
	if err:= gotenv.Load("../../.env"); err != nil{
		return err
	}
	
	if err:= envconfig.Process("DB", &config.MySQL); err != nil{
		return err
	}

	return nil
}

func parseConfigFile(path string) error{
	viper.AddConfigPath(path)
	viper.SetConfigName("server")
	viper.SetConfigType("yaml")
	if err:= viper.ReadInConfig(); err!= nil{
		if _, ok:= err.(viper.ConfigFileNotFoundError); ok{
			return err;
		}else{
			return err
		}
	}
	return viper.MergeInConfig()
}