# Estudo Engineer

Este projeto é um ambiente de desenvolvimento para estudo de engenharia de software, utilizando Docker, MongoDB, PHP e Go. Ele permite a execução de uma aplicação PHP e uma aplicação Go de forma isolada e simplificada, facilitando o desenvolvimento e teste de funcionalidades.

## Pré-requisitos

Certifique-se de ter o Docker e o Docker Compose instalados na sua máquina antes de prosseguir.

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto: `cd EngineerStudy`.
3. Execute o comando `docker-compose build` para construir as imagens dos serviços PHP e Go.
4. Execute o comando `docker-compose up -d` para iniciar os containers em segundo plano.
5. Acesse a aplicação PHP no navegador através do endereço `http://localhost`.
6. Acesse a aplicação Go no navegador através do endereço `http://localhost:8080`.

## Estrutura do Projeto

O projeto possui a seguinte estrutura de diretórios:

- docker-compose.yml
- go-app/
  - Dockerfile
  - go.mod
  - go.sum
  - ... (outros arquivos do Go)
- php-app/
  - Dockerfile
  - ... (arquivos do PHP)
