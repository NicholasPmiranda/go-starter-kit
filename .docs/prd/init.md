<context>
# Visão Geral
O Six Flow é uma plataforma interna de gestão de projetos voltada para equipes que prestam serviços a clientes. O sistema permite organizar projetos de forma simples, intuitiva e colaborativa, com foco no controle de tarefas, gestão de projetos e comunicação entre usuários da equipe. A plataforma não oferece acesso direto aos clientes, sendo exclusiva para uso interno da equipe.

# Funcionalidades Principais
## Gestão de Usuários e Clientes
- **Cadastro e gerenciamento de usuários internos**: Permite registrar membros da equipe com diferentes níveis de acesso.
- **Cadastro e gerenciamento de clientes**: Possibilita manter uma base de dados de clientes associados aos projetos.

## Gestão de Projetos
- **Criação e organização de projetos**: Cada projeto está vinculado a um cliente específico e possui um responsável interno.
- **Visualização em lista e kanban**: Oferece diferentes formas de visualizar o andamento dos projetos.
- **Controle de status e visibilidade**: Permite definir se um projeto é público (visível para toda equipe) ou privado.

## Gestão de Tarefas
- **Criação e atribuição de tarefas**: Permite criar tarefas com descrição, responsável, prazo e prioridade.
- **Subtarefas**: Possibilita dividir tarefas complexas em itens menores e verificáveis.
- **Comentários e anexos**: Facilita a comunicação contextual dentro das tarefas.

## Sistema de Notificações
- **Notificações internas**: Alerta os usuários sobre eventos relevantes como atribuição de tarefas, novos comentários e prazos próximos.

# Experiência do Usuário
## Personas
- **Gestores de Projeto**: Profissionais que coordenam projetos e equipes, necessitando de visão ampla sobre o andamento das atividades.
- **Desenvolvedores/Designers**: Membros técnicos da equipe que executam tarefas específicas e precisam de clareza sobre suas responsabilidades.
- **Administradores**: Usuários com permissões avançadas para gerenciar configurações do sistema.

## Fluxos Principais
- **Criação e acompanhamento de projetos**: Desde o cadastro do cliente até a conclusão do projeto.
- **Gerenciamento de tarefas diárias**: Atribuição, execução e finalização de tarefas.
- **Comunicação interna**: Troca de informações via comentários e notificações.

## Considerações de UI/UX
- Interface intuitiva e responsiva para uso em diferentes dispositivos.
- Organização visual clara com uso de cores para identificação de projetos.
- Navegação simplificada entre projetos, tarefas e notificações.
  </context>
  <PRD>
# Arquitetura Técnica

## Componentes do Sistema
- **Backend API REST**: Desenvolvido em Go, responsável por toda a lógica de negócio e acesso a dados.
- **Banco de Dados**: PostgreSQL para armazenamento persistente de dados.
- **Sistema de Autenticação**: JWT para gerenciamento de sessões e controle de acesso.
- **Sistema de Armazenamento**: Para anexos e avatares de usuários.
- **Sistema de Notificações**: Mecanismo interno com suporte futuro para notificações por e-mail.

## Modelos de Dados
### Entidades Principais
1. **Usuários**:
    - Identificador único, nome, email, senha (hash), avatar
    - Relacionamentos com projetos, tarefas e notificações

2. **Clientes**:
    - Identificador único, nome, email, telefone, empresa, endereço, observações
    - Data de criação e usuário responsável
    - Relacionamento com projetos

3. **Projetos**:
    - Identificador único, nome, descrição, cor, visibilidade
    - Cliente associado e usuário responsável
    - Status (ativo, concluído, em pausa)
    - Data de criação
    - Relacionamento com tarefas

4. **Tarefas**:
    - Identificador único, título, descrição, responsável
    - Projeto associado, prazo, status, prioridade
    - Ordem para visualização
    - Datas de criação e atualização
    - Relacionamento com subtarefas, comentários e anexos

5. **Subtarefas**:
    - Identificador único, título, status de conclusão
    - Tarefa associada e ordem de exibição

6. **Comentários**:
    - Identificador único, mensagem, autor
    - Tarefa associada e data de criação

7. **Anexos**:
    - Identificador único, nome do arquivo, caminho, tipo MIME
    - Tamanho, data de upload e tarefa associada

8. **Notificações**:
    - Identificador único, tipo de evento, título, mensagem
    - Usuário destinatário, status de leitura e data de criação

## APIs e Integrações
### Endpoints Principais
- **/api/auth**: Autenticação e gerenciamento de sessões
- **/api/users**: Gerenciamento de usuários
- **/api/clients**: Gerenciamento de clientes
- **/api/projects**: Gerenciamento de projetos
- **/api/tasks**: Gerenciamento de tarefas e subtarefas
- **/api/comments**: Gerenciamento de comentários
- **/api/attachments**: Upload e download de anexos
- **/api/notifications**: Gerenciamento de notificações

