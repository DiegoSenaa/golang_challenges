# Objetivo

Basicamente é ler a partir de um json aléatorio que achei na net e dá um get nesses sites, pegar o retorno, se demorar 10 segundos, a requisição é morta, qualquer status diferente de 200 é erro. Foi usado goroutine, porque é 500 sites e eu queria aprender goroutine, basicamente por isso (rs).

## Como rodar

```
Pra buildar e depois rodar:

$ go build main.go
$./main

Ou se tu for preguiçoso faz assim:

$ go run main.go
```