package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	var data string

	flag.StringVar(&data, "d", "", "Dataframe to load. Use -h or --help for more")
	flag.Parse()

	if data == "" {
		fmt.Println("No data provided. Please provide dataframe. Use --help or -h for more info")
		os.Exit(1)
	}

	file, err := os.Open(data)
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(file)
	row, column := df.Dims()
	fmt.Println("\nSummary of data provided")
	fmt.Println("\nShape of dataframe", "\nNumber of rows:", row, "\nNumber of columns:", column)
	fmt.Println("\nNumber of values:", row*column)
	fmt.Println("\nSummary Statistics of dataframe:", df.Describe())

	for column := 0; column < df.Ncol(); column++ {
		name := df.Select(column).Names()
		ds := df.Col(name[0])
		fmt.Printf("\nDoes the column %s have any null values? ", name[0])
		fmt.Println(ds.HasNaN())

	}
}
