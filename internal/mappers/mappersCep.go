package mappers

import "github.com/silmarsanches/multithreading/internal/entities"

func MapBrasilApiToCommon(dto entities.BrasilApiDto) entities.CommonCepDto {
    return entities.CommonCepDto{
        Cep:          dto.Cep,
        Street:       dto.Street,
        Neighborhood: dto.Neighborhood,
        City:         dto.City,
        State:        dto.State,
        Source:       "BrasilApi",
    }
}

func MapViaCepToCommon(dto entities.ViaCepDto) entities.CommonCepDto {
    return entities.CommonCepDto{
        Cep:          dto.Cep,
        Street:       dto.Logradouro,
        Neighborhood: dto.Bairro,
        City:         dto.Localidade,
        State:        dto.Uf,
        Source:       "ViaCep",
    }
}