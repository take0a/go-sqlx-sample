package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/roboninc/sqlx"
)

// Datatypes は、データ型テストテーブル
type Datatypes struct {
	Col01 sql.NullInt32   `db:"COL01"` // int NOT NULL,
	Col02 sql.NullInt32   `db:"COL02"` // SMALLINT,
	Col03 sql.NullInt64   `db:"COL03"` // BIGINT,
	Col04 sql.NullInt64   `db:"COL04"` // INTEGER,
	Col05 sql.NullString  `db:"COL05"` // DECIMAL(4,2),
	Col06 sql.NullString  `db:"COL06"` // NUMERIC, デフォルトは 5,0
	Col07 sql.NullFloat64 `db:"COL07"` // float,
	Col08 sql.NullFloat64 `db:"COL08"` // double,
	Col09 sql.NullString  `db:"COL09"` // decfloat,
	Col10 sql.NullString  `db:"COL10"` // char(10),
	Col11 sql.NullString  `db:"COL11"` // varchar(10),
	Col12 sql.NullString  `db:"COL12"` // char for bit data,
	Col13 sql.NullString  `db:"COL13"` // clob(10),
	Col14 sql.NullString  `db:"COL14"` // dbclob(100),
	Col15 sql.NullTime    `db:"COL15"` // date,
	Col16 sql.NullTime    `db:"COL16"` // time,
	Col17 sql.NullTime    `db:"COL17"` // timestamp,
	Col18 []byte          `db:"COL18"` // blob(10),
	Col19 sql.NullBool    `db:"COL19"` // boolean,
}

// Key は、データ型テストテーブルのキー
type Key struct {
	Col01 int32 `db:"COL01"`
}

func main() {
	sqlx.BindDriver("go_ibm_db", sqlx.QUESTION)

	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("go_ibm_db", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	// JST, _ := time.LoadLocation("Asia/Tokyo")
	key := Key{1}
	src := Datatypes{
		Col01: sql.NullInt32{Int32: 1, Valid: true},
		Col02: sql.NullInt32{Int32: 2, Valid: true},
		Col03: sql.NullInt64{Int64: 3, Valid: true},
		Col04: sql.NullInt64{Int64: 4, Valid: true},
		Col05: sql.NullString{String: "12.34", Valid: true},
		Col06: sql.NullString{String: "12345", Valid: true},
		Col07: sql.NullFloat64{Float64: 1.5, Valid: true},
		Col08: sql.NullFloat64{Float64: 2.125, Valid: true},
		Col09: sql.NullString{String: "1234567890.12345", Valid: true},
		Col10: sql.NullString{String: "char(10)  ", Valid: true},
		Col11: sql.NullString{String: "varchar", Valid: true},
		// Col12: sql.NullString{String: "31", Valid: true},		input != output
		Col13: sql.NullString{String: "clob(10)", Valid: true},
		// Col14: sql.NullString{String: "ａ", Valid: true},		input != output
		Col15: sql.NullTime{Time: time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC), Valid: true},
		Col16: sql.NullTime{Time: time.Date(0001, 1, 1, 4, 5, 6, 0, time.UTC), Valid: true}, // 0001-01-01固定（SQLServerと同じ）
		Col17: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC), Valid: true},
		Col18: []byte{1, 2, 3},
		Col19: sql.NullBool{Bool: true, Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO datatypes (
			COL01, COL02, COL03, COL04, COL05, COL06, COL07, COL08, COL09, COL10,
			COL11, COL12, COL13, COL14, COL15, COL16, COL17, COL18, COL19
		) VALUES (
			:COL01, :COL02, :COL03, :COL04, :COL05, :COL06, :COL07, :COL08, :COL09, :COL10,
			:COL11, :COL12, :COL13, :COL14, :COL15, :COL16, :COL17, :COL18, :COL19
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := Datatypes{}
	query, args, err := db.BindNamed(`
		SELECT 
			COL01, COL02, COL03, COL04, COL05, COL06, COL07, COL08, COL09, COL10,
			COL11, COL12, COL13, COL14, COL15, COL16, COL17, COL18, COL19
		FROM datatypes 
		WHERE COL01 = :COL01`,
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
		WHERE COL01 = :COL01`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
