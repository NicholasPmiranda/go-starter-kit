# Como Disparar Emails no Go Starter Kit

## Visão Geral

O Go Starter Kit inclui um provedor de email que facilita o envio de mensagens usando templates HTML. Este documento explica como configurar e utilizar este recurso em sua aplicação.

## Configuração

O provedor de email é configurado através de variáveis de ambiente definidas no arquivo `.env`:

```
MAIL_HOST=smtp.example.com
MAIL_PORT=587
MAIL_USERNAME=user@example.com
MAIL_PASSWORD=password
MAIL_FROM_ADDRESS=noreply@example.com
MAIL_FROM_NAME="Example App"
```

Certifique-se de que estas variáveis estejam configuradas corretamente antes de tentar enviar emails.

## Estrutura de Arquivos

O provedor de email é implementado nos seguintes arquivos:

```
config/emailProvider/
├── emailProvider.go    # Implementação do provedor
└── types.go            # Definição de tipos
```

Os templates de email são armazenados no diretório `template/` na raiz do projeto.

## Como Enviar Emails

### 1. Importar o Pacote

Primeiro, importe o pacote do provedor de email:

```go
import (
    emailprovider "sixTask/config/emailProvider"
)
```

### 2. Criar uma Mensagem de Email

Crie uma estrutura `EmailMessage` com os detalhes do email:

```go
// Definir os destinatários
to := []string{
    "destinatario@exemplo.com",
    "outro@exemplo.com",
}

// Criar os dados para o template.sql
templateData := map[string]interface{}{
    "Nome":    "João Silva",
    "Empresa": "Sua Empresa",
}

// Criar a mensagem
emailMsg := emailprovider.EmailMessage{
    To:           to,
    Cc:           []string{"copia@exemplo.com"},
    Bcc:          []string{"copiaoculta@exemplo.com"},
    Subject:      "Assunto do Email",
    Template:     "cadastro",  // Nome do template.sql HTML (sem a extensão)
    TemplateData: templateData,
    Attachments:  []emailprovider.EmailAttachment{
        {
            Path:     "storage/app/documento.pdf",
            Filename: "documento.pdf",
        },
    },
}
```

### 3. Enviar o Email

Envie o email usando a função `SendMail`:

```go
err := emailprovider.SendMail(emailMsg)
if err != nil {
    log.Printf("Erro ao enviar email: %v", err)
    return err
}
```

## Trabalhando com Templates

### Criando um Template

Os templates de email são arquivos HTML que usam a sintaxe de templates do Go. Crie um arquivo HTML no diretório `template/` com o nome desejado (por exemplo, `cadastro.html`):

```html
<!DOCTYPE html>
<html lang="pt">
<head>
    <meta charset="UTF-8">
    <title>Email Bonito</title>
</head>
<body style="font-family: Arial, sans-serif; background-color: #f7f9fc; padding: 20px;">
    <div style="max-width: 600px; margin: auto; background-color: #ffffff; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); overflow: hidden;">
        <div style="background-color: #4a90e2; color: white; padding: 15px 20px;">
            <h1 style="margin: 0;">Olá, {{ .Nome }}!</h1>
        </div>
        <div style="padding: 20px;">
            <p style="font-size: 16px; line-height: 1.5; color: #555;">
                Este é um email de teste usando um template mais bonito e agradável visualmente. Esperamos que goste!
            </p>
            <p style="font-size: 14px; color: #999;">
                Atenciosamente,<br>{{ .Empresa }}
            </p>
        </div>
    </div>
</body>
</html>
```

### Passando Dados para o Template

Para passar valores para o template, você precisa criar uma estrutura ou um mapa com os campos que correspondem às variáveis usadas no template:

```go
// Usando um mapa
templateData := map[string]interface{}{
    "Nome":    "João Silva",
    "Empresa": "Sua Empresa",
}

// OU usando uma estrutura
type CadastroTemplateData struct {
    Nome    string
    Empresa string
}

templateData := CadastroTemplateData{
    Nome:    "João Silva",
    Empresa: "Sua Empresa",
}
```

As variáveis no template são acessadas usando a sintaxe `{{ .NomeDaVariavel }}`.

## Exemplo Completo

Aqui está um exemplo completo de como enviar um email de boas-vindas para um novo usuário:

```go
package handler

import (
    emailprovider "sixTask/config/emailProvider"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "time"

    "sixTask/internal/database"
)

func StartHandler(c *gin.Context) {
    // Conecta ao banco de dados usando sqlx
    db := database.ConnectDBX()
    defer db.Close()

    // Gera um email único usando timestamp
    timestamp := time.Now().UnixNano()
    emailUnico := fmt.Sprintf("usuario_%d@exemplo.com", timestamp)

    // Cria um novo usuário usando sqlx
    log.Printf("Criando novo usuário com email: %s", emailUnico)

    var novoUsuario database.User
    err := db.QueryRowx(
        "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *",
        "Novo Usuário", emailUnico, "senha123",
    ).StructScan(&novoUsuario)

    if err != nil {
        c.JSON(500, gin.H{
            "erro": "Erro ao criar usuário: " + err.Error(),
        })
        return
    }

    // Definir os destinatários do email
    to := []string{
        "teste@teste.com",
        "teste2@testeco",
    }

    // Criar os dados para o template.sql
    templateData := map[string]interface{}{
        "Nome":    novoUsuario.Name,
        "Empresa": "Sua Empresa",
    }

    // Enviar o email de boas-vindas
    emailprovider.SendMail(emailprovider.EmailMessage{
        To:           to,
        Subject:      "Bem-vindo ao nosso serviço",
        Template:     "cadastro",
        TemplateData: templateData,
    })

    log.Print("email enviado")
    c.JSON(200, gin.H{
        "mensagem":      "Usuário criado com sucesso",
        "usuarioCriado": novoUsuario,
    })
}
```

## Dicas e Boas Práticas

1. **Tratamento de Erros**: Sempre verifique e trate os erros retornados pela função `SendMail`.
2. **Templates Responsivos**: Crie templates de email responsivos que funcionem bem em dispositivos móveis.
3. **Testes**: Teste seus emails em diferentes clientes de email para garantir compatibilidade.
4. **Anexos**: Não envie anexos muito grandes, pois podem ser bloqueados por servidores de email.
5. **Logs**: Adicione logs para rastrear o envio de emails e diagnosticar problemas.
6. **Filas**: Para aplicações com alto volume, considere usar filas para processar emails em segundo plano.

## Solução de Problemas

### Email não está sendo enviado

Verifique:
- As configurações SMTP no arquivo `.env`
- Se o servidor SMTP está acessível
- Se há erros nos logs da aplicação

### Variáveis não aparecem no template

Verifique:
- Se os nomes das variáveis no template correspondem exatamente aos nomes no mapa ou estrutura
- Se está usando a sintaxe correta: `{{ .NomeVariavel }}`
- Se o template está sendo carregado corretamente

### Anexos não funcionam

Verifique:
- Se o caminho do arquivo está correto
- Se o arquivo existe e é acessível pela aplicação
- Se o tamanho do arquivo não excede os limites do servidor SMTP
