package application

import (
	"github.com/jcr04/DeverCivico.go/src/domain"
	"github.com/jcr04/DeverCivico.go/src/infrastructure"
)

type CidadaoService struct {
	repo infrastructure.CidadaoRepository
}

type ProblemaService struct {
	repo infrastructure.ProblemaRepository
}

func NewProblemaService(repo infrastructure.ProblemaRepository) *ProblemaService {
	return &ProblemaService{repo: repo}
}

type DiscussaoService struct {
	repo infrastructure.DiscussaoRepository
}

func NewDiscussaoService(repo infrastructure.DiscussaoRepository) *DiscussaoService {
	return &DiscussaoService{repo: repo}
}

type InformacoesService struct {
	repo infrastructure.InformacoesRepository
}

func NewInformacoesService(repo infrastructure.InformacoesRepository) *InformacoesService {
	return &InformacoesService{repo: repo}
}

func NewCidadaoService(repo infrastructure.CidadaoRepository) *CidadaoService {
	return &CidadaoService{repo: repo}
}

func (s *CidadaoService) CadastrarCidadao(cidadao *domain.Cidadao) error {
	return s.repo.Cadastrar(cidadao)
}

func (s *ProblemaService) ReportarProblema(problema *domain.ProblemaReportado) error {
	err := s.repo.Reportar(problema)
	if err != nil {
		return err
	}
	return nil
}

func (s *DiscussaoService) CriarDiscussao(discussao domain.Discussao) error {
	return s.repo.Criar(discussao)
}

func (s *DiscussaoService) ListarDiscussoes() ([]domain.Discussao, error) {
	return s.repo.Listar()
}

func (s *InformacoesService) ObterInformacoes() (domain.InformacoesGovernamentais, error) {
	return s.repo.Obter()
}

func (s *CidadaoService) Login(credentials domain.Credentials) (domain.Cidadao, error) {
	return s.repo.ObterPorCredenciais(credentials.Email, credentials.Senha)
}
