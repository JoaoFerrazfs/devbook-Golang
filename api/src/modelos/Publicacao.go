package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publicação representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json: "id,omitempty"`
	Titulo    string    `json: "titulo,omitempty"`
	Conteudo  string    `json: "conteudo,omitempty"`
	AutorID   uint64    `json: "autorId, omitempty"`
	AutorNick string    `json: "autorNick, omitempty"`
	Curtidas  uint64    `json: "curtidas"`
	CriadaEm  time.Time `json: "criadaEm, omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar a publicacao recebida
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatorio e não pode estar em branco ")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatorio e não pode estar em branco ")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
