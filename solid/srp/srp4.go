package srp

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	ValueInDollar float64 `json:"valor_dolar"`
}

type CurrencyService struct {
	Cotacao float64
}

func (c CurrencyService) Convert(value float64) float64 {
	return value * c.Cotacao
}

func ConverterMoedaHandler(service CurrencyService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var input Input

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Erro no JSON", http.StatusBadRequest)
			return
		}

		valueInReais := service.Convert(input.ValueInDollar)

		response := map[string]float64{
			"valor_reais": valueInReais,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {

	service := CurrencyService{
		Cotacao: 5.50,
	}

	http.HandleFunc("/converter", ConverterMoedaHandler(service))
	http.ListenAndServe(":8080", nil)
}
