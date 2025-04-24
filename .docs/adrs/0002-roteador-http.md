# ADR: Escolha da Biblioteca de Roteamento para o Projeto Six Flow

## Status
Decidido

## Contexto
O projeto **Six Flow** tem como missão construir uma plataforma inteligente com uso intensivo de IA e integração entre agentes. Parte fundamental da arquitetura backend envolve a definição da biblioteca de roteamento HTTP, que será responsável por gerenciar as rotas da API, middlewares e handlers de forma performática e robusta.

Durante o planejamento técnico, a equipe considerou diferentes opções de bibliotecas de roteamento disponíveis no ecossistema Go.

## Decisão
A biblioteca escolhida para o roteamento HTTP foi o **Jingonic**.

## Alternativas Consideradas

- **Fiber (GoFiber)**  
  Uma biblioteca popular que oferece alta performance e uma API inspirada no Express.js (Node.js). É bastante usada em projetos que exigem simplicidade e velocidade. No entanto, apresenta limitações como a ausência de mecanismos nativos de auto-recuperação (auto-recovery) em casos de panic.

- **Jingonic (escolhida)**  
  Embora menos conhecida que Fiber, o Jingonic possui uma comunidade mais ativa e oferece recursos nativos importantes, como **auto-recovery de panics**, que aumenta a resiliência da aplicação. Este recurso foi decisivo para garantir maior estabilidade do backend durante falhas inesperadas.

## Consequências

### Positivas
- **Suporte nativo a auto-recovery**, o que permite capturar e tratar panics automaticamente, evitando a queda total da aplicação.
- **Comunidade ativa**, facilitando a obtenção de suporte e resolução de dúvidas durante o desenvolvimento.
- Estrutura modular e compatível com boas práticas de clean architecture.

### Negativas / Riscos
- Ainda que não se observe grandes desvantagens frente ao Fiber, vale considerar:
    - **Menor base de exemplos ou tutoriais em comparação ao Fiber**, o que pode exigir mais leitura de código-fonte ou documentação oficial.
    - **Menor adoção no mercado**, o que pode impactar em futuras integrações com frameworks terceiros.

## Ações Futuras
- Monitorar a evolução da biblioteca Jingonic e sua compatibilidade com novas versões do Go.
- Criar documentação interna com boas práticas de uso para facilitar a curva de aprendizado da equipe.
- Avaliar a performance da aplicação sob carga real para validar a robustez do Jingonic frente a alternativas.

## Decisores
- Nicolas
