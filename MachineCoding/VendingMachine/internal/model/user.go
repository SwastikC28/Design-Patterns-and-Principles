package model

var (
	CREDIT_CARD = 1
	CASH        = 2
)

type User struct {
	isAdmin        bool
	paymentMethods map[int]PaymentMode
}

func NewUser(isAdmin bool) *User {
	return &User{
		isAdmin:        isAdmin,
		paymentMethods: make(map[int]PaymentMode),
	}
}

func (user *User) AddPaymentMode(key int, mode PaymentMode) {
	user.paymentMethods[key] = mode
}

func (user *User) GetPaymentMode(key int) PaymentMode {
	mode, exists := user.paymentMethods[key]
	if !exists {
		return nil
	}

	return mode
}
