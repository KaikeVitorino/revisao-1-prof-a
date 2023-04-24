package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var Resultado float64
	var erro error

	//Funcao Switch que meio q substitui o If/Else
	switch fromScale {
	case "C":
		switch toScale {
		case "F":
			Resultado = temp*9/5 + 32
		case "K":
			Resultado = temp + 273.15
		default:
			erro = fmt.Errorf("escala inválida")
		}
	case "F":
		switch toScale {
		case "C":
			Resultado = (temp - 32) * 5 / 9
		case "K":
			Resultado = (temp-32)*5/9 + 273.15
		default:
			erro = fmt.Errorf("escala inválida")
		}
	case "K":
		switch toScale {
		case "C":
			Resultado = temp - 273.15
		case "F":
			Resultado = (temp-273.15)*9/5 + 32
		default:
			erro = fmt.Errorf("escala inválida")
		}
	default:
		erro = fmt.Errorf("escala inválida")
	}

	return Resultado, erro
}
