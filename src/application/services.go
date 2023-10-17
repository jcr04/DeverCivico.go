package application

import (
	"github.com/jcr04/DeverCivico.go/src/domain"
	"github.com/jcr04/DeverCivico.go/src/infrastructure"
)

type CidadaoService struct {
	repo infrastructure.CidadaoRepository
}

// NewCidadaoService é uma função construtora para CidadaoService.
func NewCidadaoService(repo infrastructure.CidadaoRepository) *CidadaoService {
	return &CidadaoService{repo: repo}
}

func (s *CidadaoService) CadastrarCidadao(cidadao domain.Cidadao) error {
	// Chamada ao método Cadastrar do repositório.
	err := s.repo.Cadastrar(cidadao)
	if err != nil {
		return err // Retornando o erro, se houver algum.
	}
	return nil
}
