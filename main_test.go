package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alshashiguchi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
