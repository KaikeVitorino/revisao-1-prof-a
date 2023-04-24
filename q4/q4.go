package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	//Tabela dos impostos
	//(Funcao MAP nova q eu descobri no chatgpt)
	//Basicamente vc cria uma variavel string com chaves q se completam com valores pre-definidos ou nn
	//Exemplo: imposto["2"] valor da chave 2, ou seja 0.1 nesse caso
	//tem q aprender a usar direitinho, mas ja to em um bom inicio
	var Imposto = map[int]float64{
		1: 0.05,
		2: 0.1,
		3: 0.15,
	}

	//Tabelando o frete
	//Usando a funcao MAP dnv
	//Exemplo: percentuaisFrete["SP"], valor da chave "SP" no mapa, que no caso eh 0.1
	//Criei as variaveis em float em vez de inteiros para facilitar no calculo dps
	//10% = 0.1 / 15% = 0.15 / 20% = 0.2 e por ai vai indo
	var percentuaisFrete = map[string]float64{
		"SP":     0.1,
		"RJ":     0.15,
		"MG":     0.2,
		"ES":     0.25,
		"OUTROS": 0.3,
	}

	//retornar erro se for um preco invalido
	if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}

	//Retornar erro se o codigo do imposto nn existir
	imposto, existeImposto := Imposto[taxCode]
	if !existeImposto {
		return 0, fmt.Errorf("imposto não encontrado")
	}

	//ve qual o estado e o frete
	percentualFrete, existeFrete := percentuaisFrete[state]
	if !existeFrete {
		percentualFrete = percentuaisFrete["OUTROS"]
	}

	//calculo do preco final
	precoImposto := basePrice * imposto
	precoFrete := basePrice * percentualFrete
	precoFinal := basePrice + precoImposto + precoFrete

	return precoFinal, nil
}
