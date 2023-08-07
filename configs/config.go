package configs

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Conf = new(Config)

const DefaultConfigPath = "./config.yaml"

type Config struct {
	DBConfig `mapstructure:"db"`
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func init() {
	// 默认配置文件路径
	var configPath string
	flag.StringVar(&configPath, "config", "", "配置文件绝对路径或相对路径")
	flag.Parse()
	if configPath == "" {
		configPath = DefaultConfigPath
	}
	logrus.Printf("===> config path: %s", configPath)
	// 初始化配置文件
	viper.SetConfigFile(configPath)
	viper.WatchConfig()
	// 观察配置文件变动
	viper.OnConfigChange(func(in fsnotify.Event) {
		logrus.Printf("config file has changed")
		if err := viper.Unmarshal(&Conf); err != nil {
			logrus.Errorf("failed at unmarshal config file after change, err: %v", err)
			panic(fmt.Sprintf("failed at init config: %v", err))
		}
	})
	// 将配置文件读入 viper
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("failed at ReadInConfig, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}
	// 解析到变量中
	if err := viper.Unmarshal(&Conf); err != nil {
		logrus.Errorf("failed at Unmarshal config file, err: %v", err)
		panic(fmt.Sprintf("failed at init config: %v", err))
	}
	cephalonDBHost := os.Getenv("CEPHALON_DB_HOST")
	if cephalonDBHost != "" {
		Conf.DBConfig.Host = cephalonDBHost
	}
	cephalonDBPassword := os.Getenv("CEPHALON_DB_PASSWORD")
	if cephalonDBPassword != "" {
		Conf.DBConfig.Password = cephalonDBPassword
	}
	// 返回 nil 或错误
	logrus.Infoln("global config init success...")
}
