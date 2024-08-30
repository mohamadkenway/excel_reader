package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Manager struct {
	Connection *sqlx.DB
}

var ManagerInstance Manager

func (_ Manager) Connect() {

	db, connectErr := sqlx.Connect("postgres", "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable")
	if connectErr != nil {
		log.Fatalln(connectErr)
	}
	db.Exec(RentsRepositorySchema)
	ManagerInstance.Connection = db

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatalln(pingErr)
	}

}

func (_ Manager) Disconnect() {

	pingErr := ManagerInstance.Connection.Ping()
	if pingErr != nil {
		log.Fatalln(pingErr)
	}

}
