package dbops

import (
	"errors"
	"fmt"
)

func NewAccount(name string,balance float64) error {
	i, e := Engine.Insert(&Account{Name: name, Balance: balance})
	fmt.Println(i)
	return e
}

func GetUserByName(name string) (*Account,error) {
	account := &Account{Name: name}
	_, e := Engine.Get(account)
	return account,e
}

func DepositByName(name string,deposit float64) error {
	account := &Account{Name: name}
	_, e := Engine.Get(account)
	if e != nil {
		return e
	}

	account.Balance += deposit
	_, e = Engine.Where("name = ?",account.Name).Update(account)
	return e
}

func WithdrawMoney(name string,deposit float64) error {
	account := &Account{Name: name}

	_, e := Engine.Get(account)
	if e!=nil {
		return e
	}

	account.Balance -= deposit
	_, e = Engine.Where("name = ?",account.Name).Update(account)
	return e
}

func TransferAccounts(form,to string,balance float64) error {
	session := Engine.NewSession() // 获得session
	defer session.Close() // session会消耗资源 所以要关闭
	err := session.Begin() //开启事务
	if err != nil {
		return errors.New("session begin error")
	}

	account := &Account{Name: form}
	_, e := Engine.Get(account)
	if e != nil {
		return e
	}
	if account.Balance > balance{
		account.Balance -= balance
	}
	
	account2 := &Account{Name:to}
	b, e := Engine.Get(account2)
	if b != true || e!= nil{
		return errors.New("Error")
	}
	account2.Balance += balance

	_, e = Engine.Where("name = ?",account.Name).Update(account)
	if e!=nil{
		session.Rollback() //回滚

		return e
	}

	_, e = Engine.Where("name = ?",account2.Name).Update(account2)
	if e!=nil{
		session.Rollback() //回滚
	}

	return session.Commit() //提交
}

func GetAccountsDESCId() ([]*Account,error) {
	var as []*Account
	err := Engine.Desc("id").Find(&as)
	return as,err
}

func DeleteAccountByName(name string) error {
	account := &Account{Name: name}
	_, e := Engine.Delete(account)
	return e
}


func row() {
	rows, _ := Engine.Rows(new(Account))
	defer rows.Close()

	account := new(Account)

	for rows.Next()  {
		rows.Scan(account)
	}
}