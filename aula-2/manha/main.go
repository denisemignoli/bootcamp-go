package main

import (
	"errors"
	"fmt"
)

//----------- exercicio 1 -----------
/*
func ex1(salary float64) float64 {
	var tax float64
	if salary <= 50000 {
		tax = salary * 0.17
	} else {
		tax = salary * 0.27
	}
	return tax
}
*/

//----------- exercicio 2 -----------
/*
func ex2(score ...float64) float64 {
	sum := 0.0
	for _, s := range score {
		sum += s
	}
	return sum / float64(len(score))
}
*/

//----------- exercicio 3 -----------
/*
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
*/

//----------- exercicio 4 -----------
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operationMinimum(valores ...int) (int, error) {
	if len(valores) == 0 {
		return 0, errors.New("favor informar um número válido")
	}
	var valor int = valores[0]
	for _, x := range valores {
		if x < valor {
			valor = x
		}
	}
	return valor, nil
}

func operationMaximum(valores ...int) (int, error) {
	if len(valores) == 0 {
		return 0, errors.New("favor informar um número válido")
	}
	var valor int = valores[0]
	for _, x := range valores {
		if x > valor {
			valor = x
		}
	}
	return valor, nil
}

func operationAverage(valores ...int) (int, error) {
	if len(valores) == 0 {
		return 0, errors.New("favor informar um número válido")
	}
	var valor int = 0
	for _, x := range valores {
		valor += x
	}
	return valor / len(valores), nil
}

func getFunc(operador string) (func(valores ...int) (int, error), error) {
	switch operador {
	case "minimum":
		return operationMinimum, nil
	case "maximum":
		return operationMaximum, nil
	case "average":
		return operationAverage, nil
	default:
		return nil, errors.New("presta atenção cara")
	}
}

func main() {
	/*fmt.Println("-- Exercício 1 --")
	fmt.Println("O imposto será de ", ex1(70000), "reais")
	fmt.Println("\n-- Exercício 2 --")
	fmt.Printf("Cálculo da média: %.2f\n", ex2(10, 5, 6))
	fmt.Println("\n-- Exercício 3 --")
	salary, _ := ex3("B", 176)
	fmt.Printf("Salario: R$ %.2f\n", salary)
	*/
	//minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	//averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	//maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	a, error := getFunc(minimum)
	if error == nil {
		b := 0
		b, error = a(0, 3, 3, 4, 10, 2, 4, 5)
		if error == nil {
			fmt.Println(b)
		}
	}
	fmt.Println(error)
}
