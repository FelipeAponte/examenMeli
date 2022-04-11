// package database allows the connections with a MySql database and
// implements functions to insert and query data
package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MYSQL
)

// connectMysql opens the database to perform operations
func connectMysql() *sql.DB {
	USER := "magneto"
	PASSWD := "examenmeli"
	DB := "meli"

	db, err := sql.Open("mysql", USER+":"+PASSWD+"@tcp(127.0.0.1:3306)/"+DB)

	if err != nil {
		log.Printf("error open database: %s\n", err.Error())
	}

	return db
}

// DnaInsert allows insert data in the dnaVerified table
// and checks that no sql injections occur
func DnaInsert(dna string, isMutant bool) {
	db := connectMysql()
	defer db.Close()

	_, err := db.Exec(`
	INSERT INTO dnaVerified(dnaThreads, isMutant)
	VALUES (?, ?)`, dna, isMutant)

	if err != nil {
		log.Printf("error insert element: %s\n", err.Error())
	}

}

// QueryMutants returns the among of mutants in the table dnaVerified
func QueryMutants() (r int) {
	db := connectMysql()
	defer db.Close()

	err := db.QueryRow(`
	SELECT COUNT(id)
	FROM dnaVerified
	WHERE isMutant=1`).Scan(&r)

	if err != nil {
		log.Printf("error query mutants: %s\n", err.Error())
	}

	return
}

// QueryHumans returns the among of humans in the table dnaVerified
func QueryHumans() (r int) {
	db := connectMysql()
	defer db.Close()

	err := db.QueryRow(`
	SELECT COUNT(id)
	FROM dnaVerified
	WHERE isMutant=0`).Scan(&r)

	if err != nil {
		log.Printf("error query no mutants: %s\n", err.Error())
	}

	return
}
