package main

import "github.com/scriptonist/Carehacks/server/daemon"

func main() {
	daemonConfig := daemon.Config{
		Listenspec: "localhost:9080",
	}

	daemon.StartServer(daemonConfig)

}
