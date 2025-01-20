package main

import (
	"github.com/silmarsanches/multithreading/config"
	"log"
)

func main() {
	appConfig, err := config.LoadConfig("./server/cmd")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo de configuração: %v", err)
	}

	log.Printf(appConfig.URLBrasilAPI)
}
