package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (respositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := respositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ? ,?)",
	)

	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}
