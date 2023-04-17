package q3

import (
	"github.com/revisao-1/utils"
	"testing"
)

func TestFindMinMaxAverage(t *testing.T) {
	tests := []struct {
		name            string
		numbers         []int
		expectedMax     int
		expectedMin     int
		expectedAverage float64
		expectedError   bool
	}{
		{
			name:            "Lista de números com valores positivos",
			numbers:         []int{1, 2, 3, 4, 5},
			expectedMax:     5,
			expectedMin:     1,
			expectedAverage: 3.0,
		},
		{
			name:            "Lista de números com valores negativos",
			numbers:         []int{-5, -2, -9, -1, -7},
			expectedMax:     -1,
			expectedMin:     -9,
			expectedAverage: -4.66,
		},
		{
			name:            "Lista de números com valores positivos e negativos",
			numbers:         []int{10, -2, 5, -8, 3},
			expectedMax:     10,
			expectedMin:     -8,
			expectedAverage: 2,
		},
		{
			name:            "Lista de números com um único valor",
			numbers:         []int{100},
			expectedMax:     100,
			expectedMin:     100,
			expectedAverage: 100.0,
		},
		{
			name:            "Lista de números vazia",
			numbers:         []int{},
			expectedMax:     0,
			expectedMin:     0,
			expectedAverage: 0.0,
			expectedError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			max, min, average, err := FindMinMaxAverage(test.numbers)

			if max != test.expectedMax {
				t.Errorf("Erro: valor máximo esperado %d, valor máximo retornado %d", test.expectedMax, max)
			}

			if min != test.expectedMin {
				t.Errorf("Erro: valor mínimo esperado %d, valor mínimo retornado %d", test.expectedMin, min)
			}

			if !utils.AssertFloatWithPrecision(test.expectedAverage, average, 1e-2) {
				t.Errorf("Erro: valor médio esperado %f, valor médio retornado %f", test.expectedAverage, average)
			}

			if test.expectedError && err == nil {
				t.Errorf("Erro: esperava um erro, mas não houve erro")
			}
		})
	}
}
