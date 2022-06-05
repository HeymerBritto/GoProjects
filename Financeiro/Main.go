package main

import (
	"fmt"
	"math"
)

func main() {

	investimentoMensal := 1000.00
	investimentoMensalCorrecao := investimentoMensal
	jurosIPCA := 0.05
	jurosAnuais := 0.05
	//jurosCincoAnos := 10 / 100

	periodo := 20 * 12
	anoAtual := 2022

	saltoAtualDepositado := 0.00
	saltoAtualDepositadoCorrecao := 0.00
	saldoAtualDepoisDosJuros := 0.00
	jurosAtual := 0.00

	for i := 1; i <= periodo; i++ {

		var periodoDozeMeses = i % 12
		if periodoDozeMeses == 0 {

			fmt.Println("Ano  - Depositado                   - Corrigido                    - Saldo + Juros")
			anoAtual += 1
			investimentoMensalCorrecao += (investimentoMensalCorrecao * jurosIPCA)
		}

		saltoAtualDepositado += investimentoMensal
		saltoAtualDepositadoCorrecao += investimentoMensalCorrecao

		jurosAtual = (saldoAtualDepoisDosJuros + investimentoMensalCorrecao) * (math.Pow(jurosAnuais, (1.00/12.00)) / 100)

		saldoAtualDepoisDosJuros = (saldoAtualDepoisDosJuros + investimentoMensalCorrecao) + jurosAtual

		fmt.Printf("%d - %f (+%f) - %f (+%f) - %f - %f\n", anoAtual, saltoAtualDepositado, investimentoMensal, saltoAtualDepositadoCorrecao, investimentoMensalCorrecao, saldoAtualDepoisDosJuros, jurosAtual)

	}

}
