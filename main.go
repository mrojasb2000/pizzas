package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mrojasb2000/pizzas/api/models"
)

type Pizzas []models.Pizza

func (ps Pizzas) FindByID(ID int) (models.Pizza, error) {
	for _, pizza := range ps {
		if pizza.ID == ID {
			return pizza, nil
		}
	}
	return models.Pizza{}, fmt.Errorf("Couldn't find pizza with ID: %d", ID)
}

type Orders []models.Order

type pizzasHandler struct {
	pizzas *Pizzas
}

func (ph pizzasHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		if len(*ph.pizzas) == 0 {
			http.Error(w, "Error: No pizzas found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(ph.pizzas)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type ordersHandler struct {
	pizzas *Pizzas
	orders *Orders
}

func (oh ordersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var o models.Order

		if len(*oh.pizzas) == 0 {
			http.Error(w, "Error: No pizzas found", http.StatusNotFound)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&o)
		if err != nil {
			http.Error(w, "Can't decode body", http.StatusBadRequest)
			return
		}

		p, err := oh.pizzas.FindByID(o.PizzaID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
			return
		}

		o.Total = p.Price * o.Quantity
		*oh.orders = append(*oh.orders, o)
		json.NewEncoder(w).Encode(o)
	case http.MethodGet:
		json.NewEncoder(w).Encode(oh.orders)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	var orders Orders
	pizzas := Pizzas{
		models.Pizza{
			ID:    1,
			Name:  "Pepperoni",
			Price: 12,
		},
		models.Pizza{
			ID:    2,
			Name:  "Capricciosa",
			Price: 11,
		},
		models.Pizza{
			ID:    3,
			Name:  "Margherita",
			Price: 10,
		},
	}

	mux := http.NewServeMux()
	mux.Handle("/pizzas", pizzasHandler{&pizzas})
	mux.Handle("/orders", ordersHandler{&pizzas, &orders})

	log.Fatal(http.ListenAndServe(":9090", mux))
}
