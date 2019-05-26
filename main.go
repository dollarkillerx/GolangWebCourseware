package main

import (
	"GolangWebCourseware/dbops"
	"fmt"
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
				panic("插入数据时出错")
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

		case 9:
			break forlook
		}
		
	}

	defer func() {
		fmt.Println("Bay Bay ...")
	}()
}