package main

import (
    "fmt"
)

type Endereco struct {
	Logradouro string
	Numero int
	Bairro string
}

// interfaces são implementadas automaticamente

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco
}

type Pessoa interface {
	Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func desativacao (pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente := Cliente{
		Nome: "John Doe",
        Idade: 30,
        Ativo: true,
	}

	fmt.Printf("Nome = %s, Idade = %d e situação %t\n", cliente.Nome, cliente.Idade, cliente.Ativo)
	desativacao(cliente)
}
