package expandErro

import (
	"net/http"
)

type Lasquera struct {
	Mensagem string  `json:"mensagem"`
	Err      string  `json:"erro"`
	Code     int     `json:"code"`
	Causas   []Causa `json:"causas"`
}
type Causa struct {
	Campo    string `json:"campo"`
	Mensagem string `json:"mensagem"`
}

func (l *Lasquera) Error() string {
	return l.Mensagem
}

// Construtor de Erros Lasquera
func NewLasquera(msg, err string, cod int, causas []Causa) *Lasquera {
	return &Lasquera{
		Mensagem: msg,
		Err:      err,
		Code:     cod,
		Causas:   causas,
	}
}

// Erros de chamada de função do Roteador - BadRequest
func NewBadRequestError(message string) *Lasquera {
	return &Lasquera{
		Mensagem: message,
		Err:      "bad request",
		Code:     http.StatusBadRequest,
	}
}

// Erros de chamada de função do Roteador - BadRequest Validation
func NewBadRequestValidationError(message string, causes []Causa) *Lasquera {
	return &Lasquera{
		Mensagem: message,
		Err:      "bad request",
		Code:     http.StatusBadRequest,
		Causas:   causes,
	}
}

// Erros de chamada de função do Roteador - Internal Server Error
func NewInternalServerError(message string) *Lasquera {
	return &Lasquera{
		Mensagem: message,
		Err:      "internal server error",
		Code:     http.StatusInternalServerError,
	}
}

// Erros de chamada de função do Roteador - Not Found
func NewNotFoundError(message string) *Lasquera {
	return &Lasquera{
		Mensagem: message,
		Err:      "not found",
		Code:     http.StatusNotFound,
	}
}

// Erros de chamada de função do Roteador - BadRequest
func NewForbiddenError(message string) *Lasquera {
	return &Lasquera{
		Mensagem: message,
		Err:      "forbidden",
		Code:     http.StatusForbidden,
	}
}
