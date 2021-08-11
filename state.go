package main

//  代表某种状态，能接受的某种动作
type state interface {
	addItem(count int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
