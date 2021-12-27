package logging

import (
	"errors"

	"github.com/bot-api/telegram"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type InitLoggerInput struct {
	WriteErrorsToDatabase bool
	PGDB *pg.DB
	TelegramToken string
	NotificationOnAllErrors bool
	TelegramNotificationsChannelID int64
	ServiceName string
}

type NewLogger struct {
	TelegramAPI *telegram.API
	PGDB *pg.DB
	WriteErrorsToDatabase bool
	NotificationOnAllErrors bool
	TelegramNotificationsChannelID int64
	ServiceName string
}

// A variable to determine if logging mechanism has been initiated
var LoggerInitiated bool

// Global Logger 
var Logger NewLogger

func Init(initLoggerInput *InitLoggerInput) {

	Logger.WriteErrorsToDatabase = initLoggerInput.WriteErrorsToDatabase
	Logger.NotificationOnAllErrors = initLoggerInput.NotificationOnAllErrors
	Logger.ServiceName = initLoggerInput.ServiceName

	// Create Database Table
	if initLoggerInput.WriteErrorsToDatabase == true {

		if initLoggerInput.PGDB == nil {
			panic(errors.New("InitLoggerInput: PGDB is a required field if WriteErrorsToDatabase is true."))
		}

		err := initLoggerInput.PGDB.CreateTable((*ErrorLog)(nil), &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			panic(err)
		}

		Logger.PGDB = initLoggerInput.PGDB

	}

	// Init Telegram API
	if initLoggerInput.NotificationOnAllErrors == true && initLoggerInput.TelegramToken == "" {
		panic(errors.New("TelegramToken is required when using notifications."))
	}
	if initLoggerInput.TelegramNotificationsChannelID == 0 && initLoggerInput.NotificationOnAllErrors == true {
		panic(errors.New("TelegramNotificationsChannelID is required when using notifications. test update"))
	}
	if initLoggerInput.TelegramToken != "" {
		Logger.TelegramAPI = telegram.New(initLoggerInput.TelegramToken)
		Logger.TelegramNotificationsChannelID = initLoggerInput.TelegramNotificationsChannelID
	} 


	// Set initiated to true
	LoggerInitiated = true

	return

}