package main

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func TestDbConnection(t *testing.T){
	//try to make connection to Db
	db, err :=sql.Open("mysql","root:@/mlb_test")
	if err != nil {
		t.Errorf("Connection to sql error")
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		t.Errorf("Connection to sql error")
	}
}

func TestLineupOutput(t * testing.T){
	db, err := sql.Open("mysql","root:@/mlb_test")
	defer db.Close()
	var number int
	rows, err := db.Query("Select lineupNumber from lineups where lineupID=1")
	if err != nil {
		t.Errorf("No query returned")
		return 
	} 
	if rows.Next() {
		rows.Scan(&number)
	} else {
		t.Errorf("No query returned")
		return 
	}

	for rows.Next() { 
		err := rows.Scan(&number)
		if err!= nil {
			t.Errorf("No lineups where inserted")
			return
		}
	defer rows.Close()
	}
}
	

