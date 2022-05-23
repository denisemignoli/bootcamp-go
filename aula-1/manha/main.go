package main

import "fmt"

func ex1(name string, age uint) {
	//var name string = "Denise"
	//var age uint = 28
	fmt.Println("Informe seu nome: ", name)
	fmt.Println("Informe sua idade: ", age)
}

func ex2() {
	var temperature, humidity int = 22, 18
	var pressureAtm float32 = 1013.45
	fmt.Println("Temperatura: ", temperature)
	fmt.Println("Umidade: ", humidity)
	fmt.Println("Pressão atmosférica: ", pressureAtm)
}

/*
func ex3() {
	var lnome string
	correto, apenas falta atribuir valor

	var sobrenome string
	correto, apenas falta atribuir valor

	var int idade
	incorreto, o nome da variável vem antes do tipo. Correção: var idade int

	lsobrenome := 6
	incorreto, está incoerente o valor atribuído com ao que se refere o nome da variável. Correção: lsobrenome := "Ferreira"

	var licenca_para_dirigir = true
	incorreto, para declaração de variáveis com mais palavras é recomendado utilizar o camelCase e também faltou informar o tipo de dado, que no caso é o boolean. Correção: var licencaParaDirigir = true

	var estatura da pessoa int
	incorreto, para declaração de variáveis com mais palavras é recomendado utilizar o camelCase. Correção: var estaturaDaPessoa

	quantidadeDeFilhos := 2
	correto. Vale ressltar que esse tipo de declaração usando o operador :=, só é válido quando dentro de uma função.
}

func ex4(){
	var sobrenome string = "Silva"
	correto

	var idade int = "25"
	incorreto, o tipo determinado é um inteiro, portanto o valor atribuído não deve estar dentro de aspas

	boolean := "false";
	incorreto, nesse caso caso de fato a intenção fosse ler um valor booleano, o false deveria estar sem as aspas

	var salario string = 4585.90
	incorreto, por ser tratar de salário e normalmente terem casas decimais, o tipo de variável deveria ser em float

	var nome string = "Fellipe"
	correto
}
*/

func main() {
	fmt.Println("-- Exercício 1 --")
	ex1("Denise", 28)
	fmt.Println("\n-- Exercício 2 --")
	ex2()
}
