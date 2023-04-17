package q4

import (
	"testing"
)

func TestCalculateFinalPrice(t *testing.T) {
	tests := []struct {
		name          string
		basePrice     float64
		state         string
		taxCode       int
		expectedPrice float64
		expectedError bool
	}{
		{
			name:          "base price is zero",
			basePrice:     0,
			state:         "SP",
			taxCode:       1,
			expectedPrice: 0,
			expectedError: true,
		},
		{
			name:          "tax code is invalid",
			basePrice:     100,
			state:         "SP",
			taxCode:       0,
			expectedPrice: 0,
			expectedError: true,
		},
		{
			name:          "tax code is 1 and state is SP",
			basePrice:     100,
			state:         "SP",
			taxCode:       1,
			expectedPrice: 115,
		},
		{
			name:          "tax code is 2 and state is SP",
			basePrice:     100,
			state:         "SP",
			taxCode:       2,
			expectedPrice: 120,
		},
		{
			name:          "tax code is 3 and state is SP",
			basePrice:     100,
			state:         "SP",
			taxCode:       3,
			expectedPrice: 125,
		},
		{
			name:          "tax code is 1 and state is RJ",
			basePrice:     100,
			state:         "RJ",
			taxCode:       1,
			expectedPrice: 120,
		},
		{
			name:          "tax code is 2 and state is RJ",
			basePrice:     100,
			state:         "RJ",
			taxCode:       2,
			expectedPrice: 125,
		},
		{
			name:          "tax code is 3 and state is RJ",
			basePrice:     100,
			state:         "RJ",
			taxCode:       3,
			expectedPrice: 130,
		},
		{
			name:          "tax code is 1 and state is MG",
			basePrice:     100,
			state:         "MG",
			taxCode:       1,
			expectedPrice: 125,
		},
		{
			name:          "tax code is 2 and state is MG",
			basePrice:     100,
			state:         "MG",
			taxCode:       2,
			expectedPrice: 130,
		},
		{
			name:          "tax code is 3 and state is MG",
			basePrice:     100,
			state:         "MG",
			taxCode:       3,
			expectedPrice: 135,
		},
		{
			name:          "tax code is 1 and state is ES",
			basePrice:     100,
			state:         "ES",
			taxCode:       1,
			expectedPrice: 130,
		},
		{
			name:          "tax code is 2 and state is ES",
			basePrice:     100,
			state:         "ES",
			taxCode:       2,
			expectedPrice: 135,
		},
		{
			name:          "tax code is 3 and state is ES",
			basePrice:     100,
			state:         "ES",
			taxCode:       3,
			expectedPrice: 140,
		},
		{
			name:          "tax code is 1 and state is other",
			basePrice:     100,
			state:         "other",
			taxCode:       1,
			expectedPrice: 135,
		},
		{
			name:          "tax code is 2 and state is other",
			basePrice:     100,
			state:         "other",
			taxCode:       2,
			expectedPrice: 140,
		},
		{
			name:          "tax code is 3 and state is other",
			basePrice:     100,
			state:         "other",
			taxCode:       3,
			expectedPrice: 145,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price, err := CalculateFinalPrice(tt.basePrice, tt.state, tt.taxCode)
			if tt.expectedError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if price != tt.expectedPrice {
				t.Errorf("expected price %v, got %v", tt.expectedPrice, price)
			}
		})
	}
}
