package main

import (
	"fmt"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/config"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/internal/application/usecase"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	findZip := usecase.NewRequestZipData(conf.ApiUrlZip)
	// TODO: Adicionar tratamento de CEP para converter sempre para 8 dígitos sem caracteres especiais
	zipData, err := findZip.Execute(74333110)

	fmt.Println("Resultado de Busca CEP:", zipData)
}
