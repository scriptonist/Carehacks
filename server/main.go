package main

import "github.com/scriptonist/gitnots/server/daemon"

func main() {
	daemonConfig := daemon.Config{
		Listenspec: "localhost:9080",
	}

	daemon.StartServer(daemonConfig)

}
