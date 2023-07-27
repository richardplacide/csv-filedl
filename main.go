package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type RV struct {
	Id    string
	Make  string
	Model string
	Price string
}

const (
	Id    = 0 // ID = iota then just Make Model Price will be 1,2,3
	Make  = 1
	Model = 2
	Price = 3
)

func main() {
	f, err := os.Open("csv1.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(f))
	// discard header
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		jsonB, err := json.Marshal(RV{
			Id:    record[Id],
			Model: record[Model],
			Price: record[Price],
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(os.Stdout, "%s\n", jsonB)

		/* fmt.Println(RV{
			Id:    record[Id],
			Model: record[Model],
			Price: record[Price],
		}) */
	}
}
