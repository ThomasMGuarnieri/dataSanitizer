# Data Sanitizer

Este é um serviço de manipulação de dados e persistência em base de dados relacional desenvolvido utilizando Go e o banco de dados Postgres. 

## Instalação
Inicializar o banco de dados através do [docker-compose](https://docs.docker.com/compose/install/)
```bash
docker-compose up -d
```
Acessar o banco de dados. É necessário digitar a senha que foi previamente configurada no arquivo **database.env**
```bash
docker exec datasanitizer_db_1 psql -Udatasan -W
``` 