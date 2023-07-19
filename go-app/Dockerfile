# Use a imagem base oficial do Go
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Copie todo o código fonte do aplicativo para o diretório de trabalho
COPY . .

# Compile o aplicativo Go
RUN go mod tidy && go build -o go-app ./framework/cmd/server
#RUN go mod tidy && go build ./framework/cmd/server


# Defina a porta que o aplicativo irá expor
EXPOSE 9000

# Execute o aplicativo quando o contêiner for iniciado
CMD ["./go-app"]
