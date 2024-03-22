// Command http
package main

import (
	json "encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jhpinder/order-up-go/internal"
	log "github.com/sirupsen/logrus"
)

var menuItems = [3]menuItem{
	{
		ID:            uuid.New(),
		DisplayName:   "Mac 'n Cheese",
		AmountInStock: 6,
		ImageURL:      "https://en.wikipedia.org/wiki/File:Original_Mac_n_Cheese_.jpg",
	},
	{
		ID:            uuid.New(),
		DisplayName:   "Fried Chicken Drumstick",
		AmountInStock: 5,
		ImageURL:      "https://as1.ftcdn.net/v2/jpg/05/48/91/86/1000_F_548918636_WSoVnzQaHxvHyEqgYBsB27s1cRM8T64c.webp",
	},
	{
		ID:            uuid.New(),
		DisplayName:   "Sweet Tea",
		AmountInStock: 8,
		ImageURL:      "https://as1.ftcdn.net/v2/jpg/02/71/57/10/1000_F_271571019_jAbFG0INaOJ1nonSdf5005EYYIaaxt3W.jpg",
	},
}

type menuItem struct {
	DisplayName   string    `json:"displayName"`
	ImageURL      string    `json:"imageURL"`
	ID            uuid.UUID `json:"_id"`
	AmountInStock int       `json:"amountInStock"`
}

func init() {
	internal.SetupLogging()
}

func main() {
	http.HandleFunc("/menu", handleMenu)
	server := http.Server{
		ReadTimeout: time.Minute,
	}

	log.Info(menuItems[0].DisplayName)
	log.Info(menuItems[0].ID)
	log.Info(menuItems[0].AmountInStock)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleMenu(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Info("Received GET /menu")

	err := json.NewEncoder(w).Encode(menuItems)
	if err != nil {
		log.Error(err)
	}
}
