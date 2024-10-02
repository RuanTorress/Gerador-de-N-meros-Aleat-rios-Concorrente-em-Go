# Gerador de Números Aleatórios Concorrente em Go

Este projeto é um gerador de números aleatórios escrito em Go que suporta geração concorrente de múltiplos intervalos. Ele também permite salvar os números gerados em um arquivo de texto. O uso de goroutines e canais torna o código eficiente e escalável, demonstrando o poder da concorrência no Go.

## Funcionalidades

- Gera números aleatórios em diferentes intervalos simultaneamente usando goroutines.
- Permite especificar a quantidade de números, o valor mínimo e o valor máximo para cada intervalo.
- Oferece a opção de salvar os números gerados em um arquivo de texto.
- Garante a concorrência e sincronização utilizando `sync.WaitGroup` e `channels`.

## Como Executar

### Pré-requisitos

- Certifique-se de ter o [Go](https://golang.org/dl/) instalado (versão 1.15 ou superior).
- Um ambiente de terminal para rodar o programa.

### Passos

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
