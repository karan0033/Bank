package db

import (
	_ "github.com/lib/pq"
)

// var testQuery *Queries
// var testDB *sql.DB

// func TestMain(m *testing.M) {
// 	config, err := utils.LoadConfig("../..")
// 	if err != nil {
// 		log.Fatal("Cannot load config")
// 	}

// 	testDB, err := sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("Not able to make a connection")
// 	}

// 	defer testDB.Close()

// 	testQuery = New(testDB)

// 	os.Exit(m.Run())
// }
