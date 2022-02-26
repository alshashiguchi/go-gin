package routes

import (
	"github.com/alshashiguchi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)

	r.Run() // roda na porta 8080 para definir outra r.Run(":5000")
}
