package user

import "fmt"

type Company struct {
	UID          int
	Balance      float64
	PaymentsInfo map[int]Payments
}

func NewCompany(id int, balance float64) *Company {
	return &Company{
		UID:          id,
		Balance:      balance,
		PaymentsInfo: make(map[int]Payments),
	}
}

func (c *Company) Pay(value float64, description string) (string, error) {
	if c.Balance-value < 0 {
		return "", fmt.Errorf("недостаточно средств на балансе. Текущий баланс %.2f", c.Balance)
	}
	newID := paymentID()

	c.Balance -= value
	c.PaymentsInfo[newID] = Payments{Description: description, PayValue: value, PayStatus: true}
	return fmt.Sprintf("Платеж совешен! Текущий баланс %.2f. ID платежа: %d", c.Balance, newID), nil
}

func (c *Company) CancelPay(payId int) (string, error) {
	if _, ok := c.PaymentsInfo[payId]; !ok {
		return "", fmt.Errorf("платеж с айди %d не найден. Отмена платежа невозможна", payId)
	}
	c.Balance += c.PaymentsInfo[payId].PayValue
	//перезапись всей структуры выглядит Payments выглядит как костыль
	payment := c.PaymentsInfo[payId]
	payment.PayStatus = false
	c.PaymentsInfo[payId] = payment
	return fmt.Sprintf("Платеж отменен! Текущий баланс %.2f\n", c.Balance), nil
}
