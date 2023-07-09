# Use a imagem base oficial do Go
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do módulo Go
RUN go mod download

# Copie todo o código fonte do aplicativo para o diretório de trabalho
COPY . .

# Compile o aplicativo Go
RUN go build -o go-app

# Defina a porta que o aplicativo irá expor
EXPOSE 8080

# Execute o aplicativo quando o contêiner for iniciado
CMD ["./go-app"]
