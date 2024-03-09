# Fatec Challenge - Arquitetura Orientada a Serviços

## Introdução

Com o intuito de demonstrar proficiência na matéria em questão, foi solicitada a criação de uma Web-API **REST** que tivesse os seguintes comportamentos:

- Consumir ao menos duas Web-API's externas
- Conectar em um banco de dados
- Retornar os registros de uma tabela previamente criada no banco

## Entrega

O projeto realizado é uma simples API REST que possui os seguintes endpoints:

### ViaCEP API

- **GET** /zip/**{zip}**
  Este endpoint fará a consulta do CEP inputado na URL ( **{zip}** ) e devolverá as informações sobre o mesmo.

  - **Exemplo de requisição:**

  ```shell
  curl localhost:8080/zip/05520-200
  ```

  - **Exemplo de resposta:**

  ```json
  {
    "cep": "05520-200",
    "logradouro": "Avenida Professor Francisco Morato",
    "complemento": "de 4232 a 4886 - lado par",
    "bairro": "Vila Sônia",
    "localidade": "São Paulo",
    "uf": "SP",
    "ibge": "3550308",
    "gia": "1004",
    "ddd": "11",
    "siafi": "7107"
  }
  ```

### Nationality API

- **GET** /nationality/**{name}**
  Este endpoint fará a consulta do nome inputado na URL ( **{name}** ) e devolverá as informações sobre o mesmo.

  - **Exemplo de requisição:**

  ```shell
  curl localhost:8080/name/caio
  ```

  - **Exemplo de resposta:**

  ```json
  {
    "count": 1468,
    "name": "caio",
    "country": [
      {
        "country_id": "BR",
        "probability": 0.24218256035817998
      },
      {
        "country_id": "IT",
        "probability": 0.21022431867452152
      },
      {
        "country_id": "AO",
        "probability": 0.19068948642360772
      },
      {
        "country_id": "PT",
        "probability": 0.14301595237998632
      },
      {
        "country_id": "JP",
        "probability": 0.024863386116953236
      }
    ]
  }
  ```

### Users API

- **GET** /users
  Este endpoint listará os usuários existentes na tabela de usuários.

  - **Exemplo de requisição:**

  ```shell
  curl localhost:8080/users
  ```

  - **Exemplo de resposta:**

  ```json
  [
    {
      "id": 1,
      "name": "Gopher",
      "email": "hello@gopher.com"
    },
    {
      "id": 2,
      "name": "Halie",
      "email": "Lura33@yahoo.com"
    },
    {
      "id": 3,
      "name": "Laurine",
      "email": "laurine@gmail.com"
    }
  ]
  ```

- **GET** /users/**{id}**
  Este endpoint listará o usuário inputado na URL ( **{id}** ) e devolverá as informações sobre o mesmo.

  - **Exemplo de requisição:**

  ```shell
  curl localhost:8080/users/id
  ```

  - **Exemplo de resposta:**

  ```json
  [
    {
      "id": 1,
      "name": "Gopher",
      "email": "hello@gopher.com"
    }
  ]
  ```

- **POST** /users
  Este endpoint criará o usuário e devolverá as informações criadas sobre o mesmo.

  - **Exemplo de requisição:**

  ```shell
  curl -X POST -H "Content-Type: application/json" -d '{"name":"Caio Salgado Nepomuceno", "email":"csn.caio@gmail.com"}' http://localhost:8080/users
  ```

  - **Exemplo de resposta:**

  ```json
  {
    "id": 1,
    "name": "Caio Salgado Nepomuceno",
    "email": "csn.caio@gmail.com"
  }
  ```

- **PUT** /users/**{id}**
  Este endpoint atualizará os dados do usuário inputado na URL ( **{id}** ) e devolverá as informações atualizadas sobre o mesmo.

  - **Exemplo de requisição:**

  ```shell
  curl -X PUT -H "Content-Type: application/json" -d '{"name":"Caio Salgado Nepomuceno", "email":"csn.caio@gmail.com"}' http://localhost:8080/users/1
  ```

  - **Exemplo de resposta:**

  ```json
  {
    "id": 1,
    "name": "Caio Salgado Nepomuceno",
    "email": "csn.caio@gmail.com"
  }
  ```

- **DELETE** /users/**{id}**
  Este endpoint removerá o usuário inputado na URL ( **{id}** ).

  - **Exemplo de requisição:**

  ```shell
  curl -X DELETE http://localhost:3000/users/1
  ```

  - **Exemplo de resposta:**

  ```http
  200 OK
  ```

## Ferramentas utilizadas

- Golang 1.22
- Postgres 16.2
- Docker
- Docker Compose
- Env variables
