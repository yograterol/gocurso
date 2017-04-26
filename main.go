package main

import "fmt"

func main() {
	var name string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go.\n", name)
}
