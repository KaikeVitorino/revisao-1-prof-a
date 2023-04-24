package q2

import (
	"fmt"
	"regexp"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	//Tirando os caracteres especiais com REGEXP, nn sabia q existia, descobri pelo chatgpt
	text = regexp.MustCompile("[^a-zA-Z0-9 ]+").ReplaceAllString(text, "")

	// Remove numeros, eu tava fazendo um por um ai descobri q da para fazer tudo de uma vez assim
	text = regexp.MustCompile("[0-9]+").ReplaceAllString(text, "")

	// Remove espacos em branco, Nova funcao q eu vi no chatgpt
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	// Calcula a media de letras por palavra
	Palavras := strings.Fields(text)
	numPalavras := len(Palavras)
	numLetras := 0
	for _, PalavraLetra := range Palavras {
		numLetras += len(PalavraLetra)
	}
	Media := float64(numLetras) / float64(numPalavras)

	// Verifica se tem palavras no texto
	if numPalavras == 0 {
		return 0, fmt.Errorf("texto vazio")
	}

	return Media, nil
}
