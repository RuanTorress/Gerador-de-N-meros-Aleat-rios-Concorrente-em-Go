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

   Exemplo de Uso
Entrada:
bash
Copiar código
Gerador de Números Aleatórios Avançado
-------------------------------------
Quantos intervalos de números aleatórios você deseja gerar? 2

Intervalo 1:
Quantos números aleatórios você deseja gerar? 5
Digite o valor mínimo: 10
Digite o valor máximo: 50

Intervalo 2:
Quantos números aleatórios você deseja gerar? 3
Digite o valor mínimo: 100
Digite o valor máximo: 200
Saída:
bash
Copiar código
Números gerados no intervalo 1:
13 45 22 11 38 

Números gerados no intervalo 2:
123 177 189 

Deseja salvar os números gerados em um arquivo? (s/n): s
Números salvos com sucesso no arquivo 'numeros_gerados.txt'.
Explicação do Código
Goroutines e Concorrência: Cada intervalo de geração de números é tratado em uma goroutine separada, permitindo que os números sejam gerados de forma concorrente.
Canais (chan): Os números gerados em cada goroutine são enviados de volta para o canal numbersChan, permitindo a comunicação entre as goroutines e o main goroutine.
Sincronização (sync.WaitGroup): Utilizamos um WaitGroup para garantir que todas as goroutines finalizem antes de prosseguir com o processamento dos números gerados.
Armazenamento em Arquivo: O programa permite salvar os números gerados em um arquivo de texto chamado numeros_gerados.txt para referência futura.
Estrutura de Arquivos
bash
Copiar código
├── main.go             # Código fonte principal
├── numeros_gerados.txt # Arquivo gerado com os números aleatórios (após execução)
└── README.md           # Este arquivo
Licença
Distribuído sob a licença MIT. Veja o arquivo LICENSE para mais informações.

Contribuições
Sinta-se à vontade para contribuir com melhorias! Para isso, siga estas etapas:

Faça um fork do projeto.
Crie uma branch para sua feature (git checkout -b feature/nova-feature).
Faça o commit das suas alterações (git commit -m 'Adiciona nova feature').
Envie para o repositório remoto (git push origin feature/nova-feature).
Abra um Pull Request.
Contato
Se você tiver alguma dúvida ou sugestão, entre em contato:

Nome: Ruan Fabio Cavalcante Carvalho Torres
Email: ruanfabio@gmail.com
markdown
Copiar código

---

### O que foi adicionado ou aprimorado:
1. **Título claro e descritivo** para explicar o propósito do projeto.
2. **Explicações detalhadas** sobre a funcionalidade e uso do projeto.
3. **Instruções passo a passo** de como executar o projeto, incluindo um exemplo de execução.
4. **Seções extras** como licença, estrutura de arquivos, e como contribuir.
5. **Detalhamento técnico** do código, explicando goroutines, canais e sincronização.
