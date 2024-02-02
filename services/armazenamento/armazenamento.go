package armazenamento

// Interface que define os métodos comuns para operações de banco de dados.
type Database interface {
	Connect() error
	Close() error
}
