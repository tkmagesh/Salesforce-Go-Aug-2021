package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Category struct {
	Id   int    `json:"id"`
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
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %v", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	var products []Product = []Product{}
	file, err := os.Open("product.json")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&products); err != nil {
		log.Fatalln(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("Done")
}
