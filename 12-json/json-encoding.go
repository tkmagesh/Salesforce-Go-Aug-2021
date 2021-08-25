package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Category struct {
	Id   int    `json:"id" xml:"id"`
	Name string `json:"name"`
}

type Product struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Cost     float32  `json:"cost"`
	Units    int      `json:"units"`
	Category Category `json:"category"`
}

/* fmt.Stringer interface implementation */
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	stationaryCategory := Category{1, "Stationary"}
	utencilCategory := Category{2, "Utencil"}
	electronCategory := Category{3, "Electronics"}

	products := []Product{
		{105, "Pen", 5, 50, stationaryCategory},
		{103, "Pencil", 2, 100, stationaryCategory},
		{102, "Marker", 50, 20, stationaryCategory},
		{104, "Frying Pan", 500, 5, utencilCategory},
		{101, "Phone", 5000, 3, electronCategory},
		{100, "Bowl", 100, 50, utencilCategory},
	}
	file, err := os.Create("product.json")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(products); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done")
}
