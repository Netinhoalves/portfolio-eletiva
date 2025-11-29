<table border="0" width="100%">
  <tr>
    <td width="70%">
      <h1>Portfólio de Tópicos Especiais em Linguagem de Programação</h1>
      <p>
        <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go">
        <img src="https://img.shields.io/badge/Curso-ADS-blue">
        <img src="https://img.shields.io/badge/Semestre-5º-orange">
        <img src="https://img.shields.io/badge/Status-Entregue-success">
      </p>
      <blockquote>
        <b>Aluno:</b> Dirceu Alves Neto<br>
        <b>Professor:</b> Alex Fernando de Araujo<br>
        <b>Unidade Curricular:</b> Tópicos Especiais em Linguagem de Programação<br>
        <b>Instituição:</b> Instituto Federal de Mato Grosso do Sul<br>
        <b>Data de Início:</b> 22/08/2025<br>
        <b>Entrega:</b> 28/11/2025
      </blockquote>
    </td>
    <td width="30%" align="center">
      <img src="./assets/perfil.jpeg" width="150" alt="Foto de Perfil" style="border-radius: 5%;">
    </td>
  </tr>
</table>

<br>

## 📖 1. O Contrato Pedagógico (22/08)
Este documento sintetiza a jornada de aprendizado na disciplina **Tópicos Especiais em Linguagem de Programação**.

**Definições da Aula Inaugural:**
* **Objetivo:** Compreender os fundamentos da linguagem para, posteriormente, criar aplicações com processamento concorrente.
* **Metodologia:** Foco total na prática e experimentação.
* **Acordo de Produção:** Foi estabelecida uma meta média de **10 exercícios semanais**.
* **Avaliação:** Este portfólio compõe a 1ª Nota, documentando a evolução cronológica, instalações e desafios técnicos enfrentados.

---

## ⏳ 2. Linha do Tempo (Cronograma de Aulas)

Abaixo, o registro detalhado das atividades realizadas às sextas-feiras, incluindo aulas expositivas, apresentações, desenvolvimento assíncrono e eventos acadêmicos.

