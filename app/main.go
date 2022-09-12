package main

import (
	"fmt"
	"log"
	"net/http"

	_bankinghandler "github.com/banking/banking/delivery"
	_bankrepo "github.com/banking/banking/repository"
	_bankingucase "github.com/banking/banking/usecase"
	config "github.com/banking/config"
	"github.com/gorilla/mux"


	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	err := config.GetDatabaseConfig()
	CheckError(err)

}
func main() {
	mux := mux.NewRouter()

	connection, err := GetConnection()
	CheckError(err)

	ar := _bankrepo.NewPostgresBankingRepository(connection)
	au := _bankingucase.NewBankingUsecase(ar)
	router := _bankinghandler.NewBankHandler(mux, au)

	log.Fatal(http.ListenAndServe(config.ApplicationPort.Port, router))
}

// GetConnection : returns the connection to postgres server
func GetConnection() (*gorm.DB, error) {

	db := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Databaseconfig.DBHost, config.Databaseconfig.DBPort, config.Databaseconfig.DBUser, config.Databaseconfig.DBPassword, config.Databaseconfig.DBName)

	// open database
	conn, err := gorm.Open("postgres", db)
	CheckError(err)
	// check db
	err = conn.DB().Ping()
	CheckError(err)

	return conn, nil

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
