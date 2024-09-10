package dbconnections

import (
	"TaskManagement/app/globals"
	"database/sql"
	"errors"
	"log"
	"os"
	"slices"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var MainDB *bun.DB

func OpenDB() {

	if _, err := os.Stat(globals.Configuration.Database.Database); errors.Is(err, os.ErrNotExist) {
		_, err := os.OpenFile(globals.Configuration.Database.Database, os.O_CREATE, 0666)
		if err != nil {
			log.Panicln("Error during InitLogger", err)
		}
	}

	sqldb, err := sql.Open(sqliteshim.ShimName, globals.Configuration.Database.Database)
	if err != nil {
		log.Panic("Connection To Database is not Succesfull !!!" + (err.Error()))
	}

	MainDB = bun.NewDB(sqldb, sqlitedialect.New())

	// here if database already contains the tables or not
	rws, err := MainDB.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		// handle error
		log.Println(err)
		return
	}

	defer rws.Close()
	var tableNotExists = false
	i := 0
	for rws.Next() {
		var tableName string
		err := rws.Scan(&tableName)
		if err != nil {
			tableNotExists = true
			break
		}

		if slices.Contains(globals.Configuration.Database.Tables, tableName) {
			i++
		}
	}

	if i != globals.Configuration.Database.TablesCount {
		tableNotExists = true

	}

	if tableNotExists {
		log.Println("Some tables exist in the database")
		f, err := os.ReadFile(globals.Configuration.Database.Migration)
		if err != nil {
			log.Panic("Migration Creation is not Successfully" + (err.Error()))
		}

		// here Seed the database with tables
		_, err = MainDB.Exec(string(f))
		if err != nil {
			log.Panic("Migration Creation is not Successfully" + (err.Error()))
		}
	}

	if globals.Configuration.Database.ForeignEnable {
		// here Seed the database with tables
		_, _ = MainDB.Exec("PRAGMA foreign_keys = ON")
	}

	log.Println("DB Connection Established")
}

func CloseDB() {

	if MainDB.DB != nil {
		err := MainDB.DB.Close()
		if err != nil {
			return
		}
	}
}
