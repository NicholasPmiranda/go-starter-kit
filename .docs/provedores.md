# Como Usar Provedores

## Visão Geral

Os provedores são componentes que encapsulam a integração com serviços externos, como envio de e-mails e armazenamento de arquivos. No Go Starter Kit, os provedores são organizados no diretório `config/` e fornecem interfaces simples para utilização desses serviços.

## Provedores Disponíveis

O Go Starter Kit inclui os seguintes provedores:

1. **Email Provider**: Para envio de e-mails usando SMTP
2. **Storage Provider**: Para armazenamento e recuperação de arquivos

## Provedor de Email

### Estrutura

O provedor de email é implementado nos arquivos:

```
config/emailProvider/
├── emailProvider.go    # Implementação do provedor
└── types.go            # Definição de tipos
```

### Configuração

O provedor de email é configurado através de variáveis de ambiente definidas no arquivo `.env`:

```
MAIL_HOST=smtp.example.com
MAIL_PORT=587
MAIL_USERNAME=user@example.com
MAIL_PASSWORD=password
MAIL_FROM_ADDRESS=noreply@example.com
MAIL_FROM_NAME="Example App"
```

### Como Usar

#### 1. Importar o Pacote

```go
import (
    emailprovider "boilerPlate/config/emailProvider"
)
```

#### 2. Criar uma Mensagem de Email

```go
// Criar uma mensagem de email
emailMsg := emailprovider.EmailMessage{
    To:           []string{"destinatario@example.com"},
    Cc:           []string{"copia@example.com"},
    Bcc:          []string{"copiaoculta@example.com"},
    Subject:      "Assunto do Email",
    Template:     "cadastro",  // Nome do template HTML (sem a extensão)
    TemplateData: map[string]interface{}{
        "nome":    "João Silva",
        "empresa": "Empresa XYZ",
        "link":    "https://example.com/confirmar",
    },
    Attachments: []emailprovider.Attachment{
        {
            Path:     "storage/app/documento.pdf",
            Filename: "documento.pdf",
        },
    },
}
```

#### 3. Enviar o Email

```go
// Enviar o email
err := emailprovider.SendMail(emailMsg)
if err != nil {
    log.Printf("Erro ao enviar email: %v", err)
    return err
}
```

### Templates de Email

Os templates de email são arquivos HTML localizados no diretório `template/`. Por exemplo, para o template "cadastro", o arquivo seria `template/cadastro.html`.

Exemplo de template:

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Bem-vindo!</title>
</head>
<body>
    <h1>Olá, {{.nome}}!</h1>
    <p>Bem-vindo à {{.empresa}}.</p>
    <p>Para confirmar seu cadastro, <a href="{{.link}}">clique aqui</a>.</p>
</body>
</html>
```

Os templates usam a sintaxe de templates do Go, onde `{{.nome}}` é substituído pelo valor da chave "nome" no mapa `TemplateData`.

## Provedor de Storage

### Estrutura

O provedor de storage é implementado no arquivo:

```
config/storageProvider/storageProvider.go
```

### Como Usar

#### 1. Importar o Pacote

```go
import (
    storageprovider "boilerPlate/config/storageProvider"
)
```

#### 2. Salvar um Arquivo

```go
// Salvar um arquivo
filePath, err := storageprovider.SaveFile(fileData, "nome-do-arquivo.pdf", "app/pdfs")
if err != nil {
    log.Printf("Erro ao salvar arquivo: %v", err)
    return err
}

// filePath contém o caminho completo para o arquivo salvo
// Ex: "storage/app/pdfs/57ccc4716ec6e54b875abedadfb91b8e.pdf"
```

#### 3. Obter a URL de um Arquivo

```go
// Obter a URL de um arquivo
fileURL := storageprovider.GetFileURL(filePath)

