package excel_controller

import (
	"excel_reader/entity"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func ReadExcelFile(path string) []entity.ExcelRecord {
	var list []entity.ExcelRecord
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Println(err.Error() + " in excel file = " + path)
		return []entity.ExcelRecord{}
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Println(err.Error() + " in excel file = " + path)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err.Error() + " in excel file = " + path)
		return []entity.ExcelRecord{}
	}
	for index, row := range rows {
		if index == 0 {
			continue
		}
		excelRecord := entity.ExcelRecord{}
		for collIndex, colCell := range row {
			if collIndex == 0 {
				converted, convertErr := strconv.Atoi(colCell)
				if convertErr != nil {
					log.Println(convertErr)
				}
				excelRecord.OriginCode = int64(converted)
			}
			if collIndex == 1 {
				excelRecord.Origin = colCell
			}
			if collIndex == 2 {
				converted, convertErr := strconv.Atoi(colCell)
				if convertErr != nil {
					log.Println(convertErr)
				}
				excelRecord.DestinationCode = int64(converted)
			}
			if collIndex == 3 {
				excelRecord.Destination = colCell
			}
			if collIndex == 4 {
				converted, convertErr := strconv.Atoi(colCell)
				if convertErr != nil {
					log.Println(convertErr)
				}
				excelRecord.CargoCode = int64(converted)
			}
			if collIndex == 5 {
				excelRecord.CargoName = colCell
			}
			if collIndex == 6 {
				converted, convertErr := strconv.Atoi(colCell)
				if convertErr != nil {
					log.Println(convertErr)
				}
				excelRecord.CarCode = int64(converted)
			}
			if collIndex == 7 {
				excelRecord.CarName = colCell
			}
			if collIndex == 8 {
				converted, convertErr := strconv.Atoi(colCell)
				if convertErr != nil {
					log.Println(convertErr)
				}
				excelRecord.Price = int64(converted)
			}
		}
		list = append(list, excelRecord)

	}
	return list
}
