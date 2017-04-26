package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	var name string
	lastName := "<Modificar con mi apellido>"
	var number = 100
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, name, lastName)
	fmt.Println("* Esta es otra forma de imprimir con un salto de línea")
	fmt.Println(number, a, b, c, testConst)
}
