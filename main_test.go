package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/alshashiguchi/api-go-gin/controllers"
	"github.com/alshashiguchi/api-go-gin/database"
	"github.com/alshashiguchi/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/andre", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	mockDaResposta := `{"API diz:":"E ai andre, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	// Arrange
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()

	// act
	r.ServeHTTP(res, req)

	//Assert
	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}

func TestLBuscaAlunoPorCPFHandler(t *testing.T) {
	// Arrange
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/aluno/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/aluno/cpf/12345678901", nil)
	res := httptest.NewRecorder()

	// act
	r.ServeHTTP(res, req)

	//Assert
	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/aluno/:id", controllers.BuscaAlunoPorId)
	pathDaBusca := "/aluno/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	res := httptest.NewRecorder()

	// act
	r.ServeHTTP(res, req)

	//Assert
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do aluno Teste", alunoMock.Nome, "Deveriam ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF, "Deveriam ser iguais")
	assert.Equal(t, "123456789", alunoMock.RG, "Deveriam ser iguais")
}

func TestDeleteAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/aluno/:id", controllers.DeletaAluno)
	pathDaBusca := "/aluno/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	res := httptest.NewRecorder()

	// act
	r.ServeHTTP(res, req)

	//Assert
	assert.Equal(t, http.StatusOK, res.Code, "Deveriam ser iguais")
}
