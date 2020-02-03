package main

import (
	"fmt"
)

var a, b, c, ruta int

func main() {
	menu(ruta, a, b, c)
}

func menu(ruta, a, b, c int) int {
ENTRADA:
	fmt.Println("\nDigite:\n 1 -> para Sumar\n 2 -> para Restar\n 3 -> para Multiplicar\n 4 -> para Dividir\n 5 -> para potenciar")
	fmt.Println(" 6 -> para salir")
	fmt.Scanln(&ruta)

	if ruta == 1 || ruta == 2 || ruta == 3 || ruta == 4 || ruta == 5 || ruta == 6 {
		fmt.Println("\nnúmero valido")
		switch ruta {
		case 1:
			fmt.Println("el total es: ", sumar(a, b))
			goto ENTRADA
		case 2:
			fmt.Println("el total es: ", restar(a, b))
			goto ENTRADA
		case 3:
			fmt.Println("el total es: ", multiplicar(a, b))
			goto ENTRADA
		case 4:
			fmt.Println("el total es: ", dividir(a, b))
			goto ENTRADA
		case 5:
			fmt.Println("el total es: ", potenciacion(a, b, c))
			goto ENTRADA
		case 6:
		}

	} else {
		fmt.Println("digite un numero valido")
		goto ENTRADA
	}
	return ruta
}

func sumar(a, b int) int {
	fmt.Println("\nvoy a sumar")
	fmt.Println("\nIngrese numero 1: ")
	fmt.Scanln(&a)
	fmt.Println("\nIngrese numero 2: ")
	fmt.Scanln(&b)
	return a + b
}

func restar(a, b int) int {
	fmt.Println("\nvoy a restar")
	fmt.Println("\nIngrese numero 1: ")
	fmt.Scanln(&a)
	fmt.Println("\nIngrese numero 2: ")
	fmt.Scanln(&b)
	return a - b
}

func multiplicar(a, b int) int {
	fmt.Println("\nvoy a multiplicar")
	fmt.Println("\nIngrese numero 1: ")
	fmt.Scanln(&a)
	fmt.Println("\nIngrese numero 2: ")
	fmt.Scanln(&b)
	return a * b
}

func dividir(a, b int) int {
	fmt.Println("\nvoy a dividir")
	fmt.Println("\nIngrese numero 1: ")
	fmt.Scanln(&a)
	fmt.Println("\nIngrese numero 2: ")
	fmt.Scanln(&b)
	return a / b
}
func potenciacion(a, b, c int) int {
	fmt.Println("\nvoy a potenciar")
	fmt.Println("\nIngrese la base: ")
	fmt.Scanln(&a)
	fmt.Println("\nIngrese el exponente: ")
	fmt.Scanln(&b)
	c = a
	for i := 1; i < b; i++ {
		a = a * c
	}
	return a
}
