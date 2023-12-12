package main

import (
	"fmt"

	"conta_bancaria/clientes"
	"conta_bancaria/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clientBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"}

	cliente1 := contas.ContaCorrente{Titular: clientBruno,
		Agencia: 123, Conta: 123456}

	cliente1.Depositar(300)

	clientDaniel := clientes.Titular{
		Nome:      "Daniel",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"}

	cliente2 := contas.ContaCorrente{Titular: clientDaniel,
		Agencia: 123, Conta: 123456}

	cliente2.Depositar(200)

	fmt.Println("Saldo", cliente1.ObterSaldo())

	cliente1.Sacar(10)

	fmt.Println("Saldo", cliente1.ObterSaldo())

	fmt.Println("Saldo", cliente1.Sacar(400))

	cliente1.Depositar(1000)

	fmt.Println("Saldo", cliente1.ObterSaldo())

	fmt.Println(cliente1.Depositar(-100))

	fmt.Println("Saldo", cliente1.ObterSaldo())
	fmt.Println(cliente1.Transferir(100, &cliente2))
	fmt.Println("Saldo", cliente1.ObterSaldo())
	fmt.Println("Saldo", cliente2.ObterSaldo())
	fmt.Println(cliente1.Transferir(-100, &cliente2))

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())
}
