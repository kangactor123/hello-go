package main

import (
	"fmt"
	"goproject/method/account"
	"goproject/method/alias"
)

func main() {
	a := &account.Account{Balance: 100}
	a.WithdrawPointer(30)

	fmt.Printf("%d \n", a.Balance)

	var b alias.MyInt = 10
	fmt.Println(b.Add(30))
	var c alias.MyInt = 20
	fmt.Println(alias.MyInt(c).Add(50))

	account := account.Account{Balance: 300}

	account.WithdrawPointer(30)
	account.WithdrawValue(30)

	fmt.Println(account)
}