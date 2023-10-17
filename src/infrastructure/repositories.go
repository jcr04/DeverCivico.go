package infrastructure

import (
	"database/sql"

	"github.com/jcr04/DeverCivico.go/src/domain"
)

type CidadaoRepository struct {
	db *sql.DB // conexão com o banco de dados
}

// NewCidadaoRepository é uma função construtora para CidadaoRepository.
func NewCidadaoRepository(db *sql.DB) *CidadaoRepository {
	return &CidadaoRepository{db: db}
}

func (r *CidadaoRepository) Cadastrar(cidadao domain.Cidadao) error {
	query := `
		INSERT INTO cidadao (nome, email, senha)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(query, cidadao.Nome, cidadao.Email, cidadao.Senha).Scan(&cidadao.ID)
	if err != nil {
		return err
	}

	return nil
}
