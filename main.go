package main

import (
	"fmt"
	user "goStudy/User"
)

type Payer interface {
	Pay(value float64, description string) (string, error)
	CancelPay(payId int) (string, error)
}

func PayerPays(payer Payer, value float64, description string) {
	result, err := payer.Pay(value, description)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}

func PayerCancelPay(payer Payer, id int) {
	result, err := payer.CancelPay(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}

func main() {
	u := user.NewUser(1, 100)
	u2 := user.NewUser(2, 150)

	c := user.NewCompany(1, 100000)

	PayerPays(u, 10.99, "sub") // ок платеж
	PayerPays(u, 199, "food")  // недостаточно средств
	PayerCancelPay(u, 2)       // такого айди нет для u
	fmt.Println()
	PayerPays(u2, 10, "sub")
	PayerPays(u2, 199, "food")
	PayerCancelPay(u2, 2) // такой айди есть для u платеж отменен
	fmt.Println()
	PayerPays(c, 999.99, "sub")
	PayerPays(c, 200000, "food")
	PayerCancelPay(c, 3)
}
