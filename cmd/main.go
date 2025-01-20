package main

import (
	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/web/controllers"
	"github.com/silmarsanches/multithreading/internal/web/routers"
	"github.com/silmarsanches/multithreading/internal/web/server"
	"log"
	"net/http"
)

func main() {
	appConfig, err := config.LoadConfig("./cmd")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo de configuração: %v", err)
	}

	useCaseGet := InitializeCepUseCase(appConfig)

	cepControllers := controllers.NewCepController(useCaseGet)
	cepRoutes := routers.CepRouters(cepControllers)
	srv := server.NewServer(cepRoutes)

	log.Println("Servidor iniciado na porta 8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