// fileURL contém a URL para acessar o arquivo
// Ex: "/storage/app/pdfs/57ccc4716ec6e54b875abedadfb91b8e.pdf"
```

#### 4. Verificar se um Arquivo Existe

```go
// Verificar se um arquivo existe
exists := storageprovider.FileExists(filePath)
if !exists {
    log.Printf("Arquivo não encontrado: %s", filePath)
    return errors.New("arquivo não encontrado")
}
```

#### 5. Excluir um Arquivo

```go
// Excluir um arquivo
err := storageprovider.DeleteFile(filePath)
if err != nil {
    log.Printf("Erro ao excluir arquivo: %v", err)
    return err
}
```

### Diretórios de Storage

O provedor de storage utiliza o diretório `storage/` como base para armazenamento de arquivos. Os subdiretórios comuns são:

- `storage/app/`: Para arquivos gerais da aplicação
- `storage/app/pdfs/`: Para arquivos PDF
- `storage/planilha/`: Para planilhas Excel
- `storage/log/`: Para arquivos de log

## Exemplo Completo: Upload e Envio de Arquivo por Email

```go
package filehandler

import (
    "boilerPlate/config/emailProvider"
    "boilerPlate/config/storageProvider"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func UploadAndSendFile(c *gin.Context) {
    // 1. Receber o arquivo
    file, err := c.FormFile("arquivo")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Arquivo não fornecido",
        })
        return
    }

    // 2. Abrir o arquivo
    src, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erro ao abrir arquivo",
        })
        return
    }
    defer src.Close()

    // 3. Ler o conteúdo do arquivo
    fileData := make([]byte, file.Size)
    _, err = src.Read(fileData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erro ao ler arquivo",
        })
        return
    }

    // 4. Salvar o arquivo
    filePath, err := storageProvider.SaveFile(fileData, file.Filename, "app/documentos")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erro ao salvar arquivo",
        })
        return
    }

    // 5. Enviar email com o arquivo anexado
    emailMsg := emailProvider.EmailMessage{
        To:           []string{c.PostForm("email")},
        Subject:      "Seu arquivo foi recebido",
        Template:     "upload_confirmacao",
        TemplateData: map[string]interface{}{
            "nome":     c.PostForm("nome"),
            "arquivo":  file.Filename,
        },
        Attachments: []emailProvider.Attachment{
            {
                Path:     filePath,
                Filename: file.Filename,
            },
        },
    }

    err = emailProvider.SendMail(emailMsg)
    if err != nil {
        log.Printf("Erro ao enviar email: %v", err)
        // Continua mesmo com erro no email
    }

    // 6. Retornar resposta de sucesso
    c.JSON(http.StatusOK, gin.H{
        "message": "Arquivo recebido e email enviado com sucesso",
        "file_path": filePath,
    })
}
```

## Criando Novos Provedores

Se você precisar integrar com outros serviços externos, pode criar novos provedores seguindo o mesmo padrão:

1. Crie um novo diretório em `config/` para o provedor (ex: `config/smsProvider/`)
2. Implemente a lógica de integração com o serviço externo
3. Forneça uma interface simples para utilização do serviço

Exemplo de estrutura para um novo provedor:

```go
package smsprovider

import (
    "os"
    "github.com/twilio/twilio-go"
    twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var client *twilio.RestClient
var fromNumber string

func init() {
    // Inicializar o cliente
    accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
    authToken := os.Getenv("TWILIO_AUTH_TOKEN")
    fromNumber = os.Getenv("TWILIO_FROM_NUMBER")

    client = twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })
}

// SendSMS envia uma mensagem SMS
func SendSMS(to, message string) error {
    params := &twilioApi.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom(fromNumber)
    params.SetBody(message)

    _, err := client.Api.CreateMessage(params)
    return err
}
```

## Boas Práticas

1. **Configuração**: Use variáveis de ambiente para configurar os provedores.
2. **Tratamento de Erros**: Implemente tratamento adequado de erros ao utilizar os provedores.
3. **Logging**: Adicione logs para facilitar o diagnóstico de problemas.
4. **Abstração**: Mantenha a interface dos provedores simples e abstrata.
5. **Testes**: Escreva testes para seus provedores, incluindo mocks para os serviços externos.
6. **Documentação**: Documente como configurar e usar os provedores.
7. **Segurança**: Proteja informações sensíveis, como senhas e tokens de API.
