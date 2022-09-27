# Data Sanitizer

Este é um serviço de manipulação de dados e persistência em base de dados relacional desenvolvido utilizando Go e o banco de dados Postgres. 

## Instalação
Inicializar o banco de dados através do [docker-compose](https://docs.docker.com/compose/install/)
```bash
docker-compose up -d
```
Acessar o banco de dados. É necessário digitar a senha que foi previamente configurada no arquivo **database.env**
```bash
docker exec -it datasanitizer_db_1 psql -Udatasan -W
``` 

Neste projeto existem 2 branches além da 'main', cada uma delas aborda a solução de maneira diferente. Na main está a pior delas. Na branch changeDataLog está a última e melhor delas. Nesta última foram utilizadas threads para melhorar a velocidade de processamento, enquanto o arquivo foi lido linha a linha, passando os dados para um buffered channel, permitindo um controle maior do consumo de memória ram.