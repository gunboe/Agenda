package convControllers

import (
	"Agenda/common"
	"Agenda/models"
	"Agenda/services/armazenamento"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Constantes
const Convenio = "Convênio"

/////////////////
// CRUD Convenios
/////////////////

// (CREATE) Cria convênio e salva no armazém
func CriaConvenio(conv models.Convenio) error {
	// Verifica o Convenio
	err := models.ChecarConvenio(conv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return err
	}
	// Checa se já existe Convenio pelo Nome
	convs, err := armazenamento.GetConveniosByName(conv.NomeConv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return err
	}
	if convs == nil {
		// Checa se já existe Convênio pelo Nr Prestador
		fmt.Println("conv.NrPrestador:", conv.NrPrestador)
		c, err := armazenamento.GetConveniosByNrPrestador(conv.NrPrestador)
		if err == nil {
			err = errors.New("Novo Convênio:\"" + conv.NomeConv + "\" já existe com o mesmo Nr Prestador:" + c.NrPrestador + " mas com o Nome: " + c.NomeConv)
			fmt.Println("Erro:("+Convenio+")", err)
			return err
		}
		if c.ID.IsZero() {
			result, err := armazenamento.CreateConvenio(conv)
			if err != nil {
				fmt.Println("Erro:("+Convenio+")", err)
			} else {
				fmt.Println("Convenio Criado e armazenado:", result)
			}
		}
	} else {
		err = errors.New("Convênio:\"" + conv.NomeConv + "\" já existe com o mesmo nome sob o Nr Prestador:" + conv.NrPrestador)
		fmt.Println(err.Error())
	}
	return err
}

// (READ) Retorna um Vetor de Convenios passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func GetConveniosPorNome(conv string) []models.Convenio {
	convs, err := armazenamento.GetConveniosByName(conv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return nil
	}
	if convs == nil {
		fmt.Println("Erro: Convênio: " + conv + " não encontrado.")
		return nil
	}
	return convs
}

// (READ) Retorna um Convenio passando como parâmetro o "ID" do convênio.
// Se não encontrar retorna Convênio Zerado.
func GetConvenioPorId(id primitive.ObjectID) (models.Convenio, error) {
	conv, err := armazenamento.GetConvenioById(id)
	if err != nil {
		return models.Convenio{}, err
	}
	return conv, nil
}

// (list) Retorna Lista de Convenios no formato "json" ou "bson" passando como parâmetro o "Nome" do convênio.
// Se o argumento "nome" = "*", retornará todos os convênios armazenados.
func ListaConvenio(nome string, formato ...string) interface{} {
	convs, err := armazenamento.GetConveniosByName(nome)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return nil
	} else {
		// Se houver "formato" e do tipo "bson", imprima neste.
		if len(formato) > 0 && strings.EqualFold(formato[0], "bson") {
			fmt.Println("listando Convênios em bson")
			return convs
			// Caso contrário, use por padrão "Json"
		} else {
			fmt.Println("listando Convênios em json")
			return common.PrintJSON(convs)
		}
	}
}

// (UPDATE) Atualiza os Dados de um ou mais Convênio armazenado utilizando como parâmetro o Nome do Convênio("nome"),
// o Struct do Novo Convênio("novoConv") e a opção de alterar Todos("todos") simultaneamente.
// Essa função NÃO checa os valores, LOGO NÃO DEVE SER USADA NA PRODUÇÃO. Utilize "porID".
func AtualizaConvPorNome(nome string, novoConv models.Convenio, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		var err error
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
		} else {
			// Atualiza os dados do Convênio
			result, err := armazenamento.UpdateConvenioByName(nome, novoConv, todos)
			if err != nil {
				fmt.Println("Erro:("+Convenio+")", err)
			} else {
				fmt.Println("Convenios encontrados:", result.MatchedCount)
				fmt.Println("Convenios atualizados:", result.ModifiedCount)
			}
		}
	}
}

// (UPDATE) Atualiza os Dados de um Convênio armazenado utilizando como parâmetro o ID do Convênio
// e um novo Convênio com TODOS os atributos para serem checados antes de atualizados no Armazem.
func AtualizaConvPorId(id primitive.ObjectID, novoConv models.Convenio) {
	var err error
	// Testa as alterações estão em conformidade
	err = models.ChecarConvenio(novoConv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
		return
	}
	// Atualiza os dados do Convênio
	result, err := armazenamento.UpdateConvenioById(id, novoConv)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
	} else if result.ModifiedCount > 0 {
		fmt.Println("Convenio atualizado:", id.String())
	} else {
		fmt.Println("Convênio:\"" + id.String() + "\" não foi alterado no Armazém.")
	}
}

// (UPDATE) Disponibilizar um Convênio por ID. Caso um Convênio esteja marcado como Indisponível,
// essa função o torna Disponível novamente para alteração de dados ou uso em PlanosPgtos.
func HabiliteConvPorId(id primitive.ObjectID, b bool) {
	result, err := armazenamento.AllowConveioById(id, b)
	if err != nil {
		fmt.Println("Erro:("+Convenio+")", err)
	} else if result.MatchedCount == 0 {
		fmt.Println("Erro: Convênio não encontrado.")
	} else {
		if b {
			fmt.Println("Convênio Disponibilizado.")
		} else {
			fmt.Println("Convênio Indisponibilizado.")
		}
	}
}

// (DELETE) Deleta um Convênio específico ou mais de um utilizando o Nome do Convênio como parâmetro de busca.
// Para Deletar todos os Convênios da busca é possível utilizar o parâmetro Boleano "todos".
func DeletaConveniosPorNome(nome string, todos bool) {
	// Checa se o Nome do Convenio está vazio
	if nome == "" {
		fmt.Println("Erro: Nome do Convênio nulo/vazio.")
	} else {
		result, err := armazenamento.DeleteConvenioByName(nome, todos)
		if err != nil {
			fmt.Println("Erro:("+Convenio+")", err)
			fmt.Println("Provavel que o Convênio:\"" + nome + "\" não exista no Armazém.")
		} else {
			fmt.Println("Convenios deletados:", result.DeletedCount)
		}
	}
}

// (DELETE) Deleta um Convênio específico utilizando o ID do Convênio como parâmetro de busca.
func DeletaConvenioPorId(id primitive.ObjectID) error {
	var err error
	// Checa se o Nome do Convenio está vazio
	if id.IsZero() {
		fmt.Println("iD nulo/vazio")
	} else {
		result, err := armazenamento.DeleteConvenioById(id)
		if err == nil {
			if result.DeletedCount == 0 {
				err = errors.New("convenio não encontrado")
				fmt.Println(err)
				return err
			} else {
				fmt.Println(result.DeletedCount, "convenio deletado")
				return nil
			}
		}
	}
	return err
}