| Data | Tipo | Atividade / Tópico Abordado | Evidências |
| :--- | :--- | :--- | :--- |
| **22/08** | **Começo** | **Setup e Acordos:** Definição da ementa e configuração do ambiente (Instalação do Go + VS Code). | *Presencial* |
| **29/08** | **Teórica** | **Sintaxe e Variáveis:** Diferença entre declaração explícita (`var`) e implícita (`:=`). Recebimento da **Lista (26 Exercícios)** para prática contínua. | [**Ver Lista**](#lista) |
| **05/09** | **Prática** | **Desenvolvimento Autônomo:** Resolução dos primeiros exercícios da lista. | [Códigos](./src/lista) |
| **12/09** | **Apresentação** | **Fundamentos e Comparativos:** Aula magna sobre histórico, Garbage Collector e arquitetura. **Atividade:** Apresentação em grupos comparando Go vs. Outras Linguagens. | [**Ver Análise**](#apresentacao) |
| **19/09** | **Evento** | **Semana de Administração:** Participação em palestras interdisciplinares (Liberado da aula técnica). | *Presença* |
| **26/09** | **Mentoria** | **Checkpoint:** Reunião individual com o professor para validar o andamento do portfólio e exercícios realizados. | *Presencial* |
| **03/10** | **Prática** | **Algoritmos Complexos:** Foco na resolução de exercícios avançados da Lista (ex: Torres de Hanói e Recursividade). | [Códigos](./src/lista) |
| **10/10** | **Laboratório** | **Maratona "Go By Example":** Início da bateria de testes baseada na documentação oficial (conforme indicado na aula de 12/09). | [Códigos](./src/go-by-example) |
| **17/10** | **Erro/Teste** | **Desafio de Imagem PPM (P3):** Tentativa de gerar um arquivo *Portable Pixel Map* (formato ASCII P3) via matriz de inteiros. **Falha Técnica:** Devido à falta de ambiente local configurado no laboratório, utilizamos o *Go Playground*. O ambiente "Sandbox" limitou a saída de texto e impediu a criação do arquivo `.ppm` em disco. | [Tentativa](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/desafio-imagem/ppm.go) |
| **24/10** | **Prática** | **Continuidade:** Sequência nos exercícios do *Go By Example* devido à persistência de problemas de ambiente. | [Códigos](./src/go-by-example) |
| **31/10** | **Sem Aula** | Não houve aula. Estudo individual. | - |
| **07/11** | **Evento** | **Dia da Consciência Negra:** Apresentações culturais no campus. | *Presença* |
| **14/11** | **Sem Aula** | **Colação de Grau** | - |
| **21/11** | **Sem Aula** | **Emenda de Feriado:** Não houve aula devido ao feriado da Consciência Negra (20/11). | - |
| **28/11** | **Entrega** | Finalização e entrega do Portfólio Digital. | **Entrega** |

---
<div id="apresentacao"></div>

## 🗣️ 3. Relatório da Apresentação (12/09)
Nesta data, atuei como **observador analítico** das apresentações dos grupos, registrando os oradores principais e os pontos técnicos discutidos.

### Destaques das Apresentações:

* **Grupo 1: Concorrência (Henrique, Nathan, Azam)**
    * **Nathan** focou quase integralmente na leitura da defesa teórica sobre concorrência. Houve uma observação do professor (Alex) sobre a clareza da exposição.

* **Grupo 2: Simplicidade e Desempenho (João, Cazuo, Maysson)**
    * **João** explicou bem os conceitos de simplicidade do Go, embora tenha demonstrado certa insegurança e prolongado a explicação em alguns momentos (enrolou um pouco).

* **Grupo 3: Compilação e Tipagem (Rafael, Vitor, Lucas)**
    * **Rafael** Realizou uma explicação clara, detalhando bem as diferenças de tempo de execução e tipagem entre as linguagens, demonstrando domínio dos termos.

* **Grupo 4: Garbage Collector (Gêmeos, Jonathan)**
    * **Jonathan** assumiu o protagonismo utilizando a lousa (canetão) para explicar visualmente os mecanismos do Coletor de Lixo, citando exemplos de outros grupos. A explicação foi didática e completa, sem necessidade de intervenção do professor.

* **Grupo 5: Biblioteca Padrão (Nildemar, Rayane, Misael)**
    * **Nildemar** conduziu a apresentação (com apoio de leitura), mas pontuou bem as vantagens de performance embutida e pacotes de segurança nativos do Go em comparação a certas limitações do Python.

---
<div id="lista"></div>

## 📝 4. Lista Cumulativa (29/08)
Esta lista serviu como base de atividades para os dias de estudo autônomo, cobrindo a meta de exercícios semanais.

<details>
<summary><b>🔻 Clique para expandir os 26 Exercícios Desenvolvidos</b></summary>

1.  [**Soma**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio01.go) Algoritmo que soma dois números inteiros.
2.  [**Divisão**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio02.go) Divisão de dois números inteiros.
3.  [**Sucessor/Antecessor**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio03.go) Input de inteiro e output de vizinhos.
4.  [**Verificação Numérica**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio04.go) Checar se é Ímpar, Par, Positivo ou Negativo.
5.  [**Números Primos**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio05.go) Verificação matemática.
6.  [**Ordenação (Sort)**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio06.go) Sequência numérica.
7.  [**Ordenação de Caracteres**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio07.go) Sequência ascendente.
8.  [**Árvore de Decisão**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio08.go) Baseada em sequência de respostas.
9.  [**Ponteiros**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio09.go) Imprimir variável e seu endereço de memória (`&var`).
10. [**Torres de Hanói**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio10.go) Resolução do quebra-cabeça (Recursividade).
11. [**Dia da Semana**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio11.go) Algoritmo matemático baseado na data de nascimento.
12. [**Igualdade Booleana**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio12.go) Retorno True/False.
13. [**Estatística**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio13.go) Moda de uma sequência.
14. [**Palíndromos**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio14.go) Verificação de frases/palavras (ex: "Arara").
15. [**Geometria**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio15.go) Área do Retângulo.
16. [**Conversor**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio16.go) Unidades de temperatura (Celsius/Fahrenheit).
17. [**Jogo**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio17.go) Simulação de Adivinhação.
18. [**Desafio JPEG (Matrizes)**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio18.go) Escreva um algoritmo que receba três matrizes de inteiros (uint8) e gere um arquivo de imagem JPEG.
19. [**Contador de Letras**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio19.go) Vogais e Consoantes.
20. [**Padrões**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio20.go) Ocorrência de uma palavra específica num texto.
21. [**Fatorial**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio21.go) Cálculo de N!.
22. [**Hello World**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio22.go) Revisão de estrutura básica.
23. [**Saúde**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio23.go) Cálculo de IMC (Padrão Brasileiro).
24. [**Matemática**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio24.go) Mínimo Múltiplo Comum (MMC).
25. [**Média**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio25.go) Cálculo entre dois ou mais números.
26. [**Bônus (Algoritmo Secreto)**](https://github.com/Netinhoalves/portfolio-eletiva/blob/main/src/lista/exercicio26.go) Algoritmo Evolutivo Genético.

</details>

---

## 💻 5. Ambiente de Desenvolvimento
Conforme requisito de documentação, abaixo está o setup utilizado durante o semestre:

* **Linguagem:]** [Go](https://go.dev/) (Golang) versão 1.25.1
* **Editor:** VS Code
* **Extensões:** *Go Team at Google* (IntelliSense, Linting)
* **Teste de Instalação:**
    ```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Ambiente Configurado com Sucesso - 22/08")
    }
    ```

---

## 🚀 6. Conclusão e Próximos Passos
Este portfólio documenta a construção da base sólida necessária para operar a linguagem Go.
Apesar dos desafios de infraestrutura (laboratório) e interrupções no calendário, o volume de exercícios (Lista + Go By Example) garantiu a fluência na sintaxe.

**Estou preparado para o próximo módulo.**

---
*Este portfólio foi gerado como requisito avaliativo da disciplina.*