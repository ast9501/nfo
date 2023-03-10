package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ast9501/nfo/docs"
	"github.com/ast9501/nfo/pkg/logger"
	nfo_service "github.com/ast9501/nfo/pkg/service"
	"github.com/urfave/cli/v2"
)

var NFO = &nfo_service.NFO{}

//	@title			O-RAN NFO api doc
//	@version		1.0
//	@description	winlab O-RAN NFO

//	@contact.name	ast9501
//	@contact.email	ast9501.cs10@nycu.edu.tw

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// schemes http
func main() {
	app := cli.NewApp()
	app.Name = "nfo"
	app.Usage = "3GPP NFO function for O-RAN"
	app.Action = action
	app.Flags = NFO.GetCliCmd()

	if err := app.Run(os.Args); err != nil {
		// TODO: Add logger printer
		logger.InitLog.Errorf("NFO Run Error: %v\n", err)
	}
	// generate host ip dynamicly for api doc
	docs.SwaggerInfo.Host = NFO.Config.Addr + NFO.Config.Port
	logger.InitLog.Debugln("Generate swagger api doc target server location: ", docs.SwaggerInfo.Host)

}

func action(c *cli.Context) error {
	if c.String("c") == "" {
		fmt.Println("config is null!")
		return nil
	}
	// TODO: Add log: print config file path
	NFO.Initialize(c.String("c"))

	NFO.Start(NFO.Config.Cert, NFO.Config.Key)

	return nil
}

// generate server outbound IP for api doc
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
