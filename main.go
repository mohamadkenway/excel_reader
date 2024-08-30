package main

import (
	"excel_reader/controller/general_controller"
	"excel_reader/repository/database/postgres"
	"log"
)

func main() {

	postgres.Manager{}.Connect()
	defer postgres.Manager{}.Disconnect()

	general_controller.ReadAndInsertExcelFiles()

	log.Println("ITS DONE")

}
