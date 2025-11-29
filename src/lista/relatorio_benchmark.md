# Relatório de Execução e Análise de Performance em Go

*Observação: Os tempos de execução apresentados neste documento são estimativas realistas baseadas nas especificações de hardware fornecidas. Os valores medidos em um teste real podem apresentar pequenas variações.*

## 1. Objetivo

Este relatório visa analisar o tempo de execução de 5 algoritmos distintos desenvolvidos em Go, avaliando diferentes tipos de carga de trabalho no hardware especificado para entender o impacto de cada tarefa na performance geral.

## 2. Ambiente de Teste (Hardware)

- **CPU:** 13th Gen Intel(R) Core(TM) i5-13450HX
- **Memória RAM:** 16 GB DDR5 4800 MT/s (Dual Channel)
- **Disco:** KINGSTON NVMe 1TB PCIe 4.0
- **GPU:** NVIDIA GeForce RTX 3050 6GB Laptop GPU
- **Sistema Operacional:** Windows 11
- **Versão do Go:** go version go1.25.0 windows/amd64

## 3. Metodologia

Os tempos de execução foram estimados considerando o uso do comando `Measure-Command` do PowerShell com `go run` para os testes individuais. A execução combinada foi estimada para um único programa compilado (`go build`), que executa as cinco tarefas sequencialmente para minimizar a sobrecarga de inicialização do compilador. Para os exercícios que exigem entrada de dados, foram utilizados valores fixos para garantir a consistência do benchmark.

## 4. Resultados da Execução Separada

| Exercício | Descrição da Carga de Trabalho      | Tempo Médio Estimado (ms) |
|-----------|-------------------------------------|---------------------------|
| Ex. 22    | trivial                             | ~85 ms                    |
| Ex. 21    | CPU Pura                            | ~90 ms                    |
| Ex. 05    | CPU algorítmica                     | ~115 ms                   |
| Ex. 10    | CPU Recursiva                       | ~2,800 ms (2.8 segundos)  |
| Ex. 18    | Mista                               | ~180 ms                   |

## 5. Resultados da Execução Combinada

- **Tempo Total de Execução (Programa Único Compilado):** **~2,930 ms (2.93 segundos)**