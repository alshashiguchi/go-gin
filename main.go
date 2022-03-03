package main

import (
	"github.com/alshashiguchi/api-go-gin/database"
	"github.com/alshashiguchi/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
