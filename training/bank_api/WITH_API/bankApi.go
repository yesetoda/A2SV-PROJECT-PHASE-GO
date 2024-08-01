package main

import (
	"encoding/json"
	"example/bank_api/controller"
	"example/bank_api/models"
	"fmt"

	"net/http"
	"strconv"
)

var (
	NextUsrId = 0
	NextLogId = 0
	bk        = controller.BankManeger{
		Users: make(map[int]models.User),
		Logs:  make(map[int]models.Log),
	}
)

func main() {
	server := http.NewServeMux()
	port := "9090"
	server.HandleFunc("GET /balance", handleShowBalance)
	server.HandleFunc("GET /logs", handleShowLog)
	server.HandleFunc("GET /users", handleShowusers)
	server.HandleFunc("POST /user", handleRegisterUser)
	server.HandleFunc("DELETE /user", handleRemoveUser)
	server.HandleFunc("PATCH /deposit", handleDepositMoney)
	server.HandleFunc("PATCH /send", handleSendMoney)
	server.HandleFunc("PATCH /withdraw", handleWithdrawmoney)

	fmt.Println("serving at port:", port)
	http.ListenAndServe("localhost:"+port, server)
}

// GET methods

func handleShowBalance(w http.ResponseWriter, r *http.Request) {
	str_userId := r.FormValue("id")
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		w.Write([]byte("invalid user id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ufound := bk.Users[userId]
	if !ufound {
		w.Write([]byte("user not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bk.ShowBalance(userId)
	w.WriteHeader(http.StatusAccepted)
}
func handleShowLog(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bk.Logs)
	w.WriteHeader(http.StatusAccepted)
}

func handleShowusers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bk.Users)
	w.WriteHeader(http.StatusAccepted)
}

// POST methods

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	balance, err := strconv.ParseFloat(r.FormValue("balance"), 64)
	if err != nil {
		w.Write([]byte("invalid balance"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if balance<0{
		w.Write([]byte("balance <0"))
		w.WriteHeader(http.StatusBadRequest)
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
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("sucessfully registered user"))


}

// DELETE METHODS
func handleRemoveUser(w http.ResponseWriter, r *http.Request) {
	str_userId := r.FormValue("id")
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		w.Write([]byte("invalid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ufound := bk.Users[userId]
	if !ufound {
		w.Write([]byte("user not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bk.RemoveUser(userId)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("sucessfully removed user"))

}

// PATCH methods
func handleSendMoney(w http.ResponseWriter, r *http.Request) {
	str_senderId := r.FormValue("senderId")
	str_recieverId := r.FormValue("recieverId")
	str_amount := r.FormValue("amount")

	senderId, err := strconv.Atoi(str_senderId)
	if err != nil {
		w.Write([]byte("invalid sender id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recieverId, err := strconv.Atoi(str_recieverId)
	if err != nil {
		w.Write([]byte("invalid reciever id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		w.Write([]byte("invalid amount"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sender, sfound := bk.Users[senderId]
	if !sfound {
		w.Write([]byte("sender not found "))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, rfound := bk.Users[recieverId]
	if !rfound {
		w.Write([]byte("reciever not found "))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if amount<0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("amount<0"))
		return
	}
	if sender.Balance < amount {
		w.Write([]byte("insufficient Balance "))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bk.SendMoney(senderId, recieverId, amount, &NextLogId)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("sucessfully sent"))


}

func handleWithdrawmoney(w http.ResponseWriter, r *http.Request) {
	var str_userId = r.FormValue("id")
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		w.Write([]byte("invalid user id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var str_amount = r.FormValue("amount")
	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		w.Write([]byte("invalid amount"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, ufound := bk.Users[userId]
	if !ufound {
		w.Write([]byte("user not found "))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if amount<0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("amount<0"))
		return
	}
	if user.Balance < amount {
		w.Write([]byte("insufficient Balance "))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bk.WithdrawMoney(userId, amount, &NextLogId)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("sucessfully withdrawed"))


}

func handleDepositMoney(w http.ResponseWriter, r *http.Request) {
	var str_userId = r.FormValue("id")
	userId, err := strconv.Atoi(str_userId)
	if err != nil {
		w.Write([]byte("invalid user id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var str_amount = r.FormValue("amount")
	amount, err := strconv.ParseFloat(str_amount, 64)
	if err != nil {
		w.Write([]byte("invalid amount"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ufound := bk.Users[userId]
	if !ufound {
		w.Write([]byte("user not found "))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if amount<0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("amount<0"))
		return
	}
	bk.DepositMoney(userId, amount, &NextLogId)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("sucessfully deposited"))

}
