package general_controller

import (
	"excel_reader/controller/excel_controller"
	"excel_reader/repository/database/postgres"
	"excel_reader/util"
	"log"
)

func ReadAndInsertExcelFiles() {
	rootPath := "/home/file_dir"
	files, err := util.FileManager{}.GetFilePathsFromDir(rootPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		records := excel_controller.ReadExcelFile(file)
		for _, record := range records {
			insertErr := postgres.RentsTable{}.Insert(record)
			if insertErr != nil {
				log.Println(insertErr)
			}
		}
	}
}
