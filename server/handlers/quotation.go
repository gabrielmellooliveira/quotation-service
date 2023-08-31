package handlers

import (
	"context"
	"io"
	"log"
	"net/http"
	"quotation-server/database"
	"quotation-server/models"
	"quotation-server/utils"
	"time"

	"github.com/google/uuid"
)

const AWESOME_API_URL string = "https://economia.awesomeapi.com.br"

func QuotationHandler(w http.ResponseWriter, r *http.Request) {
	res, err := GetQuotation()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Printf("Erro ao buscar a cotação: %v", err)
        w.Write([]byte("Erro ao buscar a cotação"))
        return
    }

	data, err := utils.ConvertFromJson[models.Quotation_USDBRL](res)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Printf("Erro ao fazer a conversão do JSON para a cotação: %v", err)
        w.Write([]byte("Erro ao fazer a conversão do JSON para a cotação"))
        return
    }

	err = SaveQuotation(data.UsdBrl)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Printf("Erro ao salvar cotação no banco de dados: %v", err)
        w.Write([]byte("Erro ao salvar cotação no banco de dados"))
        return
    }

	jsonResult, err := utils.ConvertToJson[models.Quotation](data.UsdBrl)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Printf("Erro ao fazer a conversão da cotação para JSON: %v", err)
        w.Write([]byte("Erro ao fazer a conversão da cotação para JSON"))
        return
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

func SaveQuotation(quotation models.Quotation) error {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 10)
    defer cancel()

    db, err := database.OpenConnection()

    if err != nil {
        return err
	}

    quotation.Id = uuid.New().String()

    err = db.WithContext(ctx).Create(quotation).Error

    if err != nil {
        return err
	}

    return nil
}

func GetQuotation() ([]byte, error) {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", AWESOME_API_URL + "/json/last/USD-BRL", nil)

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