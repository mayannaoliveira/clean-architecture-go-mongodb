

[![Udemy](https://img.shields.io/badge/Udemy-A435F0?style=flat&logo=udemy&logoColor=white&link=https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/)](https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/) ![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=flat&logo=nextdotjs&logoColor=white) [![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=Go&logoColor=white&link=https://go.dev/)](https://go.dev/) [![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=flat&logo=mongodb&logoColor=white&link=https://www.mongodb.com/)](https://www.mongodb.com/) [![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white&link=https://www.docker.com/)](https://www.docker.com/) [![Gin Framework](https://img.shields.io/badge/Gin%20Framework-008ECF?style=flat&logo=gin&logoColor=white&link=https://gin-gonic.com/)](https://gin-gonic.com/) [![Postman](https://img.shields.io/badge/Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/) [![GitHub](https://img.shields.io/badge/GitHub-181717?style=flat&logo=github&logoColor=white&link=https://github.com/mayannaoliveira)](https://github.com/mayannaoliveira) [![Studio3T](https://img.shields.io/badge/Studio%203T-17AF66?style=flat&logo=Studio3T&logoColor=white&link=https://studio3t.com/)](https://studio3t.com/)
[![VSCodium](https://img.shields.io/badge/VSCodium-2F80ED?style=flat&logo=vscodium&logoColor=white&link=https://vscodium.com/)](https://vscodium.com/)

<!-- Frontend -->
<!-- ![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=flat&logo=html5&logoColor=white) -->
<!-- ![CSS](https://img.shields.io/badge/CSS-663399?style=flat&logo=CSS&logoColor=white) -->
<!-- ![Next.js](https://img.shields.io/badge/Next.js-000000?style=flat&logo=nextdotjs&logoColor=white) -->

## API Go e MongoDB

Projetinho feito durante o curso [Clean Architecture na prática: Go, MongoDB e Docker
](https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/) na [Udemy](https://www.udemy.com/). O intuito desse repositório é de estudar Go e MongoDB.

>[!important]
>Repositório criado somente para estudar a Criação de uma API com Go e MongoDB.

<!-- TOC https://ecotrust-canada.github.io/markdown-toc/ -->
**Conteúdo:**

  * [1. Configuração do ambiente](#1-configura--o-do-ambiente)
  * [2. IDE](#2-ide)
  * [3. Docker](#3-docker)
  * [4. Introdução ao módulo](#4-introdução-ao-módulo)
  * [5. Criando API](#5-criando-api)
  * [6. Adicionando o banco de dados](#6-adicionando-o-banco-de-dados)
  * [7. Criando uma Task](#7-criando-uma-task)
  * [8. Listando as tasks](#8-listando-as-tasks)
  * [9. Atualizando as tasks](#9-atualizando-as-tasks)
  * [10. Deletando task](#10-deletando-task)
  * [11. Adicionando Swagger na aplicação](#11-adicionando-swagger-na-aplica--o)
- [Contribuidores](#contribuidores)

### 1. Configuração do ambiente

- Instalação do Go.
- Configuração variável ambiente.
- Checar a versão do Go com comando `go version`;

[![Download and install](https://img.shields.io/badge/Download%20and%20install-007f9f?style=flat&logo=go&logoColor=white)](https://go.dev/doc/install) [![Tutorial: Get started with Go](https://img.shields.io/badge/Tutorial:%20Get%20started%20with%20Go-007f9f?style=flat&logo=go&logoColor=white)](https://go.dev/doc/tutorial/getting-started)

- Para checar ser o go está rodando é simples:
    - Abra o terminal `cmd`.
    - Verifique a versão `go version`.
    - Crie um diretório `mkdir hello` e acese `cd hello`.
    - Insira um arquivo dentro chamado `hello.go` como abaixo:

    ```go
    package main
    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
    ```
    - Execute utilizando o comando `go run .`.
    - Verifique se retornou `Hello, World!`.

### 2. IDE

- Instalar a extenção do Go ![Go Team at Google](https://img.shields.io/badge/Go%20Team%20at%20Google-2F80ED?style=flat&logo=vscodium&logoColor=white) e [![Docker Microsoft](https://img.shields.io/badge/Docker%20Microsoft-2F80ED?style=flat&logo=vscodium&logoColor=white&link=https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker).
- Adicionar em `Arquivo > Preferências > Configurações` as configuração no VS Code:
```json
{
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.useLanguageServer": true
}
```
- Sempre uso a extensão [Docker da Microsoft](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker) para facilitar a visualização dos conteiners.

### 3. Docker

- Instalação do [![Docker ](https://img.shields.io/badge/Docker%20Docs-2496ED?style=flat&logo=Docker&logoColor=white&link=https://docs.docker.com/?_gl=1*bu2mvx*_gcl_au*NTA5ODMyMTIzLjE3NDgxMTAyNTE.*_ga*Mjk3NzcxNTE5LjE3NDgxMTAyNTE.*_ga_XJWPQMJYHQ*czE3NDg1NzA3ODUkbzIkZzEkdDE3NDg1NzA3OTAkajU1JGwwJGgw)](https://docs.docker.com/).
Checar Docker version com comando `docker --version`.

>[!note]
>A execução do projeto depende da criação do conteiner e que ele esteja executando.

#### 4. Introdução ao módulo

- Teste executado com old_main.go com o comando `go run old_main.go`.
- Verificar se no localhost `http://localhost:8000/` é possível visualizar a mensagem `Olá Mundo!`.

#### 5. Criando API

- Crie um novo arquivo `main.go` e adicione o código e execute no terminal `go run main`.
- Criar um servidor HTTP com a rota `/` que retorna a mensagem `Olá Mundo!`:
- Executea o comando `go run main.go`.
- Verifique se no localhost `http://localhost:8000/` é possível visualizar a mensagem `Olá Mundo!`.
- Caso erro devido ao erro de porta, altere o comando para `go run -p 8001 main.go` e execute novamente.
- Caso o Framework Gin não esteja instalado use o comando `go get -u github.com/gin-gonic/gin`.
- Execute o comando `go run main.go` para verficar `Olá,Mundo!` na página localhost.

**Projeto no Postman:**
- Baixe e instale o [![Postman](https://img.shields.io/badge/Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/).
- Importe a coleção e ambiente [![Postman](https://img.shields.io/badge/API%20do%20Curso%20Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/) para sua workspace no Postman.

>[!note] 
>As resquisições e o ambiente de teste foram salvos na pasta [Postman](\postman) para utilizar faça o `import` no Postman.

#### 6. Adicionando o banco de dados

- Verifique se o Docker está em execução.
- Crie um conteiner Docker com o comando `docker run -d -p 27017:27017 --name mongodb mongo:4`.
- Verifique com `docker ps` se a instalação foi realizada.
- Baixar pacotes para o projeto: `go mod tidy`.
- Verificar se foi gerado o arquivo `go.mod` em `go-crud-api\go.mod`.
- Continuar a parte do código para criar uma task.

#### 7. Criando uma Task

 - Criar novo repositório `models` e adicionar o arquivo `task.go`.
 - Adicionar uma struct `Task` dentro do arquivo `task.go`.
 - Adicionar um método `CreateTask` dentro do arquivo `main.go`.
 - Inserir tratamento de exceção para verificar a conexão com o banco de dados.
 - Criar o router para POST `Createtask`.
- Instalar a ferramenta [![Studio 3T](https://img.shields.io/badge/Studio%203T-17AF66?style=flat&logo=Studio3T&logoColor=white&link=https://studio3t.com/)](https://studio3t.com/) para visualizar as conexões do MongoDB.
- Para criar a conexão no Studio 3T bastas clicar em `Conect`.
- Inserir a url `mongodb://localhost:27017` para estabelecer a conexão local.
- Instalar o Postman para teste da API.
- No Postman é nescessário criar um novo worspace. 
- Adicione uma requisiçao `POST` para testar `tasks` `POST: http://localhost:8000/task`.
- Insira o JSON dentro do `Body > raw`:
```json
{
    "title":"Estudando Go",
    "description":"Exemplo do projeto",
    "completed":true,
    "created_at":"2022-03-18T14:10:00Z",
    "duo_date":"2022-03-18T14:10:00Z"
}
```
- Verifique se retornou o id similiar a esse:
```json
{
    "id": "683a043f3b9c9b864c7a67a4"
}
```
- Checar no Studio3T se os dados foram inseridos no MongoDB. Processo pode ser feito usando um JS `db.getCollection("tasks").find({})`.

#### 8. Listando as tasks

- Criar o bloco de `listTask` dentro do arquivo `main.go`.
- Inserir tratamento de exceção para verificar a conexão com o banco de dados.
- Inserir o router para `listTask`. 
- No Postman teste utilizando uma requisição `GET: http://localhost:8000/tasks` para listar todas as tasks.
- Verificar se retorna a lista de tasks adicionadas.

#### 9. Atualizando as tasks

- Implementar o código com a função 
- Criar uma request no Postman `PUT `
- Inserir o `Body > Raw > JSON`:
```json
    {
        "id": "<<ID>>",
        "title": "<<TÍTULO>>",
        "description": "<<DESCRIÇÃO>>",
        "completed": <<TRUE/FALSE>>,
        "due_date": "<<DATETIME>>",
        "created_at": "<<DATETIME>>"
    }
```
- Verificar se API foi atualizada rodando a API `GET: http://localhost:8000/tasks`
- A API retorna mensagem de sucesso:
```json
{
    "message": "Task atualizada com sucesso"
}
```
- Verificar no Studio3T se os dados foram atualizados do MongoDB. 
- Processo pode ser feito usando um JS `db.getCollection("tasks").find({})`.

#### 10. Deletando task

- Inserir em `main.go` a função para apagar a task.
- Inserir o router de apagar task.
- No Postman criar um requet `DELETE: http://localhost:8000/task/<<ID>>` para deletar a task.
- Verificar se a task foi apagada e se API retornou a mensagem de sucesso:
```json
{
    "message": "Task deletada com sucesso"
}   
```
- Verificar no Studio3T se os dados foram deletados do MongoDB. 
- Processo pode ser feito usando um JS `db.getCollection("tasks").find({})`.
- Após a finalização do CRUD popule  Tasks com os JS [populando-tasks.js](/go-crud-api/populando-tasks.js) ou [inserir-task.js](/go-crud-api/inserir-task.js). 
- Basta inserir os dados no MongoDB e checar com comando `db.getCollection("tasks").find({})` ou `db.tasks.find().pretty()`.

#### 11. Adicionando Swagger na aplicação

- Essa parte é para documentar a API é feita através da instalação de pacotes.
- Acesse o terminal na pasta do projeto `cd go-crud-api`.
- Instale os pacotes via terminal: 
    - `go get -u github.com/swaggo/gin-swagger`.
    - `go get -u github.com/swaggo/files`.
    - `go install github.com/swaggo/swag/cmd/swag@latest`.
- Fazer o import dos pacotes dentro do `main.go`.
```go
    "go-crud-api-docs" 
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
```
- Adiciona novo endpoint.
- Criar uma pasta `docs` em `go-crud-api/docs`.
- Inserir a documentação em cada função criada em `main.go`
	- `func createTask` inserir:
	```go
	// @Summary      Criando nova task
	// @Description  Cria uma nova task no banco de dados.
	// @Tags         tasks
	// @Accept       json
	// @Produce      json
	// @Param        task  body      models.Task  true  "Dados da Tarefa"
	// @Success      201   {object}  map[string]interface{}
	// @Failure      400   {object}  map[string]string
	// @Failure      500   {object}  map[string]string
	// @Router       /task [post]
	```
	
	- `func listTask` inserir:
	```go
	// @Summary      Listar tasks
	// @Description  Lista todas as tasks no banco de dados.
	// @Tags         tasks
	// @Accept       json
	// @Produce      json
	// @Success      200   {array}   models.Task
	// @Failure      500   {object}  map[string]string
	// @Router       /tasks [get]
	```
	
	- `func updateTask` inserir:
	```go
	// @Summary      Atualizando a task
	// @Description  Atualiza uma task no banco de dados.
	// @Tags         tasks
	// @Accept       json
	// @Produce      json
	// @Param        id    path      string      true  "ID da tarefa"
	// @Param        task  body      models.Task true  "Dados da Tarefa"
	// @Success      200   {object}  map[string]string
	// @Failure      400   {object}  map[string]string
	// @Failure      404   {object}  map[string]string
	// @Failure      500   {object}  map[string]string
	// @Router       /tasks/{id} [put]
	```
	- `func deleteTask` inserir:
	```go
	// @Summary      Apagar a task
	// @Description  Apaga uma task no banco de dados.
	// @Tags         tasks
	// @Accept       json
	// @Produce      json
	// @Param        id    path      string      true  "ID da tarefa"
	// @Success      200   {object}  map[string]string
	// @Failure      400   {object}  map[string]string
	// @Failure      404   {object}  map[string]string
	// @Failure      500   {object}  map[string]string
	// @Router       /tasks/{id}     [delete] {delete}
	```
	
- No terminal execute `swag --version` para checar a versão.
- Execute `swag init` para gerar todos os arquivos necessários a partir dos comentários da documentação.
- Verifique se foram criados os arquivos `docs.go`, `swagger.json`, `swagger.yaml` em `go-crud-api\docs`.
- Abrir no navegador o localhost `http://localhost:8000/swagger/index.html`.
- Agora temos a visualização dos Swaggers.

---

##### Contribuidores

<!-- Link to generate contributors: https://hub-io-mcells-projects.vercel.app/ --->
<table>
  <tbody>
    <tr><td align="left" valign="top" width="12.5%" style="word-break: break-word; white-space: normal;"><a href="https://github.com/mayannaoliveira" title="mayannaoliveira"><img src="https://avatars.githubusercontent.com/u/8138985?v=4" width="100px;" alt="mayannaoliveira" style="border-radius: 9999px;" /></a></td>
    </tr>
  </tbody>
</table>

_Caso queira usar o repositório leia [Contribuidores](/contributors.md) antes de fazer o fork._
 
