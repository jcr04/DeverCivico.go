# DeverCivico.go

DeverCivico.go é uma API destinada a promover o engajamento cívico no âmbito de políticas públicas e eventos civis. Foi criada para fornecer uma plataforma onde cidadãos possam reportar problemas, participar de discussões e obter informações governamentais.

## Funcionalidades

A API fornece os seguintes endpoints para interação:

1. **Cadastro de Cidadão**:
   - Endpoint: `/cadastrar`
   - Método: `POST`
   - Body da requisição: 
     ```json
     {
        "nome": "Seu Nome",
        "email": "seuemail@example.com",
        "senha": "suasenha"
     }
     ```

2. **Login de Cidadão**:
   - Endpoint: `/login`
   - Método: `POST`
   - Body da requisição:
     ```json
     {
        "email": "seuemail@example.com",
        "senha": "suasenha"
     }
     ```

3. **Reportar Problema**:
   - Endpoint: `/reportar-problema`
   - Método: `POST`
   - Body da requisição:
     ```json
     {
        "descricao": "Descrição do problema",
        "localizacao": "Localização do problema",
        "data_hora": "2023-10-17T10:00:00Z"
     }
     ```

4. **Listar Discussões**:
   - Endpoint: `/discussoes`
   - Método: `GET`

5. **Criar Discussão**:
   - Endpoint: `/criar-discussao`
   - Método: `POST`
   - Body da requisição:
     ```json
     {
        "titulo": "Título da discussão",
        "descricao": "Descrição da discussão",
        "data_hora": "2023-10-17T10:00:00Z"
     }
     ```

6. **Obter Informações Governamentais**:
   - Endpoint: `/informacoes-governamentais`
   - Método: `GET`

## Como Usar

1. **Configuração**:
   - Certifique-se de ter o Go instalado em sua máquina.
   - Clone o repositório e navegue até o diretório do projeto.
   - Configure o banco de dados conforme as instruções fornecidas.

    ## Configuração do Banco de Dados

    Para o funcionamento correto da API, é necessário configurar um banco de dados PostgreSQL. Abaixo estão as instruções para a criação do banco de dados e das tabelas necessárias.

    ### Criação do Banco de Dados

    ```sql
    CREATE DATABASE devercivico;
    ```
    #### Criação das Tabelas
    1. Tabela de Cidadãos:

    ```sql
    CREATE TABLE cidadao (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    senha VARCHAR(255) NOT NULL
    );
    ```
    2. Tabela de Problemas Reportados:
    ```sql
    CREATE TABLE problema_reportado (
    id SERIAL PRIMARY KEY,
    descricao TEXT NOT NULL,
    localizacao VARCHAR(255) NOT NULL,
    data_hora TIMESTAMP NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'pendente'
    );
    ```
    3. Tabela de Discussões:
    ```sql 
    CREATE TABLE discussao (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descricao TEXT NOT NULL,
    data_hora TIMESTAMP NOT NULL
    );
    ```
    4. Tabela de Informações Governamentais:
    ```sql
    CREATE TABLE informacoes_governamentais (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descricao TEXT NOT NULL,
    data_hora TIMESTAMP NOT NULL
    );
    ```
    #### Conexão com a API
    Após a criação do banco de dados e das tabelas, atualize o arquivo de configuração da API com as credenciais do seu banco de dados. 
    
    ```go
    const (
    DB_USER     = "seu_usuario"
    DB_PASSWORD = "sua_senha"
    DB_NAME     = "devercivico"
    )
    ```


2. **Execução**:
   - Execute o comando `go run main.go` para iniciar o servidor.
   - A API estará disponível em `http://localhost:8080`.

3. **Teste**:
   - Utilize um cliente HTTP como Postman ou cURL para enviar requisições à API e verificar as respostas.

## Contribuindo

Sinta-se à vontade para contribuir com o projeto reportando bugs, sugerindo melhorias ou enviando Pull Requests.
