# First Golang API

Este projeto é uma API desenvolvida em Go (Golang) para fins de estudo da linguagem com autenticação com JWT e conexão ao banco de dados MySQL.

## Pré-requisitos

![VSCode](https://img.shields.io/badge/Visual_Studio_Code-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-4479A1.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

- [VSCode](https://code.visualstudio.com/Download) Visual Studio Code ou ide da sua preferência
- [Go](https://golang.org/dl/) 1.18 ou superior instalado
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) (opcional, mas recomendado)
- MySQL rodando localmente ou em container

## Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/first-golang-api.git
cd first-golang-api
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo ou renomeia o `.env-example`:

```
PORT =

MYSQL_USER=
MYSQL_PASSWORD=
MYSQL_HOST=
MYSQL_TABLE=
MYSQL_DATABASE=
MYSQL_PORT=

SECRET_KEY=
```

> Ajuste os valores conforme sua configuração.

### 3. Suba o MySQL com Docker (opcional)

Se quiser usar Docker, crie um arquivo `docker-compose.yml`:

```yaml
version: "3.8"
services:
  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: seu_banco
      MYSQL_USER: seu_usuario
      MYSQL_PASSWORD: sua_senha
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
```

Inicie o container:

```bash
docker-compose up -d
```

### 4. Instale as dependências Go

```bash
go mod tidy
```

## Uso

### 1. Execute a API

```bash
go run main.go
```

A API estará disponível em `http://localhost:5000` (ou porta configurada).

### 2. Teste os endpoints

Use ferramentas como [Postman](https://www.postman.com/) ou [curl](https://curl.se/) para testar os endpoints da API.

Exemplo com curl:

```bash
curl http://localhost:5000/seu-endpoint
```

## Estrutura do Projeto

```
first-golang-api/
├── main.go
├── go.mod
├── go.sum
├── src/
│   ├── authentication/
│   ├── config/
│   ├── controllers/
│   ├── database/
│   ├── middlewares/
│   ├── models/
│   ├── repositories/
│   ├── responses/
│   ├── routing/
│   └── routes/
└── ...
```

## Observações

- Certifique-se de que o banco de dados está rodando e acessível.
- Ajuste as variáveis de ambiente conforme necessário.

## Instalação

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/first-golang-api.git
cd first-golang-api
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
MYSQL_USER=seu_usuario
MYSQL_PASSWORD=sua_senha
MYSQL_DATABASE=seu_banco
MYSQL_HOST=localhost
MYSQL_PORT=3306
```

> Ajuste os valores conforme sua configuração.

### 3. Suba o MySQL com Docker (opcional)

Se quiser usar Docker, crie um arquivo `docker-compose.yml`:

```yaml
version: "3.8"
services:
  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: seu_banco
      MYSQL_USER: seu_usuario
      MYSQL_PASSWORD: sua_senha
    ports:
      - "3306:3306"
```

Inicie o container:

```bash
docker-compose up -d
```

### 4. Instale as dependências Go

```bash
go mod tidy
```

## Uso

### 1. Execute a API

```bash
go run main.go
```

A API estará disponível em `http://localhost:5000` (ou porta configurada).

---

## Rotas da API (Users - CRUD)

### 1. Criar um registro (Create)

- **POST** `/users`
- **Body (JSON):**
  ```json
  {
    "Name": "João Silva",
    "NickName": "JSilva",
    "Email": "jsilva@test.com",
    "Phone": "21 97500-5810",
    "Password": "123456"
  }
  ```

### 2. Listar todos os registros (Read All)

- **GET** `/users`
- **Resposta (JSON):**
  ```json
  [
    {
      "id": 1,
      "name": "Gilberto Medeiros",
      "nickname": "Lobo-rio",
      "email": "gmedeiros@test.com",
      "phone": "21 97565-5000",
      "createdat": "2025-06-18T18:05:52-03:00"
    },
    {
      "id": 2,
      "name": "João Silva",
      "nickname": "JSilva",
      "email": "jsilva@test.com",
      "phone": "21 97500-5810",
      "createdat": "2025-06-18T19:41:00-03:00"
    }
  ]
  ```

### 3. Buscar um registro por ID (Read One)

- **GET** `/users/{id}`
- **Resposta (JSON):**
  ```json
  {
    "id": 2,
    "name": "João Silva",
    "nickname": "JSilva",
    "email": "jsilva@test.com",
    "phone": "21 97500-5810",
    "createdat": "2025-06-18T19:41:00-03:00"
  }
  ```

### 4. Atualizar um registro (Update)

- **PUT** `/users/{id}`
- **Body (JSON):**
  ```json
  {
    "Name": "GilbertoSilva",
    "NickName": "LoboRio",
    "Email": "gsilva@test.com"
  }
  ```

### 5. Remover um registro (Delete)

- **DELETE** `/users/{id}`

---

## Rotas da API (Notes - CRUD)

### 1. Criar um registro (Create)

- **POST** `/notes`
- **Body (JSON):**
  ```json
  {
    "title": "Where can I get some?",
    "description": "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in"
  }
  ```

### 2. Listar todos os registros (Read All)

- **GET** `/notes`
- **Resposta (JSON):**
  ```json
  [
    {
      "id": 4,
      "title": "Where can I get some?",
      "description": "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in",
      "author_id": 2,
      "author_nickname": "JDoe",
      "createdat": "2025-06-23T14:13:38-03:00"
    },
    {
      "id": 6,
      "title": "Where can I get some?",
      "description": "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in",
      "author_id": 2,
      "author_nickname": "JDoe",
      "createdat": "2025-06-23T14:13:41-03:00"
    }
  ]
  ```

### 3. Buscar um registro por ID (Read One)

- **GET** `/notes/{noteId}`
- **Resposta (JSON):**
  ```json
  {
    "id": 4,
    "title": "Where can I get some?",
    "description": "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in",
    "author_id": 2,
    "author_nickname": "JDoe",
    "createdat": "2025-06-23T14:13:38-03:00"
  }
  ```

### 4. Atualizar um registro (Update)

- **PUT** `/notes/{noteId}`
- **Body (JSON):**
  ```json
  {
    "Title": "Why do we use it?",
    "Description": "It is a long established fact that a reader will be distracted by the readable content"
  }
  ```

### 5. Remover um registro (Delete)

- **DELETE** `/notes/{noteId}`

---

## Observações

- Todos os endpoints que recebem dados esperam o conteúdo em JSON.
- Os campos obrigatórios podem variar conforme a modelagem do seu projeto.
- Certifique-se de que o banco de dados está rodando e acessível.
- Ajuste as variáveis de ambiente conforme necessário.
