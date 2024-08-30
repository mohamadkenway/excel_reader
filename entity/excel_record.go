package entity

type ExcelRecord struct {
	Origin          string `db:"origin"           json:"origin"`
	OriginCode      int64  `db:"origin_code"      json:"origin_code"`
	Destination     string `db:"destination"      json:"destination"`
	DestinationCode int64  `db:"destination_code" json:"destination_code"`
	CargoCode       int64  `db:"cargo_code"       json:"cargo_code"`
	CargoName       string `db:"cargo_name"       json:"cargo_name"`
	CarCode         int64  `db:"car_code"         json:"car_code"`
	CarName         string `db:"car_name"         json:"car_name"`
	Price           int64  `db:"price"            json:"price"`
}
