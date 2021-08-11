package main

import "fmt"

type noItemState struct {
	*goodMachine // 存在匿名类型 goodMachine，类型是*goodMachine
}

//  给自动售货机供货-----> 有货状态
func (i *noItemState) addItem(count int) error {
	i.incrementItemCount(count)
	if i.itemCount > 0 {
		i.setState(&hasItemState{i.goodMachine})
	}
	return nil
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("item out of  stock")
}

func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// golang: 使用指针接受者实现了state接口的全部函数，那么隐式表明*noItemState 指针类型实现了State接口
