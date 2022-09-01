package routes

import (
	"github.com/alshashiguchi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/aluno/:id", controllers.BuscaAlunoPorId)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/aluno/:id", controllers.DeletaAluno)
	r.PATCH("/aluno/:id", controllers.EditaAluno)
	r.GET("/aluno/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.GET("/index", controllers.ExibePaginaIndex)

	r.Run() // roda na porta 8080 para definir outra r.Run(":5000")
}
