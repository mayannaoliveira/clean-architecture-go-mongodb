

[![Udemy](https://img.shields.io/badge/Udemy-A435F0?style=flat&logo=udemy&logoColor=white&link=https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/)](https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/)
[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=Go&logoColor=white&link=https://go.dev/)](https://go.dev/)
[![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=flat&logo=mongodb&logoColor=white&link=https://www.mongodb.com/)](https://www.mongodb.com/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white&link=https://www.docker.com/)](https://www.docker.com/)
[![Gin Framework](https://img.shields.io/badge/Gin%20Framework-008ECF?style=flat&logo=gin&logoColor=white&link=https://gin-gonic.com/)](https://gin-gonic.com/)
[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/)
[![GitHub](https://img.shields.io/badge/GitHub-181717?style=flat&logo=github&logoColor=white&link=https://github.com/mayannaoliveira)](https://github.com/mayannaoliveira)
[![Studio3T](https://img.shields.io/badge/Studio%203T-17AF66?style=flat&logo=Studio3T&logoColor=white&link=https://studio3t.com/)](https://studio3t.com/)
[![VSCodium](https://img.shields.io/badge/VSCodium-2F80ED?style=flat&logo=vscodium&logoColor=white&link=https://vscodium.com/)](https://vscodium.com/)

## Clean Architecture na prática: Go, MongoDB e Docker

Projetinho feito durante o curso [Clean Architecture na prática: Go, MongoDB e Docker
](https://www.udemy.com/course/clean-architecture-na-pratica-com-golang/) na [Udemy](https://www.udemy.com/). O intuito desse repositório é de estudar Go e MongoDB.

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

- Instalar a extenção da go.dev [Go - Go Team at Google](https://marketplace.visualstudio.com/items?itemName=golang.go).
- Adicionar em `Arquivo > Preferências > Configurações` as configuração no VS Code:
```json
{
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.useLanguageServer": true
}
```
- Sempre uso a extensão [Docker da Microsoft](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker) para facilitar a visualização dos conteiners.

![Go Team at Google](https://img.shields.io/badge/Go%20Team%20at%20Google-2F80ED?style=flat&logo=vscodium&logoColor=white)
[![Docker Microsoft](https://img.shields.io/badge/Docker%20Microsoft-2F80ED?style=flat&logo=vscodium&logoColor=white&link=https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)

### 3. Docker
- Instalação do Docker.
Checar Docker version com comando `docker --version`.

[![Docker Docs](https://img.shields.io/badge/Docker%20Docs-2496ED?style=flat&logo=Docker&logoColor=white&link=https://docs.docker.com/?_gl=1*bu2mvx*_gcl_au*NTA5ODMyMTIzLjE3NDgxMTAyNTE.*_ga*Mjk3NzcxNTE5LjE3NDgxMTAyNTE.*_ga_XJWPQMJYHQ*czE3NDg1NzA3ODUkbzIkZzEkdDE3NDg1NzA3OTAkajU1JGwwJGgw)](https://docs.docker.com/?_gl=1*bu2mvx*_gcl_au*NTA5ODMyMTIzLjE3NDgxMTAyNTE.*_ga*Mjk3NzcxNTE5LjE3NDgxMTAyNTE.*_ga_XJWPQMJYHQ*czE3NDg1NzA3ODUkbzIkZzEkdDE3NDg1NzA3OTAkajU1JGwwJGgw)

> A execução do projeto depende da criação do conteiner e que ele esteja executando.

### Seção 2: Desenvolvimento de API REST em Go

#### 4. Introdução ao módulo: 
- Teste executado com old_main.go com o comando `go run old_main.go`.
- Verificar se no localhost `http://localhost:8000/` é possível visualizar a mensagem `Olá Mundo!`.

#### 5. Criando API: 
- Crie um novo arquivo `main.go` e adicione o código e execute no terminal `go run main`.
- Criar um servidor HTTP com a rota `/` que retorna a mensagem `Olá Mundo!`:
- Executea o comando `go run main.go`.
- Verifique se no localhost `http://localhost:8000/` é possível visualizar a mensagem `Olá Mundo!`.
- Caso erro devido ao erro de porta, altere o comando para `go run -p 8001 main.go` e execute novamente.
- Caso o Framework Gin não esteja instalado use o comando `go get -u github.com/gin-gonic/gin`.
- Execute o comando `go run main.go` para verficar `Olá,Mundo!` na página localhost.

[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/)
[![Postman](https://img.shields.io/badge/API%20do%20Curso%20Postman-FF6C37?style=flat&logo=postman&logoColor=white&link=https://www.postman.com/)](https://www.postman.com/)

> As resquisições e o ambiente de teste foram salvos na pasta [Postman](\postman) para utilizar faça o `import` no Postman.

#### 6. Adicionando o banco de dados:

- Verifique se o Docker está em execução.
- Crie um conteiner Docker com o comando `docker run -d -p 27017:27017 --name mongodb mongo:4`.
- Verifique com `docker ps` se a instalação foi realizada.
- Baixar pacotes para o projeto: `go mod tidy`.
- Continuar a parte do código para criar uma task.

#### 7. Criando uma Task (Create do CRUD)
 - Criar novo repositório `models` e adicionar o arquivo `task.go`.
 - Adicionar uma struct `Task` dentro do arquivo `task.go`.
 - Adicionar um método `CreateTask` dentro do arquivo `main.go`.
 - Inserir tratamento de exceção para verificar a conexão com o banco de dados.
 - Criar o router para POST `Createtask`.
- Instalar a ferramenta [Studio 3T](https://studio3t.com/) para visualizar as conexões do MongoDB.

[![Studio 3T](https://img.shields.io/badge/Studio%203T-17AF66?style=flat&logo=Studio3T&logoColor=white&link=https://studio3t.com/)](https://studio3t.com/)

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

#### 8. Listando as tasks (Read do CRUD):
- Criar o bloco de `listTask` dentro do arquivo `main.go`.
- Inserir tratamento de exceção para verificar a conexão com o banco de dados.
- Inserir o router para `listTask`. 
- No Postman teste utilizando uma requisição `GET: http://localhost:8000/tasks` para listar todas as tasks.
- Verificar se retorna a lista de tasks adicionadas.


#### 9. Atualizando as tasks (Update do CRUD):
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
 
#### 10. Deletando task (Delete do CRUD):
 
#### 11. Adicionando Swagger na aplicação:
 
#### 12. Validando CRUD com Frontend:


### Seção 1: Introdução ao curso e configuração de ambiente

#### 13. Introdução a Clean Architecture
 
#### 14. Conhecendo a Clean Architecture

#### 15. Refatorando o projeto para Clean Architecture
 
#### 16. Entidades

#### 17. Criando a repository
 
#### 18. Repository
 
#### 19. Criando a useCase

#### 20. UseCase

#### 21. Criando a handler

#### 22. Handler

#### 23. Configurando IoC

#### 24. Testando API - RestClient

#### 25. Validando projeto com Frontend


---

##### Contribuidores:

<!-- Link to generate contributors: https://hub-io-mcells-projects.vercel.app/ --->
<table>
  <tbody>
    <tr><td align="left" valign="top" width="12.5%" style="word-break: break-word; white-space: normal;"><a href="https://github.com/mayannaoliveira" title="mayannaoliveira"><img src="https://avatars.githubusercontent.com/u/8138985?v=4" width="100px;" alt="mayannaoliveira" style="border-radius: 9999px;" /></a></td>
    </tr>
  </tbody>
</table>

_Caso queira usar o repositório leia [Contribuidores](/contributors.md) antes de fazer o fork. Repositório criado somente para estudar Go e MongoDB_



