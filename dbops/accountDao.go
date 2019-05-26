package dbops

import "fmt"

func NewAccount(name string,balance float64) error {
	i, e := Engine.InsertOne(&Account{Name: name, Balance: balance})
	fmt.Println(i)
	return e
}

func GetUserByName(name string) (*Account,error) {
	account := &Account{Name: name}
	_, e := Engine.Get(account)
	return account,e
}