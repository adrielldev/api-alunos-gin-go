package main

import (
	"github.com/adrielldev/api-alunos-gin-go/database"
	"github.com/adrielldev/api-alunos-gin-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}
