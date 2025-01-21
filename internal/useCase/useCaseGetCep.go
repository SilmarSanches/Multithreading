package useCase

import (
	"context"
	"errors"
	"time"

	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/entities"
	"github.com/silmarsanches/multithreading/internal/infra/services"
	"github.com/silmarsanches/multithreading/internal/mappers"
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

func (u *UseCaseGetCep) GetCep(ctx context.Context, cep string) (entities.CommonCepDto, error) {
	channelViaCep := make(chan entities.ViaCepDto)
	channelBrasilApi := make(chan entities.BrasilApiDto)

	go func() {
		data, err := u.ViaCepService.GetCep(ctx, cep)
		if err != nil {
			close(channelViaCep)
			return
		}
		channelViaCep <- data
	}()

	go func() {
		data, err := u.BrasilApiService.GetCep(ctx, cep)
		if err != nil {
			close(channelBrasilApi)
			return
		}
		channelBrasilApi <- data
	}()

	for {
		select {
		case data, ok := <-channelViaCep:
			if !ok {
				return mappers.MapViaCepToCommon(data), nil
			}
		case data, ok := <-channelBrasilApi:
			if !ok {
				return mappers.MapBrasilApiToCommon(data), nil
			}
		case <-time.After(time.Second * 1):
			return entities.CommonCepDto{}, errors.New("Timeout de 1s excedido ao consultar os serviços ViaCep e BrasilApi")
		}
	}
}
