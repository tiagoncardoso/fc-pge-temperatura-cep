package usecase

import (
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/internal/application/dto"
	"github.com/tiagoncardoso/fc-pge-temperatura-cep/pkg/http_request"
	"strconv"
	"strings"
)

type RequestZipData struct {
	zipDataApiUrl string
}

func NewRequestZipData(zipDataApiUrl string) *RequestZipData {
	return &RequestZipData{
		zipDataApiUrl: zipDataApiUrl,
	}
}

func (c *RequestZipData) Execute(zipCode int) (dto.ViaCepApiDto, error) {
	zipUrl := makeUrl(c.zipDataApiUrl, zipCode)
	zipData, err := http_request.HttpGetRequest[dto.ViaCepApiDto](zipUrl)

	// TODO: Adicionar tratamentos de erro para CEP inválido e CEP não encontrado separadas para que o usuário receba o retorno correto
	if err != nil {
		return dto.ViaCepApiDto{}, err
	}

	return zipData, nil
}

func makeUrl(zipCodeBaseUrl string, zipCode int) string {
	strZipCode := strconv.Itoa(zipCode)

	return strings.Replace(zipCodeBaseUrl, "{ZIP}", strZipCode, 1)
}
