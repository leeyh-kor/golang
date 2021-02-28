package main

import (
	"fmt"

	"github.com/leeyh-kor/golang/banking"
) //  큰 따옴표 작은 따옴표 구분해야함

func main() {
	account := banking.BankAccount{Owner: "yh", Balance: 100}
	fmt.Println(account)
}
