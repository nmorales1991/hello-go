package main

import (
	"fmt"
	"github.com/nmorales1991/greetings-go"
	// se agregó nuevo package al otro proyecto. Se hizo go get pero agregando un token de github para poder descargar el paquete, ya que es privado.
	"github.com/nmorales1991/greetings-go/functions"
	"github.com/nmorales1991/greetings-go/kuky"
	"github.com/nmorales1991/hello-go/choco"
	"log"
)

func main() {
	kuky.HolaKuky()
	choco.HolaChoco()
	functions.TestFunction()

	message, err := greetings.Hello("Nicolás")
	if err != nil {
		log.SetPrefix("nombre: ")
		log.Fatal(err)
	}
	fmt.Println(message)
}
