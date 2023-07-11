# Estudo Engineer

Este projeto é um ambiente de desenvolvimento para estudo de engenharia de software, utilizando Docker, MongoDB, MYSQL, PHP e Go. Ele permite a execução de uma aplicação PHP e uma aplicação Go de forma isolada e simplificada, facilitando o desenvolvimento e teste de funcionalidades.

![Imagem do Projeto](https://i.pinimg.com/originals/76/09/46/7609468e97e15d1da8d14d534be7366c.gif)
## Pré-requisitos

Certifique-se de ter o Docker e o Docker Compose instalados na sua máquina antes de prosseguir.

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Configuração

1. Clone este repositório para a sua máquina local.
2. Navegue até o diretório do projeto: `cd EngineerStudy`.
  - Navegue até o diretório do ambiente Go: `cd go-app`.
  - Execute o comando `go mod vendor`
  - Execute o comando `go mod download`
  - Execute o comando `go mod verify`
  - Execute o comando `cd ..`
3. Execute o comando `docker-compose build` para construir as imagens dos serviços PHP e Go.
4. Execute o comando `docker-compose up -d` para iniciar os containers em segundo plano.
5. Acesse a aplicação PHP no navegador através do endereço `http://localhost:8080`.
6. Acesse a aplicação Go no navegador através do endereço `http://localhost`.

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

## Recursos Adicionais

- O ambiente Docker fornece um servidor MongoDB e um servidor MySQL para uso em suas aplicações. Você pode acessar esses bancos de dados por meio das configurações padrão de conexão.

- A aplicação PHP está configurada para usar o diretório `/var/www/html` como raiz do servidor web. Você pode adicionar seu código PHP nesse diretório e ele será automaticamente servido pelo servidor embutido do PHP.

- A aplicação Go está configurada para usar a porta `8080` como porta padrão. Você pode modificar a aplicação Go de acordo com suas necessidades no diretório `go-app`.

- O arquivo `docker-compose.yml` define a configuração dos serviços e as dependências entre eles. Sinta-se à vontade para personalizar o arquivo de acordo com suas necessidades específicas.

- O projeto utiliza o recurso `vendor` do Go para armazenar as dependências do módulo. Isso permite que o processo de construção do contêiner seja mais rápido e evita a necessidade de baixar as dependências durante a construção.

- O ambiente de desenvolvimento foi configurado para facilitar o processo de estudo e experimentação de engenharia de software. Sinta-se à vontade para adicionar mais funcionalidades, testar diferentes tecnologias ou explorar conceitos específicos em suas aplicações PHP e Go.

- Lembre-se de encerrar o ambiente Docker quando não estiver em uso, utilizando o comando `docker-compose down`, para liberar recursos do sistema.

Divirta-se estudando engenharia de software com este ambiente de desenvolvimento prático e flexível!

