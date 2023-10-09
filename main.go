package main

import (
	"github.com/adrielldev/api-alunos-gin-go/models"
	"github.com/adrielldev/api-alunos-gin-go/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Adriel", CPF: "123", RG: "123"},
	}
	routes.HandleRequests()
}
