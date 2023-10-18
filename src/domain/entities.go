// domain/entities.go
package domain

type Cidadao struct {
	ID    int
	Nome  string
	Email string
	Senha string
}

type ProblemaReportado struct {
	ID          int
	Descricao   string
	Localizacao string
	DataHora    string
	Status      string
}

type Discussao struct {
	ID            int
	Titulo        string
	Descricao     string
	DataHora      string
	Participantes []Cidadao
}

type InformacoesGovernamentais struct {
	ID        int
	Titulo    string
	Descricao string
	DataHora  string
}

type Credentials struct {
	Email string `json:"email"`
	Senha string `json:"Senha"`
}
