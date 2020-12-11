package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/anxdev")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*
		var name string
		err = db.QueryRow("SELECT site_name FROM config_site WHERE platform_id=1").Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	*/

	var id int
	var name string
	rows, err := db.Query("SELECT platform_id, site_name FROM config_site where platform_id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}
