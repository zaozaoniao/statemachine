package main

import "fmt"

type itemRequestedState struct {
	*goodMachine
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("shopping is  in  process")
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

// 付钱----> 已收钱状态
func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.goodMachine.itemPrice {
		fmt.Errorf("insert money is less, please insert %d", i.goodMachine)
	}
	fmt.Println("money entered is ok")
	i.setState(&hasMoneyState{i.goodMachine})
	return nil
}
func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}
