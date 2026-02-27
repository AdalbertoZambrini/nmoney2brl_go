CLI Currency Converter em Go 🪙
Um conversor de moedas simples via linha de comando desenvolvido em Go. O programa recebe um valor em Reais (BRL) e o converte para a moeda de destino escolhida utilizando taxas de câmbio pré-definidas.

🚀 Funcionalidades
Conversão rápida e direta direto no terminal.

Suporte a diversas moedas (USD, EUR, JPY, GBP, etc.).

Tratamento de erros para entradas inválidas ou moedas não suportadas.

Case-insensitive: funciona independentemente de você digitar a moeda em maiúsculas ou minúsculas (ex: usd ou USD).

💻 Como executar
Certifique-se de ter o Go instalado na sua máquina.

Clone este repositório ou baixe o arquivo main.go.

Abra o terminal na pasta do projeto e compile o código:

Bash

go build -o convert main.go
Execute o programa passando o valor em BRL e a sigla da moeda de destino:

Bash

./convert 10 USD
Exemplo de uso:

Bash

$ ./convert 12.50 eur
1.71