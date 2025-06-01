

## httpYac - Rest Client

É uma extensão para o VSCodium chamada de [httpYac - Rest Client](https://httpyac.github.io/guide/request.html) que permite a criação e manipulação de API.

### Instalação
A instalação é feita pela API.
- VSCodium > Extensions > [httpYac - Rest Client](https://marketplace.visualstudio.com/items?itemName=anweber.vscode-httpyac).
- Instale a extensão.

### Utilizando o httpYac no Projeto

**Requisição GET:**

- Para verificar localhost: `GET http://localhost:8000/ HTTP/1.1`
> Caso sucesso deve retornar mensagem de `Olá Mundo!`.

- Para listar todas as tasks: `GET http://localhost:8000/tasks HTTP/1.1`
> Caso sucesso deve retornar todas as tasks criadas.

**Requisição POST:**

POST http://localhost:8000/tasks
Content-Type: application/json
```json
{
    "title": "Bela e a Fera",
    "description": "Musical da Disney",
    "completed": true,
    "created_at": "2022-02-15T19:50:00Z",
    "due_date": "2022-03-01T00:00:00Z"
}
```
> Caso sucesso deve retornar O ID da task `"id": "683c867105786c610ce02ba7"`.

**Requisição PUT:**

PUT http://localhost:8000/tasks/683c791192fc3e05efe1869d
Content-Type: application/json
```json
    {
        "id": "683c791192fc3e05efe1869d",
        "title": "Atualizando a task",
        "description": "Verifique se a task atualizou",
        "completed": true,
        "due_date": "0001-01-01T00:00:00Z",
        "created_at": "2025-06-01T16:00:17.773Z"
    }
```
> Caso sucesso deve retornar mensagem `{"message": "Task atualizada com sucesso"}`

**Requisição DELETE:**

 DELETE http://localhost:8000/tasks/683c791192fc3e05efe1869d ```
> Caso sucesso deve retornar mensagem `{"message": "Task deletada com sucesso"}`

