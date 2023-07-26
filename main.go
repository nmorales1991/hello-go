package main

import (
	"fmt"
	"github.com/nmorales1991/greetings-go"
	"github.com/nmorales1991/greetings-go/kuky"
	"github.com/nmorales1991/hello-go/choco"
	"log"
)

func main() {
	kuky.HolaKuky()
	choco.HolaChoco()

	message, err := greetings.Hello("Nicolás")
	if err != nil {
		log.SetPrefix("nombre: ")
		log.Fatal(err)
	}
	fmt.Println(message)
}
