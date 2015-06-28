package main

//usage: mtr --report --raw <hostname or ip> | go run mtrparser.go

import (
	"encoding/json"
	"fmt"
	"github.com/sajal/mtrparser"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	result, err := mtrparser.NewMTROutPut(string(bytes), 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Result")
	fmt.Println(result)
	fmt.Println("Json")
	b, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	fmt.Println("Line by line")
	for idx, item := range result.Hops {
		fmt.Printf("%v : %s (%s) Avg: %v, Loss : %v%% Best: %v Worst: %v Last: %v\n", idx+1, item.Host, item.IP, item.Avg, item.Loss, item.Best, item.Worst, item.Last)
	}
}