package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func exec(db *sql.DB, sql string) sql.Result {
	result, erro := db.Exec(sql)
	if erro != nil {
		panic(erro)
	}
	return result
}

func criarBanco() {
	db, erro := sql.Open("mysql", config.StringConexaoSemBanco)
	if erro != nil {
		panic(erro)
	}

	exec(db, "create database if not exists usuarios")
	exec(db, "use usuarios")
	exec(db,
		`create table if not exists usuarios(
		id int auto_increment primary key,
		nome varchar(50) not null,
		email varchar(30) not null unique,
		senha varchar(64) not null,
		criadoEm timestamp default current_timestamp()
	) ENGINE=INNODB; `)
}

// Conectar abre a conexao com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	// criarBanco()

	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
