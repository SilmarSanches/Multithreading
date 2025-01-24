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
O server foi desenvolvido usando Wire para injeção de dependências, e o comando acima irá gerar o arquivo wire_gen.go que contém as dependências necessárias para a aplicação.

Alem destes pacotes, foram utilizados o pacote chi e viper, para expor o endpoint e controlar as variáveis de ambiente, respectivamente