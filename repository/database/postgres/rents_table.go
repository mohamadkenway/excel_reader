package postgres

import (
	"excel_reader/entity"
)

type RentsTable struct{}

var tableName = "rents"

var RentsRepositorySchema = `
	DROP TABLE Rents;
	CREATE TABLE Rents (
	 	ID  		      SERIAL PRIMARY KEY,
   	  	Origin            TEXT,
   		OriginCode    	  BIGINT,
   	  	Destination       TEXT,
   		DestinationCode   BIGINT,
   	  	CargoCode         BIGINT,
   		CargoName      	  TEXT,
   	  	CarCode           BIGINT,
   		CarName      	  TEXT,
   	  	Price             BIGINT
	);
	CREATE INDEX rent_index1 ON Rents (OriginCode);
	CREATE INDEX rent_index2 ON Rents (DestinationCode);
`

func (_ RentsTable) Insert(record entity.ExcelRecord) error {
	_, err := ManagerInstance.Connection.Exec(
		"INSERT INTO "+tableName+" (Origin,OriginCode,Destination,DestinationCode,CargoCode,CargoName,CarCode,CarName,Price) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		record.Origin,
		record.OriginCode,
		record.Destination,
		record.DestinationCode,
		record.CargoCode,
		record.CargoName,
		record.CarCode,
		record.CarName,
		record.Price,
	)
	return err
}
