package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Representa uma pessoa utilizando a aplicacao
type Usuario struct {
	ID          uint64       `json: "id"`
	Nome        string       `json: "nome"`
	Email       string       `json: "email"`
	Nick        string       `json: "nick"`
	CriadoEm    time.Time    `json: "criadoEm"`
	Seguidores  []Usuario    `json: seguidores`
	Seguindo    []Usuario    `json: seguindo`
	Publicacoes []Publicacao `json: "publicacoes"`
}

// Faz 4 requisições na API para montar o usuario
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("Erro ao buscar as publicações")
			}

			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// Chama a api para buscas os dados base do usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequsicaoComAutentificacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// Chama a api para buscas os seguidores do usuario
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequsicaoComAutentificacao(r, http.MethodGet, url, nil)

	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// Chama a api para buscas os usuarios seguidos por um  usuario
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequsicaoComAutentificacao(r, http.MethodGet, url, nil)

	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

// Chama a api para buscas as publicoes do usuario
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {

	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.ApiUrl, usuarioID)
	response, erro := requisicoes.FazerRequsicaoComAutentificacao(r, http.MethodGet, url, nil)

	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes

}
