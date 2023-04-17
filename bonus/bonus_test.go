package bonus

import (
	"fmt"
	"testing"
)

func TestCalculateDamage(t *testing.T) {
	tests := []struct {
		characterLevel int
		enemyLevel     int
		expectedDamage int
		expectedError  bool
	}{
		{characterLevel: 10, enemyLevel: 5, expectedDamage: 100},
		{characterLevel: 5, enemyLevel: 10, expectedDamage: 25},
		{characterLevel: 10, enemyLevel: 10, expectedDamage: 70},
		{characterLevel: 150, enemyLevel: 50, expectedDamage: 15000},
		{characterLevel: 60, enemyLevel: 50, expectedDamage: 600},
		{characterLevel: 50, enemyLevel: 60, expectedDamage: 250},
		{characterLevel: 50, enemyLevel: 150, expectedDamage: 200},
		{characterLevel: 30, enemyLevel: 50, expectedDamage: 150},
		{characterLevel: 50, enemyLevel: 30, expectedDamage: 500},
		{characterLevel: 0, enemyLevel: 50, expectedDamage: 0, expectedError: true},
		{characterLevel: 50, enemyLevel: 0, expectedDamage: 0, expectedError: true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("characterLevel=%d, enemyLevel=%d", test.characterLevel, test.enemyLevel), func(t *testing.T) {
			damage, err := CalculateDamage(test.characterLevel, test.enemyLevel)
			if test.expectedError && err == nil {
				t.Errorf("Erro esperado, mas nenhum erro foi retornado")
			}

			if damage != test.expectedDamage {
				t.Errorf("Dano esperado: %d, Dano obtido: %d", test.expectedDamage, damage)
			}
		})
	}
}
