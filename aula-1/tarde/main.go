package main

import (
	"fmt"
	"strings"
)

func ex1(word string) {
	split := strings.Split(word, "")

	fmt.Println("Número de letras da palavra", word, "é igual a", len(split))

	for _, x := range split {
		fmt.Println(x)
	}

}

func ex2(idade int, empregado bool, anosExperiencia float32, salario float32) {
	if idade > 22 && empregado == true && anosExperiencia > 1 && salario <= 100.000 {
		fmt.Println("empréstimo concedido com juros")
	} else if idade > 22 && empregado == true && anosExperiencia > 1 && salario > 100.000 {
		fmt.Println("empréstimo concedido sem juros")
	} else {
		fmt.Println("empréstimo NÃO concedido")
	}
}

func ex3(mes int) {

	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Print("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	}
}

func ex4() {
	var employees = map[string]int{"Beijamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Printf("A idade de Beijamin é %v\n", employees["Beijamin"]) //idade de um dos funcionários

	fmt.Printf("\nO comprimento do Map é %v\n", len(employees)) //comprimento do mapa

	maior21 := 0
	for _, value := range employees {
		if value > 21 {
			maior21++
		}
	}
	fmt.Println("\nFuncionários com mais de 21 anos:\n", maior21)

	fmt.Println("\nO funcionário Frederico foi incluído à empresa. Quadro atual:")
	employees["Frederico"] = 25 // inserindo funcionário
	fmt.Println(employees)

	fmt.Println("\nO funcionário Pedro foi desligado da empresa. Quadro atual:")
	delete(employees, "Pedro") // deletando funcionário
	fmt.Println(employees)
}

func main() {
	fmt.Println("-- Exercício 1 --")
	ex1("Banana")
	fmt.Println("\n-- Exercício 2 --")
	ex2(24, true, 3, 90.000)
	fmt.Println("\n-- Exercício 3 --")
	ex3(9)
	fmt.Println("\n-- Exercício 4 --")
	ex4()
}
