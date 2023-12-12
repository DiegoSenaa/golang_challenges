package contas

import "conta_bancaria/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	}

	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado", c.saldo
	}

	return "Valor incorreto", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
