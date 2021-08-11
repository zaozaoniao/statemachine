package main

import "fmt"

type hasItemState struct {
	*goodMachine
}

func (v *hasItemState) addItem(count int) error {
	v.incrementItemCount(count)
	return nil
}

// 有人选择了商品---> 没货状态/已经选定商品
func (v *hasItemState) requestItem() error {
	if v.itemCount == 0 {
		v.setState(&noItemState{v.goodMachine})
		return fmt.Errorf("no item present")
	}

	fmt.Print("item  requested\n")
	v.setState(&itemRequestedState{v.goodMachine})
	return nil
}

func (v *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (v *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
