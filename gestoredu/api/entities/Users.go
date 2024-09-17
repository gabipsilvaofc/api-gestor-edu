package entities

type Users struct {
	ID             string `json:"id,omitempty"`
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Telefone       string `json:"telefone"`
	Email          string `json:"email"`
	Sexo           string `json:"sexo"`
	Idade          string `json:"idade"`
	CPF            string `json:"cpf"`
	Endereco       string `json:"endereco"`
	DataNascimento string `json:"dataNascimento"`
	Matricula      string `json:"matricula"`
}

type UserRequest struct {
	// User's name (required, minimum of 4 characters, maximum of 100 characters).
	// @json
	// @jsonTag name
	// @binding required,min=4,max=100
	Nome string `json:"nome" binding:"required,min=4,max=100"`

	// User's sobrenome (required, minimum of 4 characters, maximum of 100 characters).
	// @json
	// @jsonTag sobrenome
	// @binding required,min=4,max=100
	Sobrenome string `json:"sobrenome" binding:"required,min=4,max=100"`

	// User's telefone.
	// @json
	// @jsonTag telefone
	// @jsonExample 48988881111
	Telefone string `json:"telefone" binding:"required,min=8,max=11" example:"48988881111"`

	// User's email (required and must be a valid email address).
	// Example: user@example.com
	// @json
	// @jsonTag email
	// @jsonExample user@example.com
	// @binding required,email
	Email string `json:"email" binding:"required,email" example:"test@test.com"`

	// User's
	// @json
	// @jsonTag sexo
	// @jsonExample F
	// @binding required,sexo
	Sexo string `json:"sexo" binding:"required,M,F" example:"F"`

	// User's idade (required, must be between 1 and 100).
	// @json
	// @jsonTag idade
	// @jsonExample 30
	Idade string `json:"idade" binding:"required,min=1,max=100" example:"30"`

	// User's CPF.
	// @json
	// @jsonTag cpf
	// @jsonExample 000.000.000-00
	CPF string `json:"cpf" binding:"required,min=11,max=14" example:"000.000.000-00"`

	// User's endereco.
	// @json
	// @jsonTag endereco
	// @jsonExample Rua Francisco Matos, 131
	// @binding required,min=3,max=100
	Endereco string `json:"endereco" binding:"required,min=3,max=100"`

	// User's dataNascimento (required, must be between 1 and 100).
	// @json
	// @jsonTag dataNascimento
	// @jsonExample 30/03/2003
	DataNascimento string `json:"dataNascimento" binding:"required,min=8,max=10" example:"30/03/2003"`

	// User's matricula.
	// @json
	// @jsonTag matricula
	// @jsonExample 012345678
	Matricula string `json:"matricula" binding:"required,min=8,max=10" example:"012345678"`
}
