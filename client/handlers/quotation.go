package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"quotation-client/models"
	"quotation-client/utils"
	"time"
)

const QUOTATION_SERVER_URL string = "http://localhost:8080"

func GetQuotationHandler()  {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 300)
	defer cancel()

    res, err := GetQuotation()
    if err != nil {
        log.Printf("Erro ao buscar a cotação: %v", err)
    }

    data, err := utils.ConvertFromJson[models.Quotation](res)
    if err != nil {
        log.Printf("Erro ao fazer a conversão do JSON para a cotação: %v", err)
    }

	SaveQuotationInFile("Dólar: " + data.Bid)

    log.Println("A cotação do dólar é " + data.Bid)
    log.Println("O valor da cotação foi salvo com sucesso")
}

func GetQuotation() ([]byte, error) {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", QUOTATION_SERVER_URL + "/cotacao", nil)

	if err != nil {
        return nil, err
	}

    res, err := http.DefaultClient.Do(req)

    if err != nil {
        return nil, err
	}

	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func SaveQuotationInFile(content string) {
	utils.WriteInFile("cotacao.txt", content)
}