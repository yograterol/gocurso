package main

import (
	"fmt"
	"math/rand"
)

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"
const testConst = "Test"

func main() {
	var name string
	lastName := "<Modificar con mi apellido>"
	var number = sum(50, 50)
	a, b, c := getVariables()
	fa, fb := getVariablesFloat()
	name = getName()
	unicodeString := getUnicodeString()
	fmt.Printf(helloWorld, name, lastName)
	fmt.Println("* Esta es otra forma de imprimir con un salto de línea")
	fmt.Println("int: Número entero de al menos 32 bits en adelante", a)
	fmt.Println("int32: Número entero de máximo 32 bits", b)
	fmt.Println("int32: Número entero de máximo 64 bits", c)
	fmt.Println("float32: Número flotante con precisión de 32 bits", fa)
	fmt.Println("float64: Número flotante con precisión de 64 bits", fb)
	fmt.Println(unicodeString, "Contar letras de un string 'hola'", len("Hola"))
	fmt.Println("Primera letra de mi nombre en ASCII:", name[0])
	fmt.Println("Primera letra de mi nombre:", string(name[0]))
	fmt.Println(number, a, b, c, testConst)
	array()
	slice()
	ifTest()
	switchTest()
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

func getUnicodeString() string {
	return "もしもし！"
}

func sum(a int, b int) int {
	return a + b
}

func array() {
	var arr1 [2]string
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr1[0] = "mi"
	arr1[1] = "nombre"
	fmt.Println("Array declarado con var", arr1)
	fmt.Println("Array declarado con inferencia de tipo", arr2)
}

func slice() {
	var slic1 []string
	slic2 := []int{1, 2, 3, 4, 5}
	// append agrega un elemento al final del Slice
	slic1 = append(slic1, "mi")
	slic1 = append(slic1, "nombre")
	fmt.Println("Slice declarado con var", slic1)
	fmt.Println("Slice declarado con inferencia de tipo", slic2)
}

func ifTest() {
	number := 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if number >= 5 && number <= 10 {
		fmt.Println("El número es mayor o igual a 5")
	} else if number < 5 && number > 0 {
		fmt.Println("El número es menor que 5")
	} else {
		fmt.Println("El número no está dentro del rango sugerido")
	}

	if number2 := rand.Int31n(10); number2 > 5 {
		fmt.Println("[Aleatorio] El número es mayor a 5")
	} else {
		fmt.Println("[Aleatorio] El número es menor o igual a 5")
	}
}

func switchTest() {
	number := 0
	fmt.Print("[Switch] Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	// Switch sin condicionales ni comparadores
	switch number {
	case 0:
		fmt.Println("El número es 0")
	default:
		fmt.Println("El número es mayor o menor a 0")
	}

	switch {
	case number >= 5 && number <= 10:
		fmt.Println("[Switch con condicionales y comparadores] El número es mayor o igual a 5")
	case number < 5 && number > 0:
		fmt.Println("[Switch con condicionales y comparadores] El número es menor que 5")
	default:
		fmt.Println("[Switch con condicionales y comparadores] El número no está dentro del rango sugerido")
	}
}
