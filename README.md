# Multithreading

Projeto de conclusão de pós-graduação (Desafio 2). Este projeto implementa uma API de consulta de CEP.

## Indice

1. [Descrição](#descrição)
2. [Run-Server](#run-server)

## Descrição
Este projeto é composto de uma API contendo um método chamado /buscacep. Este método recebe como parametro o CEP que se deseja pesquisar. 

## Run-Server
Para iniciar a aplicação, execute o comando abaixo:
```bash
    go run ./cmd/main.go ./cmd/wire_gen.go
```
Após o início da aplicação, acesse o endereço http://localhost:8080/buscacep/{cep} para realizar a consulta:
```bash
    http://localhost:8080/buscacep?cep=01153000
```
Depois da chamada a url acima, a aplicação irá retornar o endereço referente ao CEP informado no console, conforme exemplo:
```bash
    2025/01/24 01:49:55 Início GET /buscacep
    Cep:  01153000
    Street:  Rua Vitorino Carmilo
    Neighborhood:  Barra Funda
    City:  São Paulo
    State:  SP
    Source:  BrasilApi
    2025/01/24 01:49:55 Término GET /buscacep 20.56175ms

```
Além dos dados do CEP, temos um middleware que loga o tempo necessário para a execução da consulta nos dois endpoints.

O server foi desenvolvido usando Wire para injeção de dependências, e o comando acima irá gerar o arquivo wire_gen.go que contém as dependências necessárias para a aplicação.

Alem destes pacotes, foram utilizados o pacote chi e viper, para expor o endpoint e controlar as variáveis de ambiente, respectivamente