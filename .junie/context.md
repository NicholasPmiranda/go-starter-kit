# Documento de Contexto: Sistema de Gestão de Projetos Interno

## 🎯 Visão Geral do Sistema

Trata-se de uma plataforma interna voltada para equipes que prestam serviços a clientes e desejam organizar seus projetos de forma simples, intuitiva e colaborativa. O foco está no controle de tarefas, gestão de projetos e comunicação entre usuários da equipe — sem acesso direto dos clientes à plataforma.

---

## 👥 Usuários e Clientes

### Usuários (Equipe Interna)

- Representam os operadores da plataforma: desenvolvedores, designers, gestores, etc.
- Possuem acesso completo às funcionalidades da aplicação.

**Campos:**

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador único |
| nome | String | Nome completo |
| email | String | E-mail para login |
| senha | String | Senha do usuário (hash gerado no back-end) |
| avatar_path | String | Caminho da imagem de perfil (opcional) |

---

### Clientes

- Representam os clientes da equipe, vinculados aos projetos.
- Não possuem acesso direto à aplicação.

**Campos:**

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador único |
| nome | String | Nome do cliente |
| email | String | Contato principal |
| telefone | String | Opcional |
| empresa | String | Nome da empresa associada |
| endereço | Texto | Endereço completo (opcional) |
| observações | Texto | Notas gerais |
| criado_em | Date | Data de cadastro |
| responsável_id | Int | Usuário responsável por esse cliente |

---

## 📁 Organização de Projetos

Cada projeto está associado a um cliente e gerenciado internamente pelos usuários.

**Campos:**

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador único |
| nome | String | Nome do projeto |
| descricao | Texto | Descrição mais detalhada |
| cor | String | Cor visual de identificação |
| visibilidade | Enum | `publico` ou `privado` (apenas para uso interno) |
| cliente_id | Int | Cliente relacionado |
| responsavel_id | Int | Usuário que lidera o projeto |
| criado_em | Date | Data de criação |
| status | Enum | `ativo`, `concluido`, `em pausa` |

> Visualizações suportadas: Lista e Kanban com colunas fixas.
>

---

## ✅ Tarefas e Subtarefas

### Tarefas

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador |
| projeto_id | Int | Projeto vinculado |
| titulo | String | Nome curto |
| descricao | Texto | Detalhes da tarefa |
| responsavel_id | Int | Usuário atribuído |
| prazo | Date | Data limite |
| status | Enum | `a_fazer`, `em_progresso`, `concluido` |
| ordem | Int | Para ordenação visual |
| prioridade | Enum | (opcional) `baixa`, `média`, `alta` |
| criado_em | DateTime | Data de criação |
| atualizado_em | DateTime | Última modificação |

### Subtarefas

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador |
| tarefa_id | Int | Referência à tarefa principal |
| titulo | String | Nome da subtarefa |
| concluida | Bool | Status |
| ordem | Int | Ordem de exibição |

---

## 💬 Comunicação em Tarefas

### Comentários

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador |
| tarefa_id | Int | Referência à tarefa |
| autor_id | Int | Usuário que comentou |
| mensagem | Texto | Texto do comentário |
| criado_em | DateTime | Data de envio |

> Sem suporte a menções. Comentários podem futuramente ser estendidos para edição ou exclusão.
>

### Anexos

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador |
| tarefa_id | Int | Referência à tarefa |
| nome_arquivo | String | Nome original |
| path_arquivo | String | Caminho salvo no back-end |
| tipo | String | Tipo MIME |
| tamanho | Int | Tamanho em bytes |
| criado_em | DateTime | Data de upload |

---

## 🔔 Notificações

### Notificações Internas

| Campo | Tipo | Descrição |
| --- | --- | --- |
| id | Int | Identificador |
| usuario_id | Int | Destinatário |
| tipo_evento | Enum | Tipo da notificação |
| titulo | String | Título resumido |
| mensagem | Texto | Texto explicativo (opcional) |
| visto | Bool | Status de leitura |
| criado_em | DateTime | Data da notificação |

> Eventos que geram notificação: atribuição de tarefa, novo comentário, prazo próximo.
>

---

## 🚀 Resumo das Decisões

- Não há conceito de “time”, apenas usuários internos e clientes.
- Os projetos são internos e não expostos aos clientes.
- Visualização do projeto é exclusivamente por **Lista** e **Kanban** fixo.
- Preferências de interface e histórico de ações foram **excluídos** do escopo.
- Comentários e anexos são simples e diretos.
- Toda notificação é gerada internamente, com **opcional envio por e-mail** (em fase futura).
