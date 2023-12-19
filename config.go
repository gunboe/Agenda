package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

type config struct {
	NomeFantasia     string
	HoraInicioAtende time.Duration
	DuracaoAtende    time.Duration
	HoraFimAtende    time.Duration
	HoraInterval     time.Duration
	DuraInterval     time.Duration
	DiasSemanaAtende []int
	admSecret        string
}

func (c *config) SetSecret(s string) error {
	if s == "" {
		return errors.New("[SetSecret] Segredo Nulo ou vazio!")
	} else {
		c.admSecret = s
		return nil
	}
}
func (c *config) GetSecret() (string, error) {
	if c.admSecret == "" {
		return "", errors.New("[GetSecret] Segredo Nulo ou vazio!")
	} else {
		return c.admSecret, nil
	}
}

func init() {
	// Testes iniciais do programa e verificação de requisitos
	// Por exemplo:
	// - Testar no config.ini os dias da semana de 0 à 6
	// - Lançar os Erros como Logs
}

func (conf *config) carregaConfig(file string) {

	// Definir valores iniciais a partir do arquivo config.ini
	inidata, err := ini.Load(file)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	section := inidata.Section("config")
	conf.NomeFantasia = section.Key("NomeFantasia").String()
	conf.HoraInicioAtende, _ = time.ParseDuration(section.Key("HoraInicioAtende").String())
	conf.DuracaoAtende, _ = time.ParseDuration(section.Key("DuracaoAtende").String())
	conf.HoraFimAtende, _ = time.ParseDuration(section.Key("HoraFimAtende").String())
	conf.HoraInterval, _ = time.ParseDuration(section.Key("HoraInterval").String())
	conf.DuraInterval, _ = time.ParseDuration(section.Key("DuraInterval").String())
	conf.DiasSemanaAtende, err = string2Int(section.Key("DiasSemanaAtende").String())
	if err != nil {
		fmt.Println(err)
	}
	conf.admSecret = section.Key("admSecret").String()
}

func string2Int(str string) ([]int, error) {
	// Divide a string usando a função Split do pacote strings
	strSlice := strings.Split(str, ",")
	// Cria um slice para armazenar os inteiros
	intSlice := make([]int, len(strSlice))
	// Converte os strings para inteiros usando a função Atoi do pacote strconv
	for i, s := range strSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, errors.New("Erro ao converter string para inteiro:" + err.Error())
		} else if num < 0 || num > 6 {
			return nil, errors.New("Dia da semana invalido, deve ser entre 0 e 6 (Dom..Sab):")
		} else {
			// Verifica valores do dia da semana repetidos
			for j := 0; j < len(intSlice); j++ {
				if intSlice[j] == num {
					return nil, errors.New("Erro: Valor de dia da semana repetido" + string(num))
				}
			}
			intSlice[i] = num
		}
	}
	return intSlice, nil
}
