package q1

import (
	"testing"
)

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		name             string
		currentPurchase  float64
		purchaseHistory  []float64
		expectedDiscount float64
		expectedError    bool
	}{
		{
			name:             "Compra atual de 500 e histórico de compras acima de 1000",
			currentPurchase:  500.0,
			purchaseHistory:  []float64{1000, 800, 800, 900},
			expectedDiscount: 500 * 0.1,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 200 e histórico de compras entre 500 e 1000",
			currentPurchase:  200.0,
			purchaseHistory:  []float64{300, 400, 150, 150},
			expectedDiscount: 200 * 0.05,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 100 e histórico de compras abaixo de 500",
			currentPurchase:  100.0,
			purchaseHistory:  []float64{50, 75, 80, 90},
			expectedDiscount: 100 * 0.02,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 1000 e histórico de compras vazio",
			currentPurchase:  1000.0,
			purchaseHistory:  []float64{},
			expectedDiscount: 1000 * 0.1,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 700 e histórico de compras acima de 1000, média de compras acima de 1000",
			currentPurchase:  700.0,
			purchaseHistory:  []float64{1500, 800, 1000},
			expectedDiscount: 140,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 1500 e histórico de compras acima de 1000, média de compras acima de 1000",
			currentPurchase:  1500.0,
			purchaseHistory:  []float64{800, 1200, 700, 1500},
			expectedDiscount: 300,
			expectedError:    false,
		},
		{
			name:             "Compra atual de 0",
			currentPurchase:  0.0,
			purchaseHistory:  []float64{500, 800, 1000},
			expectedDiscount: 0.0,
			expectedError:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			discount, err := CalculateDiscount(test.currentPurchase, test.purchaseHistory)

			if discount != test.expectedDiscount {
				t.Errorf("Erro: valor do desconto esperado %f, valor do desconto retornado %f", test.expectedDiscount, discount)
			}

			if test.expectedError && err == nil {
				t.Errorf("Erro: esperava um erro, mas não houve erro")
			}
		})
	}
}
