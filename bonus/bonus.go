package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	// Implemente sua solução aqui
	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("nível inválido")
	}

	//Clacula dano
	var dano int
	if characterLevel > 100 {
		dano = characterLevel * 2
	} else if enemyLevel > 100 {
		dano = characterLevel * 20
	} else {
		levelDiff := characterLevel - enemyLevel
		if levelDiff > 20 {
			dano = characterLevel * 5
		} else if levelDiff < -20 {
			dano = characterLevel / 2
		} else {
			switch {
			case levelDiff > 0:
				dano = characterLevel * 10
			case levelDiff < 0:
				dano = characterLevel * 5
			default:
				dano = characterLevel * 7
			}
		}
	}

	return dano, nil
}
