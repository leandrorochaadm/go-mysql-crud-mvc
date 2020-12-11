package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa o modelo de usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Validar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Validar(etapa string) error {
	if erro := usuario.verificarCampoObrigatorio(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) verificarCampoObrigatorio(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode está em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode está em branco")
	}

	if usuario.Senha == "" && etapa == "cadastro" {
		return errors.New("A senha é obrigatório e não pode está em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
