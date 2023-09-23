package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main(){
	//open csv
	irisFile, err := os.Open("./Iris.csv")
	if err != nil{
		log.Fatal(err)
	}
	defer irisFile.Close()

	//create datframe from csv
	//types of columns will be infered
	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)

}

