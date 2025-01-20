package useCase

import (
	"context"
	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/infra/services"
)

type UseCaseGetCep struct {
	ViaCepService    services.ExternalServiceViaCepInterface
	BrasilApiService services.ExternalServiceBrasilApiInterface
	appConfig        *config.Config
}

func NewUseCaseGetCep(viaCepService services.ExternalServiceViaCepInterface, brasilApiService services.ExternalServiceBrasilApiInterface, appConfig *config.Config) *UseCaseGetCep {
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
