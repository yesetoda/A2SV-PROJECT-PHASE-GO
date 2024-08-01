package main

import (
	"bufio"
	"example/bank_api/controller"
	"example/bank_api/models"
	"fmt"
	"os"
	"strconv"
)

var (
	reader    = bufio.NewReader(os.Stdin)
	NextUsrId = 0
	NextLogId = 0
	bk        = controller.BankManeger{
		Users: make(map[int]models.User),
		Logs:  make(map[int]models.Log),
	}
)

func main() {
	for {
		fmt.Println(`
			1,Register User
			2,Remove User
			3,Show Balance
			4,Send Money
			5,Withdraw Money
			6,Show Logs
			7,Show Users
			8,Deposit Money
		`)
		var str_choice string
		fmt.Scan(&str_choice)
		choice, err := strconv.Atoi(str_choice)
		if err != nil {
			fmt.Println("invalid input!")
			continue
		}
		switch choice {
		case 1:
			Register_User()
		case 2:
			Remove_User()
		case 3:
			Show_Balance()
		case 4:
			Send_Money()
		case 5:
			Withdraw_money()
		case 6:
			Show_Log()
		case 7:
			Show_users()
		case 8:
			Deposit_Money()
		default:
			fmt.Println("invalid input")
			continue
		}

	}

}
func Register_User() {
	fmt.Print("Enter Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("invalid name")
		return
	}
	fmt.Print("Enter Balance: ")

	var str_balance string
	fmt.Scan(&str_balance)
	balance, err := strconv.ParseFloat(str_balance, 64)
	if err != nil {
		fmt.Println("invalid Balance")
		return
	}

	user := models.User{
		Id:      NextUsrId,
		Name:    name,
		Balance: balance,
		Logs:    make(map[int]models.Log),
	}
	bk.RegisterUser(user)
	NextUsrId += 1
}

func Remove_User() {
	var str_userId string
	fmt.Print("Enter User Id: ")
	fmt.Scan(&str_userId)
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		fmt.Println("invalid id")
		return
	}
	bk.RemoveUser(userId)
}
func Send_Money() {
	var str_senderId string
	fmt.Print("Enter Senders Id: ")

	fmt.Scan(&str_senderId)
	senderId, err := strconv.Atoi(str_senderId)
	if err != nil {
		fmt.Println("invalid sender id")
		return
	}

	var str_recieverId string
	fmt.Print("Enter Recievers Id: ")

	fmt.Scan(&str_recieverId)
	recieverId, err := strconv.Atoi(str_recieverId)
	if err != nil {
		fmt.Println("invalid reciever id")
		return
	}
	var str_amount string
	fmt.Print("Enter amount: ")

	fmt.Scan(&str_amount)
	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		fmt.Println("invalid amount")
		return
	}

	bk.SendMoney(senderId, recieverId, amount, &NextLogId)
}

func Withdraw_money() {
	var str_userId string
	fmt.Print("Enter User Id: ")
	fmt.Scan(&str_userId)
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		fmt.Println("invalid user id")
		return
	}
	var str_amount string
	fmt.Print("Enter amount: ")

	fmt.Scan(&str_amount)
	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		fmt.Println("invalid amount")
		return
	}

	bk.WithdrawMoney(userId, amount, &NextLogId)
}
func Show_Balance() {
	var str_userId string
	fmt.Print("Enter User Id: ")

	fmt.Scan(&str_userId)
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		fmt.Println("invalid user id")
		return
	}
	bk.ShowBalance(userId)
}
func Show_Log() {
	bk.Log()
}

func Show_users() {
	fmt.Println("here is the list of users")
	for _, i := range bk.Users {
		fmt.Println(i)
	}
}

func Deposit_Money() {
	var str_userId string
	fmt.Print("Enter User Id: ")
	fmt.Scan(&str_userId)
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		fmt.Println("invalid user id")
		return
	}
	var str_amount string
	fmt.Print("Enter amount: ")

	fmt.Scan(&str_amount)
	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		fmt.Println("invalid amount")
		return
	}
	bk.DepositMoney(userId, amount, &NextLogId)
}
