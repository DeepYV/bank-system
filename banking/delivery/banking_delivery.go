package delivery

import (
	"encoding/json"
	"net/http"

	domain "github.com/banking/domain"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
)

// BankHandler takes usecase interface and implements it.
type BankHandler struct {
	bankUsecase domain.BankingUseCase
}

// NewBankHandler route to end point
func NewBankHandler(router *mux.Router, BankUsecase domain.BankingUseCase) *mux.Router {

	handler := &BankHandler{

		bankUsecase: BankUsecase,
	}

	router.HandleFunc("/bank/create", handler.create).Methods("POST")
	router.HandleFunc("/bank/transaction", handler.transaction).Methods("POST")
	router.HandleFunc("/bank/{id}", handler.balance).Methods("GET")
	router.HandleFunc("/bank/records/{id}", handler.history).Methods("GET")

	return router
}

// create calls the CreateAccount usecase
func (handler BankHandler) create(w http.ResponseWriter, r *http.Request) {

	var user domain.Account

	// intialize first amount
	user.Balance = domain.DefaultAmount

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		internalServerError(w, r)

		return
	}

	err := handler.bankUsecase.CreateAccount(user)
	if err != nil {

		log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot create User with this information"))

		return

	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Created"))

	return
}

// transaction calls transfer usecase.
func (handler BankHandler) transaction(w http.ResponseWriter, r *http.Request) {

	var transaction domain.Transfer

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {

		internalServerError(w, r)

		return
	}

	err := handler.bankUsecase.Transfer(transaction)
	if err != nil {
		log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Trasaction incomplete try again !!"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("trasaction completed"))

	return

}

// balance call usecase to check the balance
func (handler BankHandler) balance(w http.ResponseWriter, r *http.Request) {

	id := GetId(r)

	balance, err := handler.bankUsecase.GetBalance(id)
	if err != nil {

		log.Error(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user not found"))

		return

	}

	amount, err := json.Marshal(balance)
	if err != nil {
		log.Error(err)

		internalServerError(w, r)

		return

	}
	w.Write([]byte(amount))

	return

}

// hisotry call the gethistory usecall to get the entry
func (handler BankHandler) history(w http.ResponseWriter, r *http.Request) {

	accountId := GetId(r)

	records, err := handler.bankUsecase.GetHistory(accountId)

	if err != nil {

		log.Error(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to fetch record"))

		return
	}
	entries, err := json.Marshal(records)
	if err != nil {

		log.Error(err)

		internalServerError(w, r)

		return

	}

	w.Write([]byte(entries))

	return

}
