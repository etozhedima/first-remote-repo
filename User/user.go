package user

import "fmt"

var paymentID = gemPayID()

func gemPayID() func() int { // генерируем уникальный айди платежа, не повторяющийся ни у одного платящегося
	id := 0
	return func() int {
		id++
		return id
	}

}

type User struct {
	UID          int
	Balance      float64
	PaymentsInfo map[int]Payments
}

type Payments struct {
	Description string
	PayValue    float64
	PayStatus   bool
}

func NewUser(id int, balance float64) *User {
	return &User{
		UID:          id,
		Balance:      balance,
		PaymentsInfo: make(map[int]Payments),
	}
}

func (u *User) Pay(value float64, description string) (string, error) {
	if u.Balance-value < 0 {
		return "", fmt.Errorf("Платеж на сумму %.2f совершить не удалось. "+
			"Недостаточно средств на балансе. Текущий баланс %.2f", value, u.Balance)
	}

	newID := paymentID()

	u.Balance -= value
	u.PaymentsInfo[newID] = Payments{Description: description, PayValue: value, PayStatus: true}
	return fmt.Sprintf("Платеж совешен! Текущий баланс %.2f. ID платежа: %d", u.Balance, newID), nil
}

func (u *User) CancelPay(payId int) (string, error) {
	if _, ok := u.PaymentsInfo[payId]; !ok {
		return "", fmt.Errorf("платеж с айди %d не найден. Отмена платежа невозможна", payId)
	}
	u.Balance += u.PaymentsInfo[payId].PayValue
	payment := u.PaymentsInfo[payId]
	payment.PayStatus = false
	u.PaymentsInfo[payId] = payment
	return fmt.Sprintf("Платеж отменен! Текущий баланс %.2f\n", u.Balance), nil
}
