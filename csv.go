package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	flag.Parse()

	file1, err := os.Open("KEN_ALL.CSV")
	failOnError(err)
	defer file1.Close()

	reader := csv.NewReader(transform.NewReader(file1, japanese.ShiftJIS.NewDecoder()))

	//reader := csv.NewReader(file1)
	reader.LazyQuotes = true

	log.Printf("Start")

	for num := 0; num < 125; num++ {

		hoge := num + 1
		phoge := fmt.Sprintf("%03d", hoge) //0埋め
		//numi := strconv.Itoa(hoge)
		//整数が格納された変数hogeをint型からstring型に変換。strconv.Atoiでstring→int

		filex, err := os.Create("outout_" + phoge + ".csv")
		//filex, err := os.Create("outout_" + numi + ".csv")
		if err != nil {
			defer filex.Close()

		}

		//writer := csv.NewWriter(filex)
		writer := csv.NewWriter(transform.NewWriter(filex, japanese.ShiftJIS.NewEncoder()))

		writer.UseCRLF = true
		//  writer.Comma = '\t'

		for z := 0; z < 1000; z++ {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else {
				failOnError(err)
			}
			var new_record []string
			for _, v := range record {
				new_record = append(new_record, v)
			}

			writer.Write(new_record)
			writer.Flush()
		}

	}

	log.Printf("Finish !")
}
