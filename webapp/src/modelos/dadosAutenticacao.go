package modelos

// DadosAuteticacao contem os o ID e Token do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json: "id`
	Token string `json: "token"`
}
