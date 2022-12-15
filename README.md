# Client-Server-API

Código de resolução do desafio client-server-api proposto no curso [GoExpert](https://goexpert.fullcycle.com.br/)

## Desafio

Olá dev, tudo bem?

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.

Você precisará nos entregar dois sistemas em Go:

- `client.go`
- `server.go`

Os requisitos para cumprir este desafio são:

O `client.go` deverá realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.

O `server.go` deverá consumir a API contendo o câmbio de Dólar e Real no endereço: [https://economia.awesomeapi.com.br/json/last/USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL) e em seguida deverá retornar no formato JSON o resultado para o cliente.

Usando o package "context", o `server.go` deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

O `client.go` precisará receber do `server.go` apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o `client.go` terá um timeout máximo de 300ms para receber o resultado do `server.go`.

O `client.go` terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

O endpoint necessário gerado pelo `server.go` para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

Ao finalizar, envie o link do repositório para correção.

## Solução

### Client

O client é um programa que faz uma requisição HTTP para o servidor e salva o valor da cotação em um arquivo.

Para executar o client, basta executar o comando abaixo:

```bash
go run cmd/client/client.go
```

É possível passar o nome do arquivo que será salvo o valor da cotação como argumento. Caso não seja passado nenhum argumento, o valor da cotação será salvo no arquivo `cotacao.txt`.

```bash
go run cmd/client/client.go --filename=meu-arquivo.txt
```

Também é possível passar o endereço do servidor como argumento. Caso não seja passado nenhum argumento, o endereço do servidor será `http://localhost:8080`.

```bash
go run cmd/client/client.go --addr=http://localhost:8080
```

### Server

O server é um programa que faz uma requisição HTTP para a API de cotação do dólar e salva o valor da cotação no banco de dados.

Para executar o server, basta executar o comando abaixo:

```bash
go run cmd/server/server.go
```

É possível passar o endereço em que o servidor HTTP será executado como argumento. Caso não seja passado nenhum argumento, o endereço do servidor será `:8080`.

```bash
go run cmd/server/server.go --addr=:8080
```

Tamém é possível passar a string de conexão com o banco de dados como argumento. Caso não seja passado nenhum argumento, a string de conexão será `file:db.sqlite3?cache=shared`.

```bash
go run cmd/server/server.go --dsn="file:db.sqlite3?cache=shared"
```
