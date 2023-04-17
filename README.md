# Revisão

## Instruções

1. Resolva as questões em seus respectivos arquivos. Ex: ``q1/q1.go``
2. Utilize o arquivo `main.go` para testar suas funções.
3. Ao finalizar, não esqueça de fazer o commit e o push.

## Questão 1

Uma loja de roupas precisa de uma função para verificar se um cliente é elegível para um desconto com base em seu
histórico de compras. A função deve receber o valor da compra atual e o histórico de compras do cliente.
A função deve retornar o valor do desconto e um erro. O erro deve ser nulo (ou seja, nil) se não houver nenhum erro.
Caso o valor da compra atual seja menor ou igual a zero, a função deve retornar um erro com a mensagem "valor da compra
inválido".

Regras de desconto:

* Se o valor total de compras do cliente for maior que R$ 1000,00, o desconto é de 10% do valor da compra atual.
* Se o valor total de compras do cliente for menor ou igual a R$ 1000,00, o desconto é de 5% do valor da compra atual.
* Se o valor total de compras do cliente for menor ou igual a R$ 500,00, o desconto é de 2% do valor da compra atual.
* Se a compra atual for a primeira do cliente, o desconto é de 10% do valor da compra atual.
* Se a média de compras do cliente for maior que R$ 1000,00, o desconto é de 20% do valor da compra atual.

Os descontos não são cumulativos. Ou seja, se o cliente tiver direito a mais de um desconto, deve ser aplicado o maior
desconto.

### Exemplo de entrada

```
Valor da compra atual: R$ 1500,00
Histórico de compras do cliente: [R$ 800,00, R$ 1200,00, R$ 700,00, R$ 1500,00]
```

### Exemplo de saída:

```
Valor do desconto: R$ 300,00
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

O valor da compra atual é de R$ 1500,00 e o cliente possui um histórico de compras com os seguintes valores: R$ 800,00,
R$ 1200,00, R$ 700,00 e R$ 1500,00.

A soma do histórico de compras do cliente é de R$ 4200,00, o que faz com que o cliente seja elegível para um desconto de
20% do valor da compra atual, de acordo com a regra de desconto de média de compras acima de R$ 1000,00.

Portanto, o valor do desconto é de R$ 300,00 (20% de R$ 1500,00) e não há erros, já que todas as condições foram
atendidas.

## Questão 2

Um aplicativo de processamento de texto precisa de uma função para contar a média de letras por
palavra em um texto. A função deve receber um texto e retornar a média de letras por palavra.
A função deve retornar um erro caso o texto não possua nenhuma palavra. O erro deve conter a mensagem "texto
vazio".

### Exemplo de entrada

```
Texto: "O rato roeu a roupa do rei de Roma"
```

### Exemplo de saída:

```
Média de letras por palavra: 3.5
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

O texto possui 9 palavras, sendo que a maior palavra possui 4 letras e a menor possui 2 letras. A média de letras por
palavra é de 2.88 (26 letras / 9 palavras).

### Exemplo de entrada

```
Texto: ""
```

### Exemplo de saída:

```
Média de letras por palavra: 0
Erro: "texto vazio"
```

### Explicação:

O texto não possui nenhuma palavra, portanto a função deve retornar um erro com a mensagem "texto vazio".

### Exemplo de entrada

```
Texto: " "
```

### Exemplo de saída:

```
Média de letras por palavra: 0
Erro: "texto vazio"
```

### Explicação:

O texto possui apenas um espaço, portanto a função deve retornar um erro com a mensagem "texto vazio".

## Questão 3

Você está desenvolvendo um programa que recebe uma lista de números inteiros como entrada e precisa encontrar o maior
valor e o menor valor da lista. Além disso, você precisa calcular a média dos valores da lista, desconsiderando o maior
e o menor valor na média.

Escreva uma função que recebe uma lista de números inteiros como parâmetro e
retorna uma tupla contendo o maior valor, o menor valor e a média dos valores da lista, desconsiderando o maior e o
menor valor.

A função deve retornar um erro caso a lista esteja vazia. O erro deve conter a mensagem "lista vazia".

### Exemplo de entrada

```
Lista de números: [1, 2, 3, 4, 5]
```

### Exemplo de saída:

