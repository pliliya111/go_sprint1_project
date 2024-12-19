package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pliliya111/go_sprint1_project/internal/calculator"
)

type RequestBody struct {
	Expression string `json:"expression"`
}
type ResponseBody struct {
	Result float64 `json:"result"`
}
type ResponseErrorBody struct {
	Error string `json:"error"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := calculator.Calc(reqBody.Expression)
	if err != nil {
		response := ResponseErrorBody{Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := ResponseBody{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
