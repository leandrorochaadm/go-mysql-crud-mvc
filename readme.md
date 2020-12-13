## CRUD MVC em Linguagem Go com MySql

**Alguns recursos**
* Validação campos vazios e nulos no POST e PUT
* Validaçao email
* Retornando mensagem de erro json e status code
* Salvando senha com hash no banco de dados


**Comando pra executar no terminal dento da pasta API:** `go run main.go`

**URI base:** `localhost:5000/usuarios`


**Exemplo de POST:**

``` json
{
	"nome":"Leandro",
	"email":"leandro@user.com",
	"senha":"123"
}
```

**Retorno esperado** 

status: 200 OK
```json
{
  "id": 1,
  "nome": "Leandro",
  "email": "leandro@user.com",
  "senha": "123",
  "CriadoEm": "0001-01-01T00:00:00Z"
}

```
