# Use a imagem base oficial do Go
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do módulo Go, comando abaixo deve ser feito a mão
# RUN go mod download

# Copie todo o código fonte do aplicativo para o diretório de trabalho
COPY . .

COPY vendor/ ./vendor/

# Compile o aplicativo Go
RUN go build -mod vendor -o go-app

# Defina a porta que o aplicativo irá expor
EXPOSE 8080

# Execute o aplicativo quando o contêiner for iniciado
CMD ["./go-app"]
