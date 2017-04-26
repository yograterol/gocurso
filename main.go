package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	var name string
	lastName := "<Modificar con mi apellido>"
	var number = sum(50, 50)
	a, b, c := getVariables()
	fa, fb := getVariablesFloat()
	name = getName()
	fmt.Printf(helloWorld, name, lastName)
	fmt.Println("* Esta es otra forma de imprimir con un salto de línea")
	fmt.Println("int: Número entero de al menos 32 bits en adelante", a)
	fmt.Println("int32: Número entero de máximo 32 bits", b)
	fmt.Println("int32: Número entero de máximo 64 bits", c)
	fmt.Println("float32: Número flotante con precisión de 32 bits", fa)
	fmt.Println("float64: Número flotante con precisión de 64 bits", fb)
	fmt.Println(number, a, b, c, testConst)
}

func getName() string {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int32, int64) {
	return 9223372036854775807, 2147483647, 9223372036854775807
}

func getVariablesFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func sum(a int, b int) int {
	return a + b
}
