package main

import (
	"fmt"
	"log"
)

func main() {
	goodMachine := newGoodMachine(1, 10)
	err := goodMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = goodMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = goodMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println()

	err = goodMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = goodMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = goodMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
