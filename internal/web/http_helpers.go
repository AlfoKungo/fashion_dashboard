package web

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const maxAmount = 50

func ParseAmount(r *http.Request, def int) (int, error) {
	raw := r.URL.Query().Get("amount")
	if raw == "" {
		return def, nil
	}
	amount, err := strconv.Atoi(raw)
	if err != nil || amount < 1 || amount > maxAmount {
		return 0, errInvalidAmount
	}
	return amount, nil
}

type invalidAmountError struct{}

func (invalidAmountError) Error() string { return "amount must be an integer between 1 and 50" }

var errInvalidAmount invalidAmountError

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
