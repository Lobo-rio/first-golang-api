package main

import (
	"fmt"
	"log"
	"modules/src/config"
	"modules/src/routing"
	"net/http"
)

func main() {
	config.LoadConfig()
	r := routing.Generate()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
