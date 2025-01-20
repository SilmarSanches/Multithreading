package useCase

import (
	"context"
	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/infra/services"
)

type UseCaseGetCep struct {
	ViaCepService    services.HttpExternalServiceViaCep
	BrasilApiService services.ExternalServiceBrasilApi
	appConfig        *config.Config
}

func NewUseCaseGetCep(viaCepService services.HttpExternalServiceViaCep, brasilApiService services.ExternalServiceBrasilApi, appConfig *config.Config) *UseCaseGetCep {
	return &UseCaseGetCep{
		ViaCepService:    viaCepService,
		BrasilApiService: brasilApiService,
		appConfig:        appConfig,
	}
}

func (u *UseCaseGetCep) GetCep(ctx context.Context, cep string) (map[string]interface{}, error) {
	data, err := u.ViaCepService.GetCep(ctx, cep)
	if err != nil {
		return nil, err
	}

	return data, nil
}
