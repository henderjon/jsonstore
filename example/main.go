package main

import (
	"github.com/henderjon/jsonstore"
	"log"
)

func main() {

	data := struct{
		Nuts, Langs []string
		Tomato, Potato, Kale string
		Five int
	}{
		[]string{"Pistachio", "Cashew"},
		[]string{"C", "Go", "Ruby", "PHP"},
		"Red", "White", "Green",
		5,
	}

	j, err := jsonstore.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	j.Put("data.json", data)

	// data.Nuts = append(data.Nuts, "Pecans")

	// j.Put("data.json", data)

	// data.Langs = data.Langs[1:2]

	// j.Put("data.json", data)

}
