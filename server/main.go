package main

import (
	"flag"

	"github.com/scriptonist/Carehacks/server/daemon"
)

func main() {

	daemonConfig := daemon.Config{}

	flag.StringVar(&daemonConfig.Listenspec, "listen", "0.0.0.0:9080", "HTTP listen spec")
	flag.Parse()
	daemon.StartServer(daemonConfig)

}
