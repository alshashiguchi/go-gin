package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de propósito, não se preocupe")
}
