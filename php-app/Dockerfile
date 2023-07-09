FROM php:latest

# Copia todos os arquivos e diretórios para o diretório /var/www/html dentro do container
COPY . /var/www/html

# Define o diretório de trabalho como /var/www/html
WORKDIR /var/www/html

# Expõe a porta 80 para acesso externo
EXPOSE 80

# Define o comando a ser executado quando o container for iniciado
CMD ["php", "-S", "0.0.0.0:80"]
