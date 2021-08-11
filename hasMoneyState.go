package main

import "fmt"

type hasMoneyState struct {
	*goodMachine
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("shopping is in process")
}
func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("shopping is in process")
}
func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("already pay money")
}
func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("dispensing item")
	i.itemCount = i.goodMachine.itemCount - 1
	if i.itemCount == 0 {
		i.setState(&noItemState{i.goodMachine})
	} else {
		i.setState(&hasItemState{i.goodMachine})
	}
	return nil
}
