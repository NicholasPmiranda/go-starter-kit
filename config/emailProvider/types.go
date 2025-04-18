package emailprovider

// EmailAttachment representa um anexo do email
type EmailAttachment struct {
	Filename string
	Path     string
}

// EmailMessage representa a configuração completa de um email
type EmailMessage struct {
	To           []string
	Cc           []string `json:"cc,omitempty"`
	Bcc          []string `json:"bcc,omitempty"`
	Subject      string
	Template     string
	TemplateData interface{}       `json:"templateData,omitempty"`
	Attachments  []EmailAttachment `json:"attachments,omitempty"`
}
