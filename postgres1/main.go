package main

import (
	"database/sql"
	"log"
	"math"
	"os"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Datatypes は、データ型テストテーブル
type Datatypes struct {
	Col01 sql.NullInt64   `db:"col01"` // bigint
	Col02 sql.NullString  `db:"col02"` // bigserial
	Col03 sql.NullString  `db:"col03"` // bit(1)
	Col04 sql.NullString  `db:"col04"` // varbit(1)
	Col05 sql.NullBool    `db:"col05"` // boolean
	Col06 sql.NullString  `db:"col06"` // box
	Col07 []byte          `db:"col07"` // bytea
	Col08 sql.NullString  `db:"col08"` // varchar(10)
	Col09 sql.NullString  `db:"col09"` // character(10)
	Col10 sql.NullString  `db:"col10"` // cidr
	Col11 sql.NullString  `db:"col11"` // circle
	Col12 sql.NullTime    `db:"col12"` // date
	Col13 sql.NullFloat64 `db:"col13"` // float8
	Col14 sql.NullString  `db:"col14"` // inet
	Col15 sql.NullInt64   `db:"col15"` // integer
	Col16 sql.NullString  `db:"col16"` // interval
	Col17 sql.NullString  `db:"col17"` // json
	Col18 sql.NullString  `db:"col18"` // line
	Col19 sql.NullString  `db:"col19"` // lseg
	Col20 sql.NullString  `db:"col20"` // macaddr
	Col21 sql.NullString  `db:"col21"` // money
	Col22 sql.NullString  `db:"col22"` // numeric(10,4)
	Col23 sql.NullString  `db:"col23"` // path
	Col24 sql.NullString  `db:"col24"` // point
	Col25 sql.NullString  `db:"col25"` // polygon
	Col26 sql.NullFloat64 `db:"col26"` // real
	Col27 sql.NullInt64   `db:"col27"` // smallint
	Col28 sql.NullString  `db:"col28"` // smallserial
	Col29 sql.NullString  `db:"col29"` // serial
	Col30 sql.NullString  `db:"col30"` // text
	Col31 sql.NullTime    `db:"col31"` // time
	Col32 sql.NullTime    `db:"col32"` // timetz
	Col33 sql.NullTime    `db:"col33"` // timestamp
	Col34 sql.NullTime    `db:"col34"` // timestamptz
	Col35 sql.NullString  `db:"col35"` // tsquery
	Col36 sql.NullString  `db:"col36"` // tsvector
	Col37 sql.NullString  `db:"col37"` // txid_snapshot
	Col38 sql.NullString  `db:"col38"` // uuid
	Col39 sql.NullString  `db:"col39"` // xml
}

// Key は、データ型テストテーブルのキー
type Key struct {
	Col01 int64 `db:"col01"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	JST, _ := time.LoadLocation("Asia/Tokyo")
	key := Key{1}
	src := Datatypes{
		Col01: sql.NullInt64{Int64: 1, Valid: true},
		Col02: sql.NullString{String: "1", Valid: true}, // bigserial not null
		Col03: sql.NullString{String: "0", Valid: true},
		Col04: sql.NullString{String: "1", Valid: true},
		Col05: sql.NullBool{Bool: true, Valid: true},
		Col06: sql.NullString{String: "(1,1),(0,0)", Valid: true},
		Col07: []byte{1, 2, 3}, // select 結果は nil でなく empty
		Col08: sql.NullString{String: "varchar", Valid: true},
		Col09: sql.NullString{String: "character ", Valid: true}, // space padding
		Col10: sql.NullString{String: "192.168.1.0/24", Valid: true},
		Col11: sql.NullString{String: "<(0,0),1>", Valid: true},
		Col12: sql.NullTime{Time: time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC), Valid: true},
		Col13: sql.NullFloat64{Float64: 1.5, Valid: true},
		Col14: sql.NullString{String: "192.168.1.1", Valid: true},
		Col15: sql.NullInt64{Int64: math.MaxInt32, Valid: true},
		Col16: sql.NullString{String: "1 year", Valid: true},
		Col17: sql.NullString{String: `{"key":1}`, Valid: true},
		Col18: sql.NullString{String: "{1,2,3}", Valid: true},
		Col19: sql.NullString{String: "[(1,1),(0,0)]", Valid: true},
		Col20: sql.NullString{String: "08:00:2b:01:02:03", Valid: true},
		Col21: sql.NullString{String: "$123.45", Valid: true},
		Col22: sql.NullString{String: "123456.1234", Valid: true},
		Col23: sql.NullString{String: "[(1,1),(0,0)]", Valid: true},
		Col24: sql.NullString{String: "(1,2)", Valid: true},
		Col25: sql.NullString{String: "((1,2),(0,0))", Valid: true},
		Col26: sql.NullFloat64{Float64: 1.25, Valid: true},
		Col27: sql.NullInt64{Int64: math.MinInt16, Valid: true},
		Col28: sql.NullString{String: "2", Valid: true}, // smallserial not null
		Col29: sql.NullString{String: "3", Valid: true}, // serial not null
		Col30: sql.NullString{String: "text", Valid: true},
		Col31: sql.NullTime{Time: time.Date(0, 1, 1, 1, 2, 3, 4000000, time.UTC), Valid: true},
		Col32: sql.NullTime{Time: time.Date(0, 1, 1, 1, 2, 3, 4000000, JST), Valid: true},
		Col33: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000000, time.UTC), Valid: true},
		Col34: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000000, JST), Valid: true},
		Col35: sql.NullString{String: "'super':*", Valid: true},
		Col36: sql.NullString{String: "'Fat' 'Rats' 'The'", Valid: true},
		Col37: sql.NullString{String: "10:20:10,14,15", Valid: true},
		Col38: sql.NullString{String: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", Valid: true},
		Col39: sql.NullString{String: "<xml /><root></root>", Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO datatypes (
			col01, col02, col03, col04, col05, col06, col07, col08, col09, col10,
			col11, col12, col13, col14, col15, col16, col17, col18, col19, col20,
			col21, col22, col23, col24, col25, col26, col27, col28, col29, col30,
			col31, col32, col33, col34, col35, col36, col37, col38, col39
		) VALUES (
			:col01, :col02, :col03, :col04, :col05, :col06, :col07, :col08, :col09, :col10,
			:col11, :col12, :col13, :col14, :col15, :col16, :col17, :col18, :col19, :col20,
			:col21, :col22, :col23, :col24, :col25, :col26, :col27, :col28, :col29, :col30,
			:col31, :col32, :col33, :col34, :col35, :col36, :col37, :col38, :col39
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := Datatypes{}
	query, args, err := db.BindNamed(`
		SELECT 
			col01, col02, col03, col04, col05, col06, col07, col08, col09, col10,
			col11, col12, col13, col14, col15, col16, col17, col18, col19, col20,
			col21, col22, col23, col24, col25, col26, col27, col28, col29, col30,
			col31, col32, col33, col34, col35, col36, col37, col38, col39
		FROM datatypes 
		WHERE col01 = :col01`,
		key,
	)
	if err != nil {
		log.Printf("db.BindNamed error %s", err)
	}
	err = db.QueryRowx(query,
		args...,
	).StructScan(
		&dst,
	)
	if err != nil {
		log.Printf("db.QueryRow error %s", err)
	}
	if !reflect.DeepEqual(src, dst) {
		// log.Printf("\nsrc = %#v\ndst = %#v\n", src, dst)
		diff := cmp.Diff(src, dst)
		if len(diff) > 0 {
			log.Print(diff)
		}
	}

	_, err = db.NamedExec(`
		DELETE FROM datatypes
		WHERE col01 = :col01`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
