package srp

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type OrderRequest struct {
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	Amount    float64 `json:"amount"`
	Email     string  `json:"email"`
}

type (
	OrderHandler struct {
		Service *OrderService
	}

	OrderService struct {
		Repo     *OrderRepository
		Notifier *EmailNotifiere
	}

	OrderRepository struct {
		DB *sql.DB
	}

	EmailNotifiere struct{}
)

func (h *OrderHandler) ProcessCheckout(w http.ResponseWriter, r *http.Request) {

	var req OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.Service.Process(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order processed"))
}

func (s *OrderService) Process(req OrderRequest) error {

	if req.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	err := s.Repo.Save(req.UserID, req.Amount)
	if err != nil {
		return err
	}

	s.Notifier.Sendi(req.Email)
	return nil
}

func (r *OrderRepository) Save(userID int, amount float64) error {
	_, err := r.DB.Exec(
		"INSERT INTO orders (user_id, amount) VALUES ($1, $2)",
		userID,
		amount,
	)
	return err
}

func (*EmailNotifiere) Sendi(email string) {
	fmt.Printf("Enviando email de confirmação para %s...\n", email)
}
