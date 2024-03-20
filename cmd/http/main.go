// Command http
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jhpinder/order-up-go/internal"
	log "github.com/sirupsen/logrus"
)

var menuItems = [3]menuItem{
	{
		id:            uuid.New(),
		displayName:   "Mac 'n Cheese",
		amountInStock: 6,
		imageURL:      "https://en.wikipedia.org/wiki/File:Original_Mac_n_Cheese_.jpg",
	},
	{
		id:            uuid.New(),
		displayName:   "Fried Chicken Drumstick",
		amountInStock: 5,
		imageURL:      "https://as1.ftcdn.net/v2/jpg/05/48/91/86/1000_F_548918636_WSoVnzQaHxvHyEqgYBsB27s1cRM8T64c.webp",
	},
	{
		id:            uuid.New(),
		displayName:   "Sweet Tea",
		amountInStock: 8,
		imageURL:      "https://as1.ftcdn.net/v2/jpg/02/71/57/10/1000_F_271571019_jAbFG0INaOJ1nonSdf5005EYYIaaxt3W.jpg",
	},
}

type menuItem struct {
	displayName   string
	imageURL      string
	id            uuid.UUID
	amountInStock int
}

func init() {
	internal.SetupLogging()
}

func main() {
	http.HandleFunc("/menu", handleMenu)
	server := http.Server{
		ReadTimeout: time.Minute,
	}

	log.Info(menuItems[0].displayName)
	log.Info(menuItems[0].id)
	log.Info(menuItems[0].amountInStock)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleMenu(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Info("Received GET /menu")

	_, err := fmt.Fprintf(w, "%+v", menuItems)
	if err != nil {
		log.Fatal(err)
	}
}
