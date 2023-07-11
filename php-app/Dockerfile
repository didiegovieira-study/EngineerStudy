FROM php:latest

# Copia todos os arquivos e diretórios para o diretório /var/www/html dentro do container
COPY . /var/www/html

# Define o diretório de trabalho como /var/www/html/public
WORKDIR /var/www/html/public

# Expõe a porta 8080 para acesso externo
EXPOSE 8080

# Define o comando a ser executado quando o container for iniciado
CMD ["php", "-S", "0.0.0.0:8080"]
