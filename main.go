package main

import (
	"maneeshaindrachapa.github.io/go-server-gin/configs"
	"maneeshaindrachapa.github.io/go-server-gin/routes"
)

func main() {
	// load env variables just once in here so can be use in any other place
	configs.InitEnvConfigs()
	routes.InitRoutes()

	routes.Router.Run(configs.EnvConfigs.LocalServerPort)
}
