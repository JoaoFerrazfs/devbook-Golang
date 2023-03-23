package modelos

// DadosAtuteticacao contem o token e id do usuario
type DadosAutenticacao struct {
	ID    string `json: "id"`
	Token string `json: "token"`
}
