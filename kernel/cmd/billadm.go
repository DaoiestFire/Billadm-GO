package cmd

import "github.com/billadm/kernel/server"

func main() {
	ginServer := server.NewGinServer()
	ginServer.Run()
}
