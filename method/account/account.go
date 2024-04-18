package account

type Account struct {
	Balance int
}

// func (리시버) 메서드명(매개변수) 리턴타입 {}
func (a *Account) WithdrawPointer(amount int) {
	a.Balance -= amount
}

func (a Account) WithdrawValue(amount int) Account {
	a.Balance -= amount
	return a
}

