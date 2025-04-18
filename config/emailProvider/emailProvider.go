package emailprovider

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var mailer *gomail.Dialer
var fromAddress string
var fromName string

func init() {
	// Carregar variáveis do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Ler configurações SMTP do arquivo .env
	host := os.Getenv("MAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Fatalf("Erro ao converter MAIL_PORT: %v", err)
	}
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	fromAddress = os.Getenv("MAIL_FROM_ADDRESS")
	fromName = os.Getenv("MAIL_FROM_NAME")

	// Configurar o Dialer do Gomail
	mailer = gomail.NewDialer(host, port, username, password)
}

// SendMail envia um e-mail usando um template HTML
func SendMail(emailMsg EmailMessage) error {
	// Ler e parsear o template
	tmpl, err := template.ParseFiles("template/" + emailMsg.Template + ".html")
	if err != nil {
		return fmt.Errorf("erro ao carregar template: %v", err)
	}

	// Executar o template com os dados fornecidos
	var body bytes.Buffer
	if err = tmpl.Execute(&body, emailMsg.TemplateData); err != nil {
		return fmt.Errorf("erro ao executar template: %v", err)
	}

	// Criar a mensagem
	msg := gomail.NewMessage()
	msg.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, fromAddress))

	// Configurar destinatários
	if len(emailMsg.To) == 0 {
		return fmt.Errorf("pelo menos um destinatário é necessário")
	}
	msg.SetHeader("To", emailMsg.To...)

	// Adicionar CC se houver
	if len(emailMsg.Cc) > 0 {
		msg.SetHeader("Cc", emailMsg.Cc...)
	}

	// Adicionar BCC se houver
	if len(emailMsg.Bcc) > 0 {
		msg.SetHeader("Bcc", emailMsg.Bcc...)
	}

	msg.SetHeader("Subject", emailMsg.Subject)
	msg.SetBody("text/html", body.String())

	// Adicionar anexos se houver
	for _, attachment := range emailMsg.Attachments {
		msg.Attach(attachment.Path, gomail.Rename(attachment.Filename))
	}

	// Enviar o email
	if err := mailer.DialAndSend(msg); err != nil {
		return fmt.Errorf("erro ao enviar email: %v", err)
	}

	return nil
}
