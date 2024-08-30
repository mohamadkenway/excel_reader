package postgres

import (
	"excel_reader/entity"
	"excel_reader/repository/database/postgres"
)

type RentsRepository struct{}

func (_ RentsRepository) Insert(record entity.ExcelRecord) error {
	return postgres.RentsTable{}.Insert(record)
}
