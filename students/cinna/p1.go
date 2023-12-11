package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main(){
	//open the csv file
	file,err := os.Open("/workspaces/quiz/problems.csv")
	if err!=nil {
		fmt.Println("error when opening csv file",err)
		return
	}
	defer file.Close()

	//create a reader
	reader := csv.NewReader(file)

	//read data lines
	records, err:= reader.ReadAll()
	if err!=nil {
		fmt.Println("error when reading data",err)
		return
	}

	//process data
	for _,record :=  range records{
		for _,value := range record {
			fmt.Printf("%s",value)
		}
		fmt.Println()
	}
}

