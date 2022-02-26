package main

import (
	"github.com/alshashiguchi/api-go-gin/database"
	"github.com/alshashiguchi/api-go-gin/models"
	"github.com/alshashiguchi/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	models.Alunos = []models.Aluno{
		{Nome: "André Luiz", CPF: "999999999-99", RG: "00.000.000-0"},
		{Nome: "Arthur Yoshiharu", CPF: "000000000-00", RG: "11.000.000-0"},
	}
	routes.HandleRequest()
}
