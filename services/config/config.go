package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

type Config struct {
	NomeFantasia     string
	HoraInicioAtende time.Duration
	DuracaoAtende    time.Duration
	HoraFimAtende    time.Duration
	HoraInterval     time.Duration
	DuraInterval     time.Duration
	DiasSemanaAtende []int
	admSecret        string
	Canais           []string
	ArmazemDados     string
	ArmazemHost      string
	ArmazemPort      int
	ArmazemUser      string
	ArmazemPassword  string
	ArmazemDatabase  string
	ArmazemCert      string
	ArmazemCa        string
	ArmazemChave     string
	ArmazemExtra     string
	ApiServerPort    int
}

var ConfigInicial Config

func init() {
	ConfigInicial.CarregaConfig("config.ini")
}

func (c *Config) SetSecret(s string) error {
	if s == "" {
		return errors.New("[SetSecret] Segredo Nulo ou vazio!")
	} else {
		c.admSecret = s
		return nil
	}
}

func (c *Config) GetSecret() (string, error) {
	if c.admSecret == "" {
		return "", errors.New("[GetSecret] Segredo Nulo ou vazio!")
	} else {
		return c.admSecret, nil
	}
}

func (conf *Config) CarregaConfig(file string) {

	// Definir valores iniciais a partir do arquivo config.ini
	inidata, err := ini.Load(file)
	if err != nil {
		fmt.Printf("Erro: \"%v\". O arquivo \"config.ini\" existe? ", err)
		os.Exit(1)
	}
	// TODO: Tratar os erros abaixo
	sectionAgenda := inidata.Section("Agenda")
	sectionConfig := inidata.Section("Config")
	conf.NomeFantasia = sectionAgenda.Key("NomeFantasia").String()
	conf.HoraInicioAtende, _ = time.ParseDuration(sectionAgenda.Key("HoraInicioAtende").String())
	conf.DuracaoAtende, _ = time.ParseDuration(sectionAgenda.Key("DuracaoAtende").String())
	conf.HoraFimAtende, _ = time.ParseDuration(sectionAgenda.Key("HoraFimAtende").String())
	conf.HoraInterval, _ = time.ParseDuration(sectionAgenda.Key("HoraInterval").String())
	conf.DuraInterval, _ = time.ParseDuration(sectionAgenda.Key("DuraInterval").String())
	conf.DiasSemanaAtende, err = string2VetorInt(sectionAgenda.Key("DiasSemanaAtende").String())
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
	conf.admSecret = sectionConfig.Key("admSecret").String()
	conf.Canais, err = string2VetorString(sectionConfig.Key("canais").String())
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
	conf.ArmazemDados = sectionConfig.Key("armazemDados").String()
	conf.ArmazemHost = sectionConfig.Key("host").String()
	conf.ArmazemPort, _ = sectionConfig.Key("port").Int()
	conf.ArmazemDatabase = sectionConfig.Key("database").String()
	conf.ApiServerPort, _ = sectionConfig.Key("apiServerPort").Int()
}

func string2VetorInt(str string) ([]int, error) {
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
					return nil, errors.New("Erro: Valor de dia da semana repetido" + fmt.Sprint(num))
				}
			}
			intSlice[i] = num
		}
	}
	return intSlice, nil
}

func string2VetorString(str string) ([]string, error) {
	// Divide a string usando a função Split do pacote strings
	strSlice := strings.Split(str, ",")
	// Cchecar os nomes reservados para os canis: "WAPP,EMAIL,VOZ,SMS,WEB"
	// TODO: Esses canais devem ser uma lista configurada externamente com seus parametros
	//       e não fixa no código abaixo
	for _, s := range strSlice {
		if s != "WAPP" && s != "EMAIL" && s != "VOZ" && s != "SMS" && s != "WEB" {
			return nil, errors.New("Erro: Canal \"" + s + "\" inválido. Deve ser um dos tipos: WAPP,EMAIL,VOZ,SMS,WEB")
		}
	}
	return strSlice, nil
}
