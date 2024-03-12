package main

import "fmt"

func main() {
	name := "Letícia"
	var version float32 = 1.1
	fmt.Println("Oie", name)
	fmt.Println("A versão atual desse programa é:", version)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("comando", comando)

	if comando == 1 {

	} else if comando == 2 {

	} else if comando == 0 {

	} else {
		fmt.Println("Comando não reconhecido, tente novamente!")
	}

}
