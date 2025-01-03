# Calculator

![Go](https://img.shields.io/badge/Go-1.23.4-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Descrição

Este é um projeto de uma calculadora básica escrita em Go. Ele suporta operações de adição, subtração, multiplicação e divisão para números inteiros e de ponto flutuante.


## Funcionalidades

- **Adição**: `sum`
- **Subtração**: `sub`
- **Multiplicação**: `mult`
- **Divisão**: `div`

## Como Usar

### Pré-requisitos

- Go 1.23.4 ou superior

### Compilação

Para compilar o projeto, execute:

```sh
go build -o main.exe cmd/cli/main.go
```

### Execução
Para executar a calculadora, use o comando:

```sh
./main.exe <method> <parametro1> <parametro2>
```

#### Exemplos:

```sh
./main.exe sum 2 3
./main.exe sub 5 2
./main.exe mult 3 4
./main.exe div 10 2
```

## Testes
Para rodar os testes, execute:

```sh
go test ./internal/math/...
```

## Licença
Este projeto está licenciado sob a licença MIT - veja o arquivo LICENSE para mais detalhes.

## Autor
Fabio S. Oliveira


Este README fornece uma visão geral clara do projeto, incluindo como compilar, executar e testar o código, além de informações sobre como contribuir.