```
Maior valor: 5
Menor valor: 1
Média: 3
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

O maior valor da lista é 5, o menor valor é 1 e a média é 3 (2 + 3 + 4 = 9 / 3 = 3).

## Questão 4

Uma empresa de comércio eletrônico precisa de uma função para calcular o preço final de um produto,
que inclui o preço base, impostos e frete. A função deve receber o preço base, o estado de destino
e o código do imposto.

A alíquota do imposto e o percentual de frete se baseiam nas tabelas abaixo:

| Código do imposto | Alíquota |
|-------------------|----------|
| 1                 | 5%       |
| 2                 | 10%      |
| 3                 | 15%      |

| Estado de destino | Percentual de frete |
|-------------------|---------------------|
| SP                | 10%                 |
| RJ                | 15%                 |
| MG                | 20%                 |
| ES                | 25%                 |
| OUTROS            | 30%                 |

O preço final é calculado da seguinte forma:

``preço final = preço base + preço bas * alíquota do imposto + preço base * percentual de frete.``

A função deve retornar o preço final e um erro. O erro deve ser nulo (ou seja, nil) se não houver nenhum erro. Caso o
código do imposto não exista, a função deve retornar um erro com a mensagem "imposto não
encontrado". Caso o preço base seja menor ou igual a zero, a função deve retornar um
erro com a mensagem "preço base inválido".

### Exemplo de entrada

```
Preço base: 100
Estado de destino: SP
Código do imposto: 1
```

### Exemplo de saída:

```
Preço final: 115
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

O preço final é calculado da seguinte forma:

``preço final = 100 + 100 * 5% + 100 * 10% = 100 + 5 + 10 = 115``

## Questão 5

Uma aplicação de clima precisa de uma função para converter a temperatura entre diferentes escalas: Celsius, Fahrenheit
e Kelvin. A função deve receber a temperatura, a escala de origem e a escala de destino.

As escalas de temperatura são representadas pelas seguintes letras:

* C: Celsius
* F: Fahrenheit
* K: Kelvin

A função deve retornar a temperatura convertida na escala de destino. As fórmulas de conversão são:

* Celsius para Fahrenheit: `F = C * 9/5 + 32`
* Celsius para Kelvin: `K = C + 273.15`
* Fahrenheit para Celsius: `C = (F - 32) * 5/9`
* Fahrenheit para Kelvin: `K = (F - 32) * 5/9 + 273.15`
* Kelvin para Celsius: `C = K - 273.15`
* Kelvin para Fahrenheit: `F = (K - 273.15) * 9/5 + 32`

A função deve retornar um erro caso a escala de origem ou a escala de destino não sejam válidas, com a mensagem "escala
inválida".

### Exemplo de entrada

```
Temperatura: 0
Escala de origem: C
Escala de destino: F
```

### Exemplo de saída:

```
Temperatura convertida: 32
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

A temperatura de 0 graus Celsius é igual a 32 graus Fahrenheit.

## Questão Bônus

Um aplicativo de jogos precisa de uma função para calcular o dano causado por um personagem em um inimigo. A função
deve receber o nível do personagem e o nível do inimigo e retornar o dano causado. A função deve retornar um erro
caso o nível do personagem ou do inimigo seja inválido. O erro deve conter a mensagem "nível inválido".

O dano causado é calculado da seguinte forma:

* Se o nível do personagem for maior que o nível do inimigo, o dano causado é igual ao nível do personagem multiplicado
  por 10.
* Se o nível do personagem for menor que o nível do inimigo, o dano causado é igual ao nível do personagem multiplicado
  por 5.
* Se o nível do personagem for igual ao nível do inimigo, o dano causado é igual ao nível do
  personagem multiplicado por 7.
* Se o nível do personagem for maior que 100, o dano causado é igual ao nível do personagem multiplicado por 20.
* Se o nível do inimigo for maior que 100, o dano causado é igual ao nível do personagem multiplicado por 2.
* Se a diferença entre o nível do personagem e o nível do inimigo for maior que 20, o dano causado é multiplicado por 5.
* Se a diferença entre o nível do personagem e o nível do inimigo for menor que 20, o dano causado é por 2.

### Exemplo de entrada

```
Nível do personagem: 50
Nível do inimigo: 40
```

### Exemplo de saída:

```
Dano causado: 350
Erro: nil (ou seja, nenhum erro)
```

### Explicação:

O nível do personagem é maior que o nível do inimigo, mas a diferença entre eles é menor que 20. O dano causado é
calculado da seguinte forma:

``dano causado = 50 * 3 = 150``
