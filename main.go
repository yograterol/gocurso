package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	var name string
	lastName := "<Modificar con mi apellido>"
	var number = sum(50, 50)
	a, b, c := getVariables()
	name = getName()
	fmt.Printf(helloWorld, name, lastName)
	fmt.Println("* Esta es otra forma de imprimir con un salto de línea")
	fmt.Println(number, a, b, c, testConst)
}

func getName() string {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int, int) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}
