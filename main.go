package main

import (
	"GolangWebCourseware/dbops"
	"fmt"
	"log"
)

const (
	prompt = `
Please enter number of operation:
1. Create new account
2. Show detail of account
3. Deposit
4. Withdraw
5. Make transfer
6. List account by Id
7. List account by balance
8. Delete account
9. Exit
`
)

func main() {
	fmt.Println("Welcome bank of xorm!")
	forlook:
	for {
		fmt.Printf(prompt)
		var num int
		_, err := fmt.Scanf("%d\n", &num)
		if err !=nil {
			fmt.Printf("输入错误!类型必须为int")
			continue
		}
		switch num {
		case 1:
			fmt.Println("Please enter <name> <balance>:")
			var name string
			var balance float64
			i, err := fmt.Scanf("%s %f\n", &name, &balance)
			if err != nil{
				fmt.Printf("输错错入!")
				break
			}

			if err = dbops.NewAccount(name, balance);err != nil{
				panic(err.Error())
			}
			fmt.Printf("%v\n",i)
		case 2:
			fmt.Println("Please enter <name> :")
			var name string
			_, err := fmt.Scanf("%s\n", &name)
			if err != nil{
				fmt.Printf("输错错入!")
				break
			}
			if account, err := dbops.GetUserByName(name);err != nil{
				fmt.Printf("查询数据不存在")
				break
			}else{
				fmt.Printf("%#v",account)
			}
		case 3:
			fmt.Println("Please enter <name> <deposit>:")
			var name string
			var deposit float64
			_, err := fmt.Scanf("%s %f\n",&name,&deposit)
			if err != nil{
				fmt.Println("数据输入错误!")
				break
			}
			err = dbops.DepositByName(name, deposit)
			if err != nil {
				log.Println(err.Error())

				fmt.Println("数据输入错误!")
				break
			}
			fmt.Println("存款成功")
		case 4:
			fmt.Println("Please enter <name> <deposit>:")
			var name string
			var deposit float64
			_, err := fmt.Scanf("%s %f\n", &name, &deposit)
			if err != nil {
				fmt.Println("数据输入错误!")
				break
			}
			err = dbops.WithdrawMoney(name, deposit)
			if err != nil {
				log.Println(err.Error())
				fmt.Println("数据输入错误!")
				break
			}
			fmt.Println("存款成功")
		case 5:
			fmt.Println("Please enter <name> <balance> <name>:")
			var name1,name2 string
			var balance float64
			_, err := fmt.Scanf("%s %s %f\n", &name1, &name2, &balance)
			if err != nil {
				log.Print(err.Error())
				fmt.Println("数据输入错误!01")
				break
			}

			err = dbops.TransferAccounts(name1, name2, balance)
			if err!= nil{
				log.Print(err.Error())
				fmt.Println("数据输入错误!02")
				break
			}
			fmt.Println("转账成功!")
		case 6:
			if account, err := dbops.GetAccountsDESCId();err != nil{
				fmt.Println("数据错误!")
				break
			}else{
				for _,k := range account {
					fmt.Println(k)
				}
			}
		case 7:
			fmt.Println("Please enter <name> :")
			var name string
			_, err := fmt.Scanf("%s\n",&name)
			if err != nil {
				fmt.Println("数据输入错误!")
				break
			}

			err = dbops.DeleteAccountByName(name)
			if err!=nil{
				fmt.Println("数据错误!")
				break
			}
			fmt.Println("删除成功")
		case 9:
			break forlook
		}
		
	}

	defer func() {
		fmt.Println("Bay Bay ...")
	}()
}