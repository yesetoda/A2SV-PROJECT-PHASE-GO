package controller

import (
	"example/bank_api/models"
	"fmt"
)

type BankManeger struct {
	Users map[int]models.User
	Logs  map[int]models.Log
}

func (bk *BankManeger) RegisterUser(user models.User) {
	bk.Users[user.Id] = user
}

func (bk *BankManeger) RemoveUser(userId int) {
	_, ufound := bk.Users[userId]
	if !ufound {
		fmt.Println("this user is not found")
		return
	}
	delete(bk.Users, userId)
}

func (bk *BankManeger) ShowBalance(userId int) {
	user, ufound := bk.Users[userId]
	if !ufound {
		fmt.Println("no user with this id")
		return
	}
	fmt.Println(user.Balance)
}

func (bk *BankManeger) SendMoney(senderId int, recieverId int, amount float64, logid *int) bool { //if theis function returns true increment the next
	sender, sfound := bk.Users[senderId]
	if !sfound {
		fmt.Println("there is no such sender account")
		return false
	}
	reciever, rfound := bk.Users[recieverId]
	if !rfound {
		fmt.Println("there is no such reciever account")
		return false
	}

	if sender.Balance < amount {
		fmt.Println("insufficient balance ")
		return false
	}
	bk.Logs[*logid] = models.Log{
		SenderId:   senderId,
		RecieverId: recieverId,
		Amount:     amount,
	}
	sender.Logs[*logid] = bk.Logs[*logid]
	reciever.Logs[*logid] = bk.Logs[*logid]
	*logid += 1 //does this work check it

	sender.Balance -= amount
	reciever.Balance += amount
	bk.Users[senderId] = sender
	bk.Users[recieverId] = reciever
	return true
}

func (bk *BankManeger) WithdrawMoney(userId int, amount float64, logid *int) bool { //if theis function returns true increment the next
	user, ufound := bk.Users[userId]
	if !ufound {
		fmt.Println("no user with this id")
		return false
	}
	if user.Balance < amount {
		fmt.Println("insufficient balance")
		return false
	}

	bk.Logs[*logid] = models.Log{
		SenderId:   userId,
		RecieverId: -1,
		Amount:     -amount,
	}
	user.Logs[*logid] = bk.Logs[*logid]
	*logid += 1 //does this work check it
	user.Balance -= amount
	bk.Users[userId] = user
	return true
}

func (bk *BankManeger) Log() {
	for _, i := range bk.Logs {
		fmt.Println(i)
	}
}

func (bk *BankManeger) DepositMoney(userId int, amount float64, logid *int) {
	user, ufound := bk.Users[userId]
	if !ufound {
		fmt.Println("no user with this id")
		return
	}
	bk.Logs[*logid] = models.Log{
		SenderId:   userId,
		RecieverId: -1,
		Amount:     amount,
	}
	user.Logs[*logid] = bk.Logs[*logid]
	*logid += 1 //does this work check it
	user.Balance += amount
	bk.Users[userId] = user
}
