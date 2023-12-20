module github.com/take0a/go-sqlx-sample

go 1.18

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/ibmdb/go_ibm_db v0.4.5
	github.com/lib/pq v1.10.9
	github.com/microsoft/go-mssqldb v0.21.0
	github.com/roboninc/sqlx v1.4.0
	github.com/sijms/go-ora/v2 v2.7.6
)

// replace github.com/roboninc/sqlx => ../sqlx

require (
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
)
