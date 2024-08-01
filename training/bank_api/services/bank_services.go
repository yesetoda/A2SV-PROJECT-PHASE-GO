package services

type Bank interface{
	RegisterUser()
	RemoveUser()
	ShowBalance()
	SendMoney()
	WithdrawMoney()
	DepositMoney()
	Log()
}