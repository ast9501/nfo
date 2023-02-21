package service

import (
	"github.com/ast9501/nfo/internal/service/management"
	"github.com/ast9501/nfo/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli/v2"
)

// Define config file fields
type Config struct {
	DMAAPServer     string `mapstructure:"DMAAP_SERVER"`
	DMAAPPort       string `mapstructure:"DMAAP_PORT"`
	DMAAPWriteTopic string `mapstructure:"DMAAP_WRITE_TOPIC"`
	DMAAPReadTopic  string `mapstructure:"DMAAP_READ_TOPIC"`
	Cert            string `mapstructure:"TLS_CERT"`
	Key             string `mapstructure:"TLS_KEY"`
	Addr            string `mapstructure:"NFO_BIND_ADDR"`
	Port            string `mapstructure:"NFO_BIND_PORT"`
}

type NFO struct {
	Config Config
}

var cliCmd = []cli.Flag{
	&cli.StringFlag{
		Name:  "c",
		Usage: "Path for config file",
	},
}

func (*NFO) GetCliCmd() (flags []cli.Flag) {
	return cliCmd
}

func (nfo *NFO) Initialize(configPath string) {
	nfo.LoadConfig(configPath)

	// derive cert path
	nfo.Config.Cert = configPath + "/" + nfo.Config.Cert
	nfo.Config.Key = configPath + "/" + nfo.Config.Key

	// derive port binding
	nfo.Config.Port = ":" + nfo.Config.Port
}

// schemes http
func (nfo *NFO) Start(certPath string, keyPath string) {
	// TODO: external function call for customize gin engine loger
	router := gin.Default()

	// Add service to router
	management.AddService(router)

	// api server by swagger in debug mode
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Load server binding port from conifg
	router.RunTLS(nfo.Config.Port, certPath, keyPath)
}

func (n *NFO) LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("nfo")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	n.Config = config
	logger.InitLog.Infoln("Load config successful!")

	return
}
