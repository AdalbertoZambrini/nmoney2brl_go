# CLI Currency Converter em Go 🪙

Um conversor de moedas simples via linha de comando desenvolvido em Go. O programa recebe um valor em Reais (BRL) e o converte para a moeda de destino escolhida utilizando taxas de câmbio pré-definidas.

## 🚀 Funcionalidades

* Conversão rápida e direta direto no terminal.
* Suporte a diversas moedas (USD, EUR, JPY, GBP, etc.).
* Tratamento de erros para entradas inválidas ou moedas não suportadas.
* Case-insensitive: funciona independentemente de você digitar a moeda em maiúsculas ou minúsculas (ex: `usd` ou `USD`).

## 💻 Como executar

1. Certifique-se de ter o [Go instalado](https://go.dev/) na sua máquina.
2. Clone este repositório ou baixe o arquivo `main.go`.
3. Abra o terminal na pasta do projeto e compile o código:
```bash
go build -o convert main.go

```


4. Execute o programa passando o valor em BRL e a sigla da moeda de destino:
```bash
./convert 10 USD

```



**Exemplo de uso:**

```bash
$ ./convert 12.50 eur
1.71

```