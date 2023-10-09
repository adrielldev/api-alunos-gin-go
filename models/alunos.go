package models

type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"CPF"`
	RG   string `json:"RG"`
}

var Alunos []Aluno
