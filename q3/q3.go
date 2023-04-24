package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	// retorna um erro caso a lista esteja vazia
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}

	// inicia as variáveis maior, menor e soma
	maior := numbers[0]
	menor := numbers[0]
	soma := 0

	// loop para descobrir maior número, menor número e soma
	for _, numero := range numbers {
		if numero > maior {
			maior = numero
		}
		if numero < menor {
			menor = numero
		}
		soma += numero
	}

	// calcula a média desconsiderando o maior e o menor número
	media := float64(soma-menor-maior) / float64(len(numbers)-2)

	return maior, menor, media, nil
}
