# ADR: Escolha do Banco de Dados e Ferramentas de Acesso para o Projeto Six Flow

## Status
Decidido

## Contexto
No projeto **Six Flow**, era necessário definir qual banco de dados relacional seria utilizado e também quais bibliotecas de acesso a dados seriam adotadas para interagir com esse banco, considerando a linguagem Go como base do projeto.

O objetivo principal era escolher uma combinação que oferecesse **performance**, **flexibilidade** e **produtividade**, além de garantir escalabilidade no médio e longo prazo. A solução precisava também se integrar bem com o modelo de desenvolvimento desejado pela equipe, que prioriza controle direto sobre as queries e estruturação limpa do código.

## Decisão
Foi decidido utilizar **PostgreSQL** como banco de dados e **SQLC + SQLX** como ferramentas para geração e manipulação de queries SQL.

## Alternativas Consideradas

### Banco de Dados
- **PostgreSQL (escolhido)**  
  Banco de dados relacional robusto, leve, com excelente suporte a operações complexas, JSON, concorrência e extensibilidade.

- **MySQL**  
  Amplamente utilizado no mercado, com boa performance, porém com limitações em algumas operações transacionais e menos recursos avançados que o PostgreSQL.

### Acesso a Dados
- **SQLC + SQLX (escolhido)**  
  SQLC permite gerar código Go tipado a partir de queries SQL puras, enquanto o SQLX expande o suporte nativo de SQL do Go, oferecendo mais flexibilidade e produtividade.

- **ORMs (como GORM ou Ent)**  
  Soluções que abstraem a manipulação de SQL, permitindo desenvolvimento mais rápido para quem está acostumado, mas com possíveis impactos em performance, perda de controle sobre as queries e dificuldade em debug para cenários complexos.

## Consequências

### Positivas
- PostgreSQL é mais leve e robusto que o MySQL em diversas operações e com maior capacidade de extensibilidade.
- SQLC + SQLX oferecem excelente performance e controle total sobre as queries.
- Maior produtividade na escrita de código SQL com segurança de tipos.
- Boa escalabilidade futura com uma stack de dados simples e eficiente.
- Mais alinhado com boas práticas modernas em desenvolvimento backend com Go.

### Negativas / Riscos
- Curva de aprendizado maior para desenvolvedores acostumados apenas com ORMs.
- Necessidade de escrever queries manualmente, o que pode ser mais custoso inicialmente para quem não domina SQL.

## Ações Futuras
- Criar materiais de onboarding para novos devs sobre uso de SQLC e SQLX.
- Escrever um repositório de exemplos internos com boas práticas para acesso a dados.
- Avaliar constantemente a necessidade de abstrações para manter o equilíbrio entre produtividade e controle.

## Decisores
- Nicolas