### Integrações Futuras
- Sistema de e-mail para notificações externas
- Possível integração com ferramentas de calendário

## Requisitos de Infraestrutura
- Servidor para hospedagem da API Go
- Banco de dados PostgreSQL
- Sistema de armazenamento para arquivos
- Certificados SSL para comunicação segura
- Sistema de backup e recuperação de dados

# Roteiro de Desenvolvimento

## Requisitos do MVP
### Fase 1: Fundação
- Configuração do ambiente de desenvolvimento
- Implementação da estrutura básica do projeto em Go
- Configuração do banco de dados PostgreSQL
- Implementação do sistema de autenticação e autorização

### Fase 2: Funcionalidades Essenciais
- Implementação do CRUD de usuários
- Implementação do CRUD de clientes
- Implementação do CRUD de projetos
- Implementação do CRUD de tarefas e subtarefas
- Implementação da visualização em lista e kanban

### Fase 3: Comunicação e Colaboração
- Implementação do sistema de comentários
- Implementação do sistema de anexos
- Implementação do sistema de notificações internas
- Testes e ajustes finais para lançamento do MVP

## Melhorias Futuras
### Fase 4: Aprimoramentos
- Implementação de filtros avançados para projetos e tarefas
- Melhorias na interface de usuário
- Implementação de relatórios e dashboards
- Sistema de busca avançada

### Fase 5: Expansão
- Integração com serviços de e-mail para notificações externas
- Implementação de templates de projetos e tarefas
- Sistema de etiquetas e categorização
- Implementação de métricas e análises de desempenho

# Cadeia de Dependência Lógica

## Fundação (Prioridade Máxima)
1. Configuração da estrutura do projeto Go com Gingonic
2. Implementação do banco de dados PostgreSQL com SQLC e SQLX
3. Sistema de autenticação e gerenciamento de usuários
4. Implementação da camada de serviços e repositórios base

## Desenvolvimento Incremental
1. **Gestão de Clientes** → Dependência para projetos
2. **Gestão de Projetos** → Dependência para tarefas
3. **Gestão de Tarefas** → Dependência para subtarefas
4. **Subtarefas** → Funcionalidade independente
5. **Comentários e Anexos** → Dependem de tarefas
6. **Sistema de Notificações** → Depende de todas as entidades anteriores

## Priorização de Interface
1. Implementação da visualização em lista de projetos
2. Implementação da visualização em kanban
3. Implementação da interface de tarefas e subtarefas
4. Implementação da interface de comentários e anexos
5. Implementação da interface de notificações

# Riscos e Mitigações

## Desafios Técnicos
- **Risco**: Curva de aprendizado da equipe com Go e as bibliotecas escolhidas
  **Mitigação**: Criar documentação interna detalhada e realizar sessões de capacitação

- **Risco**: Performance do sistema com crescimento de dados
  **Mitigação**: Implementar paginação, indexação adequada e monitoramento de performance

- **Risco**: Segurança dos dados e controle de acesso
  **Mitigação**: Implementar testes de segurança e revisões de código focadas em segurança

## Definição do MVP
- **Risco**: Escopo muito amplo para o MVP
  **Mitigação**: Focar nas funcionalidades essenciais (usuários, clientes, projetos, tarefas básicas)

- **Risco**: Funcionalidades incompletas ou mal implementadas
  **Mitigação**: Definir critérios claros de "pronto" para cada funcionalidade

## Restrições de Recursos
- **Risco**: Dependência de conhecimento técnico concentrado
  **Mitigação**: Documentação detalhada e compartilhamento de conhecimento

- **Risco**: Tempo de desenvolvimento limitado
  **Mitigação**: Priorização clara de funcionalidades e desenvolvimento incremental

# Apêndice

## Decisões Arquiteturais
- **Linguagem**: Go (Golang) escolhida por sua performance, escalabilidade e suporte a concorrência
- **Roteador HTTP**: Jingonic escolhido por seu suporte a auto-recovery e comunidade ativa
- **Banco de Dados**: PostgreSQL com SQLC e SQLX para controle direto sobre queries e performance

## Convenções de Desenvolvimento
- Estrutura de pastas seguindo o padrão golang-standards/project-layout
- Nomenclatura de arquivos e funções seguindo convenções Go
- Organização da lógica de aplicação dentro da pasta internal/
- Criação de serviços para abstrair comportamentos
- Agrupamento lógico de serviços relacionados

## Referências Técnicas
- Documentação oficial do Go: https://golang.org/doc/
- Documentação do PostgreSQL: https://www.postgresql.org/docs/
- Documentação do SQLC: https://sqlc.dev/
- Padrões de projeto em Go: https://github.com/tmrts/go-patterns
  </PRD>
