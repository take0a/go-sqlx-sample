package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/roboninc/sqlx"
)

// Customer は、顧客
type Customer struct {
	CustomerID int
	Name       string
	Address    string
}

// CustomerKey は、顧客のキー
type CustomerKey struct {
	CustomerID int
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := CustomerKey{1}
	src := Customer{key.CustomerID, "Shohei Otani", "Los Angeles Angels"}

	_, err = db.Exec(`
		INSERT INTO CUSTOMER (
			CUSTOMER_ID, NAME, ADDRESS
		) VALUES (
			$1, $2, $3)`,
		src.CustomerID,
		src.Name,
		src.Address,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := Customer{}
	err = db.QueryRow(`
		SELECT 
			CUSTOMER_ID, NAME, ADDRESS
		FROM CUSTOMER 
		WHERE CUSTOMER_ID = $1`,
		key.CustomerID,
	).Scan(
		&dst.CustomerID,
		&dst.Name,
		&dst.Address,
	)
	if err != nil {
		log.Printf("db.QueryRow error %s", err)
	}
	log.Printf("\nsrc = %#v\ndst = %#v\n", src, dst)

	_, err = db.Exec(`
		DELETE FROM CUSTOMER
		WHERE CUSTOMER_ID = $1`,
		key.CustomerID,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
