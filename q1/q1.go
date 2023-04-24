package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("Valor da compra inválido")
	}
	var Desconto float64
	var total float64
	for _, purchase := range purchaseHistory {
		total += purchase
	}
	if len(purchaseHistory) == 0 {
		Desconto = currentPurchase * 0.1
	} else {
		média := total / float64(len(purchaseHistory))
		if média > 1000 {
			Desconto = currentPurchase * 0.2
		} else if total > 1000 {
			Desconto = currentPurchase * 0.1
		} else if total > 500 {
			Desconto = currentPurchase * 0.05
		} else {
			Desconto = currentPurchase * 0.02
		}
	}

	return Desconto, nil
}
