//é necessário para saber onde está o main
package main

// Import block, it allows multiple packages to be imported without
// repeating the "import" keyword

import (
	"fmt"

	"github.com/godarksystem/collections"
	"github.com/godarksystem/functions"
)

/*
Função main
When part of the main package, identifies the entry point of an application
*/
func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Olá Mundo")
	fmt.Println("Hola Mundo")

	collections.Arrays()
	collections.Olamundo3()

	functions.MainDemoWebservice()

}

//Ola mundo
