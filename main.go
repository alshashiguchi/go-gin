package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Run() // roda na porta 8080 para definir outra r.Run(":5000")
}
