package main

import (
	"errors"
	"fmt"
)

func ex1(salary float64) float64 {
	var tax float64
	if salary <= 50000 {
		tax = salary * 0.17
	} else {
		tax = salary * 0.27
	}
	return tax
}

func ex2(score ...float64) float64 {
	sum := 0.0
	for _, s := range score {
		sum += s
	}
	return sum / float64(len(score))
}

func ex3(category string, hours int) (float64, error) {
	if category == "C" {
		return 1000 * float64(hours), nil
	}
	if category == "B" {
		salary := 1500 * float64(hours)
		if hours > 160 {
			return salary * 1.2, nil
		}
		return salary, nil
	}
	if category == "A" {
		salary := 3000 * float64(hours)
		if hours > 160 {
			return salary * 1.5, nil
		}
		return salary, nil
	}
	return 0.0, errors.New("Categoria Inválida")
}

func main() {
	fmt.Println("-- Exercício 1 --")
	fmt.Println("O imposto será de ", ex1(70000), "reais")
	fmt.Println("\n-- Exercício 2 --")
	fmt.Printf("Cálculo da média: %.2f\n", ex2(10, 5, 6))
	fmt.Println("\n-- Exercício 3 --")
	salary, _ := ex3("B", 176)
	fmt.Printf("Salario: R$ %.2f\n", salary)
}
