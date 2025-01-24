package useCase

import (
	"context"
	"errors"
	"fmt"
	"github.com/silmarsanches/multithreading/config"
	"github.com/silmarsanches/multithreading/internal/entities"
	"github.com/silmarsanches/multithreading/internal/infra/services"
	"github.com/silmarsanches/multithreading/internal/mappers"
	"time"
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
	if u.ViaCepService == nil || u.BrasilApiService == nil {
		return entities.CommonCepDto{}, errors.New("Serviços ViaCep ou BrasilApi não foram inicializados")
	}

	channelViaCep := make(chan entities.ViaCepDto, 1)       // Canal com buffer de 1
	channelBrasilApi := make(chan entities.BrasilApiDto, 1) // Canal com buffer de 1

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recuperado de um panic em ViaCep: %v\n", r)
			}
		}()
		data, err := u.ViaCepService.GetCep(ctx, cep)
		if err != nil {
			return
		}
		channelViaCep <- data
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recuperado de um panic em BrasilApi: %v\n", r)
			}
		}()
		data, err := u.BrasilApiService.GetCep(ctx, cep)
		if err != nil {
			return
		}
		channelBrasilApi <- data
	}()

	for {
		select {
		case data := <-channelViaCep:
			return mappers.MapViaCepToCommon(data), nil
		case data := <-channelBrasilApi:
			return mappers.MapBrasilApiToCommon(data), nil
		case <-time.After(1 * time.Second):
			return entities.CommonCepDto{}, errors.New("Timeout de 1s excedido ao consultar os serviços ViaCep e BrasilApi")
		}
	}
}
