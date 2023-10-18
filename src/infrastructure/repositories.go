package infrastructure

import (
	"database/sql"
	"time"

	"github.com/jcr04/DeverCivico.go/src/domain"
)

type CidadaoRepository struct {
	db *sql.DB // conexão com o banco de dados
}

func NewCidadaoRepository(db *sql.DB) *CidadaoRepository {
	return &CidadaoRepository{db: db}
}

type ProblemaRepository struct {
	db *sql.DB
}

type DiscussaoRepository struct {
	db *sql.DB
}

func NewDiscussaoRepository(db *sql.DB) *DiscussaoRepository {
	return &DiscussaoRepository{db: db}
}

type InformacoesRepository struct {
	db *sql.DB
}

func NewInformacoesRepository(db *sql.DB) *InformacoesRepository {
	return &InformacoesRepository{db: db}
}

func NewProblemaRepository(db *sql.DB) *ProblemaRepository {
	return &ProblemaRepository{db: db}
}

func (r *CidadaoRepository) Cadastrar(cidadao *domain.Cidadao) error {
	query := `
        INSERT INTO cidadao (nome, email, senha)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	return r.db.QueryRow(query, cidadao.Nome, cidadao.Email, cidadao.Senha).Scan(&cidadao.ID)
}

func (r *ProblemaRepository) Reportar(problema *domain.ProblemaReportado) error {
	query := `
		INSERT INTO problema_reportado (descricao, localizacao, data_hora, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(query, problema.Descricao, problema.Localizacao, problema.DataHora, problema.Status).Scan(&problema.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DiscussaoRepository) Listar() ([]domain.Discussao, error) {
	query := `
		SELECT id, titulo, descricao, data_hora 
		FROM discussao
		ORDER BY data_hora DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discussoes []domain.Discussao
	for rows.Next() {
		var discussao domain.Discussao
		err := rows.Scan(&discussao.ID, &discussao.Titulo, &discussao.Descricao, &discussao.DataHora)
		if err != nil {
			return nil, err
		}
		discussoes = append(discussoes, discussao)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return discussoes, nil
}

func (r *DiscussaoRepository) Criar(discussao *domain.Discussao) error {
	query := `
		INSERT INTO discussao (titulo, descricao, data_hora)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	dataHora, err := time.Parse(time.RFC3339, discussao.DataHora)
	if err != nil {
		return err
	}
	err = r.db.QueryRow(query, discussao.Titulo, discussao.Descricao, dataHora).Scan(&discussao.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *InformacoesRepository) Obter() (domain.InformacoesGovernamentais, error) {
	query := `
		SELECT id, titulo, descricao, data_hora 
		FROM informacoes_governamentais
		ORDER BY data_hora DESC
		LIMIT 1
	`
	row := r.db.QueryRow(query)

	var info domain.InformacoesGovernamentais
	err := row.Scan(&info.ID, &info.Titulo, &info.Descricao, &info.DataHora)
	if err != nil {
		return domain.InformacoesGovernamentais{}, err
	}

	return info, nil
}

func (r *CidadaoRepository) ObterPorCredenciais(email, senha string) (domain.Cidadao, error) {
	query := `
        SELECT id, nome, email, senha
        FROM cidadao
        WHERE email = $1 AND senha = $2
    `
	row := r.db.QueryRow(query, email, senha)
	var cidadao domain.Cidadao
	err := row.Scan(&cidadao.ID, &cidadao.Nome, &cidadao.Email, &cidadao.Senha)
	if err != nil {
		return domain.Cidadao{}, err
	}
	return cidadao, nil
}
