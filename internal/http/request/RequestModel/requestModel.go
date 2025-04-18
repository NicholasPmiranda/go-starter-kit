package RequestModel

type Pessoa struct {
	Nome  string `json:"nome,omitempty"`
	Email string `json:"email,omitempty"`
	Idade int    `json:"idade,omitempty"`
}
