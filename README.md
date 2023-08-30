# Quotation Service (Serviço de cotação)

Esse repositório faz parte de um desafio do curso [Go Expert](https://goexpert.fullcycle.com.br/curso).

## Descrição do desafio

Você precisará nos entregar dois sistemas em Go:
- client
- server
 
Os requisitos para cumprir este desafio são:
 
O client deverá realizar uma requisição HTTP no server solicitando a cotação do dólar.
 
O server deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client terá que salvar a cotação atual em um arquivo `cotacao.txt` no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server para este desafio será: `/cotacao` e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Como rodar o projeto

### Rodar o server

- Entrar na pasta `/server`
- Rodar o comando `go run main.go`

### Rodar o client

- Entrar na pasta `/client`
- Rodar o comando `go run main.go`
