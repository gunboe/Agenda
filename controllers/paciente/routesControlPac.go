package pacControllers

import (
	"Agenda/common"
	"Agenda/controllers"
	"Agenda/models"
	"Agenda/services/expandErro"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/////////////////////////////////
// Routes Control para Pacientes
/////////////////////////////////

// Alteração do Password(Secret) do Paciente
func (pacFunc *PacienteFunc) ChangePassword(c *gin.Context) {
	// Cria o objeto para receber os atributos do Login
	var pacSecret controllers.PacienteSecret
	// Avalia os Atributos do Login de Paciente vindos de: "c *gin.Context"
	if err := controllers.AvaliarRequest(c, &pacSecret); err != nil {
		return
	}
	email := pacSecret.Email
	secretPac := common.MD5(pacSecret.Secret)
	// Procura Paciente por email
	pac, err := pacFunc.GetPacientePorEmailSecret(email, secretPac)
	if err != nil {
		reqErro := expandErro.NewForbiddenError("email ou secret incorreto")
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Altera o Secret(Password) do Paciente
	err = pacFunc.ChangePasswordPacPorId(pac.ID, common.MD5(pacSecret.ConfirmSecret))
	if err != nil {
		reqErro := expandErro.NewInternalServerError("Erro: alteração do Password(Secret) do Paciente")
		c.JSON(reqErro.Code, reqErro)
		return
	}
	c.JSON(http.StatusOK, pac)
}

// Login Paciente
func (pacFunc *PacienteFunc) LoginPaciente(c *gin.Context) {
	// Cria o objeto para receber os atributos do Login
	var pacLogin controllers.PacienteLogin
	// Avalia os Atributos do Login de Paciente vindos de: "c *gin.Context"
	if err := controllers.AvaliarRequest(c, &pacLogin); err != nil {
		return
	}
	email := pacLogin.Email
	secretPac := common.MD5(pacLogin.Secret)
	// Procura Paciente por email
	pac, err := pacFunc.GetPacientePorEmailSecret(email, secretPac)
	if err != nil {
		reqErro := expandErro.NewForbiddenError("email ou secret incorreto")
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Geração do Token
	// Obtem o Secret da configuração
	secretConf, err := pacFunc.Config.GetSecret()
	if err != nil {
		reqErro := expandErro.NewInternalServerError("secret não definida")
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Cria os atributos do Token
	claims := jwt.MapClaims{
		"id":      pac.ID,
		"email":   pac.Email,
		"nome":    pac.Nome,
		"expData": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Gera o Hash do Token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secretConf))
	// Adicionao Token no Header da Resposta
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, pac)
}

// Cria Paciente por Json
func (pacFunc *PacienteFunc) CreatePac(c *gin.Context) {
	// Avalia os atributos do Request de Paciente
	var pacRequest models.Paciente
	var err error
	if err = controllers.AvaliarRequest(c, &pacRequest); err != nil {
		return
	}
	// Cria o Paciente se não houver erros legados(checagem de campo) ou Erros de Negocio
	result, err := pacFunc.CriaPaciente(pacRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na regra de negócio: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, result)
}

// Insere Plano de Pagamento em um Paciente passando um ID e o PlanoPgto
func (pacFunc *PacienteFunc) InserePlanoPac(c *gin.Context) {
	var planoRequest models.PlanoPgto
	var err error
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &planoRequest); err != nil {
		return
	}
	// Obtem o ID e verifica se ID formatado corretamente
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na inserção do Plano de Pagamento:(pacId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Checa Erros de Negocio
	result, err := pacFunc.InsPlanoPgtoPaciente(id, planoRequest)
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erros na regra de negócio: " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, result)
}

// RC: Retornar um objeto Json do Paciente por ID
func (pacFunc *PacienteFunc) FindPacById(c *gin.Context) {
	// Checa o Id recebido
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(pacId) " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	pac, err := pacFunc.GetPacientePorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	fmt.Println("paciente encontrado com id:", id)
	c.JSON(http.StatusOK, pac)
}

// RC: Retornar TODOS ("*") objetos Json do Paciente ou por Nome
func (pacFunc *PacienteFunc) FindPacientes(c *gin.Context) {
	// Tenta realizar a busca
	pacs := pacFunc.ListaPaciente(c.Param("nome"), "bson")
	jsize := len(pacs.([]models.Paciente))
	if jsize == 0 {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: pacientes não encontrados")

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Converte o Tipo Interface no Tipo dos dados reais e calucula o tamanho do array
	fmt.Println("Paciente(s) encontrado(s):", jsize)
	c.JSON(http.StatusOK, pacs)
}

// RC: Retornar um objeto Json do Paciente por CPF
func (pacFunc *PacienteFunc) FindPacByCPF(c *gin.Context) {
	var err error
	cpf := c.Param("pacCPF")
	// Checa o CPF recebido
	if _, ok := common.CPFvalido(cpf); !ok {
		err = errors.New("CPF inválido")
		reqErro := expandErro.NewBadRequestError("Erro na pesquisa:(pacCPF) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a busca
	pac, err := pacFunc.GetPacientePorCPF(cpf)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na pesquisa: " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	fmt.Println("Paciente encontrado com CPF:", cpf)
	c.JSON(http.StatusOK, pac)
}

// RC: Atualiza um Paciente por Json
func (pacFunc *PacienteFunc) UpdatePac(c *gin.Context) {
	var pacRequest models.Paciente
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na atualização do Paciente: (pacId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Realiza o Marshal dos Campos da requição no Objeto
	if err = controllers.AvaliarRequest(c, &pacRequest); err != nil {
		return
	}
	// Atualiza o Objeto
	err = pacFunc.AtualizaPacPorId(id, pacRequest)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro na atualização do Paciente (regras de negócio): " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, id)
}

// RC: (Des)Bloqueia um Paciente por Query
func (pacFunc *PacienteFunc) BloqPac(c *gin.Context) {
	var reqErro *expandErro.Lasquera
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro no (Des)Bloqueio do Paciente: (pacId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Obter o valor do atributo "disponivel"
	s := c.Query("bloqueado")
	b, err := strconv.ParseBool(s)
	if err != nil {
		if s == "" {
			reqErro = expandErro.NewBadRequestError("Erro no Bloqueio: Query (bloqueado) não informado")
		} else {
			reqErro = expandErro.NewBadRequestError("Erro no Bloqueio: Query (bloqueado) deve ser true ou false, recebido: " + s)
		}
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a alterção do atributo
	err = pacFunc.HabilitePacPorId(id, b)
	if err != nil {
		reqErro = expandErro.NewBadRequestError("Erro no (Des)Bloqueio do Paciente: (regras de negócio): " + err.Error())
		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, id)
}

// RC: Deleta um Paciente por ID
func (pacFunc *PacienteFunc) DeletePacById(c *gin.Context) {
	var err error
	// Checa ID
	id, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na deleção: (pacId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a deleção
	err = pacFunc.DeletaPacientePorId(id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na deleção: " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, id)
}

// RC: Deleta Plano de Pagamento por ID (PlanoPgto.ID)
func (pacFunc *PacienteFunc) DelPlanoPac(c *gin.Context) {
	var err error
	// Checa ID do Paciente
	pacid, err := primitive.ObjectIDFromHex(c.Param("pacId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na deleção: (pacId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Checa ID do PlanoPgto
	id, err := primitive.ObjectIDFromHex(c.Param("planoId"))
	if err != nil {
		reqErro := expandErro.NewBadRequestError("Erro na deleção: (planoId) " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Tenta realizar a deleção
	err = pacFunc.DeletaPlanoPorId(pacid, id)
	if err != nil {
		reqErro := expandErro.NewNotFoundError("Erro na deleção: " + err.Error())

		c.JSON(reqErro.Code, reqErro)
		return
	}
	// Processa e envia resposta (Qdo é para criação é necessário obter o ID do resultado do Mongo)
	controllers.Resposta(c, id)
}
