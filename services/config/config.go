package config

import (
	"errors"
	"fmt"
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
	LogOutput        string
	LogLevel         string
}

// Carrega as configurações no Objeto Config a partir do arquivo passado como parâmetro
func (conf *Config) CarregaConfig(file string) error {
	var err error

	// Definir valores iniciais a partir do arquivo config.ini
	inidata, err := ini.Load(file)
	if err != nil {
		err = errors.New(err.Error() + ". O arquivo \"config.ini\" existe? ")
		return err
	}
	// Define as configurações padrão
	conf.NomeFantasia = "Agenda"
	conf.HoraInicioAtende, _ = time.ParseDuration("8h")
	conf.DuracaoAtende, _ = time.ParseDuration("45m")
	conf.HoraFimAtende, _ = time.ParseDuration("18h")
	conf.HoraInterval, _ = time.ParseDuration("12h")
	conf.DuraInterval, _ = time.ParseDuration("1h")
	conf.DiasSemanaAtende, _ = string2VetorInt("1,2,3,4,5") // Seg-Qua
	conf.admSecret = "Agend@123"
	conf.Canais, _ = string2VetorString("EMAIL,WEB")
	conf.ArmazemDados = "MongoDB"
	conf.ArmazemHost = "localhost"
	conf.ArmazemPort = 27017
	conf.ArmazemDatabase = "Agenda"
	conf.ApiServerPort = 8080
	conf.LogOutput = "stdout"
	conf.LogLevel = "info"

	// Lewitur das Configurações do arquivo "config.ini"
	sectionAgenda := inidata.Section("Agenda")
	sectionConfig := inidata.Section("Config")
	sectionLogger := inidata.Section("Logger")
	conf.NomeFantasia = sectionAgenda.Key("NomeFantasia").String()
	if conf.HoraInicioAtende, err = time.ParseDuration(sectionAgenda.Key("HoraInicioAtende").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.DuracaoAtende, err = time.ParseDuration(sectionAgenda.Key("DuracaoAtende").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.HoraFimAtende, err = time.ParseDuration(sectionAgenda.Key("HoraFimAtende").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.HoraInterval, err = time.ParseDuration(sectionAgenda.Key("HoraInterval").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.DuraInterval, err = time.ParseDuration(sectionAgenda.Key("DuraInterval").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.DiasSemanaAtende, err = string2VetorInt(sectionAgenda.Key("DiasSemanaAtende").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.admSecret = sectionConfig.Key("AdmSecret").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.Canais, err = string2VetorString(sectionConfig.Key("Canais").String()); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemDados = sectionConfig.Key("ArmazemDados").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemHost = sectionConfig.Key("Host").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemPort, err = sectionConfig.Key("Port").Int(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemDatabase = sectionConfig.Key("Database").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemUser = sectionConfig.Key("User").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ArmazemPassword = sectionConfig.Key("Password").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.ApiServerPort, err = sectionConfig.Key("ApiServerPort").Int(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.LogOutput = sectionLogger.Key("LogOutput").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	if conf.LogLevel = sectionLogger.Key("LogLevel").String(); err != nil {
		return errors.New("Leitura das Configurações: " + err.Error())
	}
	return nil
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
					return nil, errors.New("Valor de dia da semana repetido" + fmt.Sprint(num))
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
	// Checar os nomes reservados para os canis: "WAPP,EMAIL,VOZ,SMS,WEB"
	for _, s := range strSlice {
		if s != "WAPP" && s != "EMAIL" && s != "VOZ" && s != "SMS" && s != "WEB" {
			return nil, errors.New("Canal \"" + s + "\" inválido. Deve ser um dos tipos: WAPP,EMAIL,VOZ,SMS,WEB")
		}
	}
	return strSlice, nil
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
