package app

import (
	"github.com/ayushjnv1/Gobank/config"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Init() {

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func initDB() (err error) {
	dbConfig := config.Database()

	db, err = sqlx.Open(dbConfig.Driver(), dbConfig.ConnectionURL())
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	// db.SetMaxIdleConns(dbConfig.MaxPoolSize())
	// db.SetMaxOpenConns(dbConfig.MaxOpenConns())
	// db.SetConnMaxLifetime(time.Duration(dbConfig.MaxLifeTimeMins()) * time.Minute)

	return
}

func GetDB() *sqlx.DB {
	return db
}

func Close() {

	db.Close()
}
