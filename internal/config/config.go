package config

import (
	"os"
	"time"
	mysql "wikivin/pkg/database/MySQL"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	HTTP HTTPConfig
	MySQL mysql.MySQLConfig
	Auth AuthConfig
}

type AuthConfig struct{
	JWT  JWTConfig
	PasswordSalt string
}

type JWTConfig struct{
	AccessTokenTTL time.Duration `mapstructure:"accessTokenTTL"`
	RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	SecretAccessKey string
	SecretRefreshKey string
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
	if err:= viper.UnmarshalKey("auth", &config.Auth.JWT); err != nil{
		return err
	}
	return nil
}

func setFromEnv(config *Config) error{
	if err:= gotenv.Load("/root/.env"); err != nil{
		return err
	}
	
	if err:= envconfig.Process("DB", &config.MySQL); err != nil{
		return err
	}

	config.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	config.Auth.JWT.SecretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	config.Auth.JWT.SecretRefreshKey = os.Getenv("SECRET_REFRESH_KEY")
	
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