package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

type Item struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Group   string    `json:"group"`
	Count   int32     `json:"count"`
	DateRec time.Time `json:"daterec"`
	DateExp time.Time `json:"dateexp"`
	Price   float64   `json:"price"`
	Unit    string    `json:"unit"`
}

var items []Item

func generateFakeItems(count int) {
	groups := []string{"Cold-Storage", "Dry-Goods", "Produce", "Meat", "Dairy"}
	units := []string{"kg", "lbs", "pcs", "bottles", "cans"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		item := Item{
			ID:      faker.UUIDDigit(),
			Name:    faker.Word(),
			Group:   groups[r.Intn(len(groups))],
			Count:   int32(r.Intn(100) + 1),
			DateRec: time.Now().AddDate(0, 0, -r.Intn(30)),
			DateExp: time.Now().AddDate(0, 0, r.Intn(90)),
			Price:   float64(r.Intn(10000)) / 100,
			Unit:    units[r.Intn(len(units))],
		}
		items = append(items, item)
	}
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func addItem(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func main() {
	generateFakeItems(50)

	router := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/items", getItems)
	router.POST("/items", addItem)

	fmt.Println("Server is running on http://localhost:8080")
	router.Run("localhost:8080")
}
