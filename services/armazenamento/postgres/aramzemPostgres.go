package armazenamento

import "github.com/jmoiron/sqlx"

// Implementação da interface Database para PostgreSQL.
type PostgresDB struct {
	// campos específicos do PostgreSQL, se necessário
	Conn *sqlx.DB
}

func (p *PostgresDB) Connect() error {
	// lógica de conexão com PostgreSQL
	return nil // so para não ficar em erro
}

func (p *PostgresDB) Close() error {
	// lógica de fechamento de conexão
	return nil // so para não ficar em erro
}
