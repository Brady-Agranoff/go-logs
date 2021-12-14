package logging

import (
	"fmt"
	"runtime"
)

type ErrorLog struct {
	ID int `pg:",pk"`
	ErrorMessage string
	LineNumber int
	FileName string
	CreatedAt string `pg:"default:now()"`
}

func HandleError(err error) {

	if err != nil {

		_, filename, line, _ := runtime.Caller(1)
		
		// Write Logs to Database
		if Logger.WriteErrorsToDatabase {
			InsertErrorLog(&ErrorLog{
				ErrorMessage: err.Error(),
				LineNumber: line,
				FileName: filename,
			})
		}

		// Send Telegram Notification
		if Logger.NotificationOnAllErrors {

			errorMessage := fmt.Sprintf("New Error in file %s on line %d: %s", filename, line, err.Error())
			if Logger.ServiceName != "" {
				errorMessage = fmt.Sprintf("%s: New Error in file %s on line %d: %s ", Logger.ServiceName, filename, line, err.Error())
			}

			SendMessage(errorMessage, Logger.TelegramNotificationsChannelID)
		}

	}

}

// InsertErrorLog inserts an error log into the database
func InsertErrorLog(errorLog *ErrorLog) error {

	if err := Logger.PGDB.Insert(errorLog); err != nil {
		return err
	}

	return nil

} 