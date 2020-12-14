## CRUD MVC em Linguagem Go com MySql

**Alguns recursos**
* Validação campos vazios e nulos no POST e PUT
* Validação email
* Retornando mensagem de erro json e status code
* Salvando senha com hash no banco de dados
* login autenticando por token


**Comando pra executar no terminal dento da pasta API:** `go run main.go`

[![](http://img.youtube.com/vi/I1ItsyI4P9A/0.jpg)](http://www.youtube.com/watch?v=I1ItsyI4P9A "Video de como criar usuarios e logar na api GO-MVC")


**HOST:PORT** `localhost:5000/`

---

**Exemplo de POST:**

URI: localhost:5000/usuarios
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
  "senha": "$2a$10$Qq1A4IOZ69EeWyN4ifZGRutYWTRz041zJfPrUohjD3Se2BMEcPSaC",
  "CriadoEm": "0001-01-01T00:00:00Z"
}

```

---

**login**

URI: localhost:5000/login

``` json
{
	"email":"leandro@user.com",
	"senha":"123"
}

```

**Retorno esperado** 

status: 200 OK

`Bearer Token ...` 

---

**Atualizar senha**

URI: localhost:5000/usuarios/`usuario`/atualizar-senha

``` json
{
  "atual":"123",
  "nova":"456"
}

```
