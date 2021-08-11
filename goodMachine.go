package main

import (
	"fmt"
	"reflect"
)

type goodMachine struct {
	currentState state
	itemCount    int
	itemPrice    int
}

func newGoodMachine(itemCount, itemPrice int) *goodMachine {
	v := &goodMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	if itemCount <= 0 {
		v.setState(&noItemState{v}) // 实现state接口的是*noItemState 指针类型
	} else {
		v.setState(&hasItemState{v})
	}
	return v
}

func (v *goodMachine) setState(s state) {
	fmt.Println("enter state: ", reflect.TypeOf(s))
	v.currentState = s
}

func (v *goodMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *goodMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *goodMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *goodMachine) incrementItemCount(count int) {
	v.itemCount += count
}

func (v goodMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}
