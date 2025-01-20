//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/infra/services"
	"github.com/silmarsanches/multithreading/internal/useCase"
)

var setUseCaseCepDependency = wire.NewSet(
	services.NewHttpExternalServiceViaCep,
	wire.Bind(new(services.ExternalServiceViaCepInterface), new(*services.HttpExternalServiceViaCep)),
	services.NewHttpExternalServiceBrasilApi,
	wire.Bind(new(services.ExternalServiceBrasilApiInterface), new(*services.HttpExternalServiceBrasilApi)),
)

func InitializeCepUseCase(appConfig *config.Config) *useCase.UseCaseGetCep {
	wire.Build(
		setUseCaseCepDependency,
		useCase.NewUseCaseGetCep,
	)
	return &useCase.UseCaseGetCep{}
}
