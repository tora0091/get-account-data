package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tora0091/get-account-data/domain/repositories"
	"github.com/tora0091/get-account-data/domain/results"
)

type AccountHandler struct {
	AccountRepo *repositories.AccountRepository
}

func NewAccountHandler(ar *repositories.AccountRepository) *AccountHandler {
	return &AccountHandler{AccountRepo: ar}
}

func (h *AccountHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accounts, err := h.AccountRepo.FindAll(100)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	num := len(accounts)
	resp := make(map[string]interface{}, num)

	for _, account := range accounts {
		primaryKey := fmt.Sprintf("%d-%d-%s", account.PROPERTY_ID, account.CATEGORY_ID, account.START_DATE)
		resp[primaryKey] = map[string]interface{}{
			"plan_id":    account.PLAN_ID,
			"name":       account.NAME,
			"price":      account.PRICE,
			"start_date": account.START_DATE,
			"end_date":   account.END_DATE,
		}
	}

	b, _ := json.Marshal(results.AccountReuslt{
		Code: http.StatusOK,
		Body: resp,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *AccountHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	propertyId := r.FormValue("property_id")
	if len(propertyId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("property_id is not found."))
		return
	}

	pId, err := strconv.Atoi(propertyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("property_id : strconv a to i error."))
		return
	}

	accounts, err := h.AccountRepo.FindById(int64(pId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	num := len(accounts)
	resp := make(map[string]interface{}, num)

	for _, account := range accounts {
		primaryKey := fmt.Sprintf("%d-%d-%s", account.PROPERTY_ID, account.CATEGORY_ID, account.START_DATE)
		resp[primaryKey] = map[string]interface{}{
			"plan_id":    account.PLAN_ID,
			"name":       account.NAME,
			"price":      account.PRICE,
			"start_date": account.START_DATE,
			"end_date":   account.END_DATE,
		}
	}

	b, _ := json.Marshal(results.AccountReuslt{
		Code: http.StatusOK,
		Body: resp,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
