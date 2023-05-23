package main

import (
	"database/sql"
	"database/sql/driver"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	ora "github.com/sijms/go-ora/v2"
)

// Datatypes は、データ型テストテーブル
type Datatypes struct {
	Col01 sql.NullString  `db:"COL01"` // varchar2
	Col02 sql.NullString  `db:"COL02"` // nvarchar2
	Col03 sql.NullInt64   `db:"COL03"` // number(10)
	Col04 sql.NullString  `db:"COL04"` // number(10,4)
	Col05 sql.NullFloat64 `db:"COL05"` // float(10)
	Col06 sql.NullString  `db:"COL06"` // long
	Col07 sql.NullTime    `db:"COL07"` // date
	Col08 sql.NullFloat64 `db:"COL08"` // binary_float
	Col09 sql.NullFloat64 `db:"COL09"` // binary_double
	Col10 OraTimeStamp    `db:"COL10"` // timestamp
	Col11 OraTimeStamp    `db:"COL11"` // timestamp with tz
	Col12 OraTimeStamp    `db:"COL12"` // timestamp with local tz
	Col13 sql.NullString  `db:"COL13"` // interval year to month
	Col14 sql.NullString  `db:"COL14"` // interval day to second
	Col15 []byte          `db:"COL15"` // raw
	// Col16 []byte          `db:"COL16"` // long raw
	Col17 sql.NullString `db:"COL17"` // rowid
	Col18 sql.NullString `db:"COL18"` // urowid
	Col19 sql.NullString `db:"COL19"` // char(10)
	Col20 sql.NullString `db:"COL20"` // nchar(10)
	Col21 sql.NullString `db:"COL21"` // clob
	Col22 sql.NullString `db:"COL22"` // nclob
	Col23 []byte         `db:"COL23"` // blob
	// Col24 []byte          `db:"COL24"` // bfile
}

// Key は、データ型テストテーブルのキー
type Key struct {
	Col01 string `db:"COL01"`
}

func main() {
	sqlx.BindDriver("oracle", sqlx.NAMED)

	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("oracle", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	JST, _ := time.LoadLocation("Asia/Tokyo")
	key := Key{"1"}
	src := Datatypes{
		Col01: sql.NullString{String: "1", Valid: true},
		Col02: sql.NullString{String: "123456789", Valid: true},
		Col03: sql.NullInt64{Int64: 123, Valid: true},
		Col04: sql.NullString{String: "123456.1234", Valid: true},
		Col05: sql.NullFloat64{Float64: 1.25, Valid: true},
		Col06: sql.NullString{String: "abc", Valid: true},
		Col07: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 0, time.UTC), Valid: true}, // 時刻も持つ（仕様通り）
		Col08: sql.NullFloat64{Float64: 2.5, Valid: true},
		Col09: sql.NullFloat64{Float64: 3.125, Valid: true},
		Col10: OraTimeStamp{sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000, time.UTC), Valid: true}}, // マイクロ秒持てない
		Col11: OraTimeStamp{sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000, JST), Valid: true}},      // マイクロ秒もTZも持てない
		Col12: OraTimeStamp{sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000, JST), Valid: true}},      // マイクロ秒もTZも持てない
		Col13: sql.NullString{String: "+01-11", Valid: true},
		Col14: sql.NullString{String: "+09 23:59:59.000001", Valid: true},
		Col15: []byte{1, 2, 3},

		Col17: sql.NullString{String: "AAARXwAAEAAAAKbAAE", Valid: true},
		Col18: sql.NullString{String: "AAARXwAAEAAAAKbAAE", Valid: true},
		Col19: sql.NullString{String: "12345678  ", Valid: true}, // 末尾はスペース
		Col20: sql.NullString{String: "12345678  ", Valid: true}, // 末尾はスペース
		Col21: sql.NullString{String: "CLOB", Valid: true},
		Col22: sql.NullString{String: "NCLOB", Valid: true},
		Col23: []byte{1, 2, 3, 4},
	}

	_, err = db.NamedExec(`
		INSERT INTO DATATYPES (
			col01, col02, col03, col04, col05, col06, col07, col08, col09, col10,
			col11, col12, col13, col14, col15,        col17, col18, col19, col20,
			col21, col22, col23
		) VALUES (
			:COL01, :COL02, :COL03, :COL04, :COL05, :COL06, :COL07, :COL08, :COL09, :COL10,
			:COL11, :COL12, :COL13, :COL14, :COL15,         :COL17, :COL18, :COL19, :COL20,
			:COL21, :COL22, :COL23
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
			col11, col12, col13, col14, col15,        col17, col18, col19, col20,
			col21, col22, col23
		FROM DATATYPES 
		WHERE col01 = :COL01`,
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
		WHERE col01 = :COL01`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}

// OraTimeStamp は、Null を許容する Time
type OraTimeStamp struct {
	sql.NullTime
}

// Scan implements the Scanner interface.
func (ot *OraTimeStamp) Scan(value any) error {
	if value == nil {
		ot.Time, ot.Valid = time.Time{}, false
		return nil
	}
	ot.Valid = true
	switch temp := value.(type) {
	case ora.TimeStamp:
		ot.NullTime.Time = time.Time(temp)
	case *ora.TimeStamp:
		ot.NullTime.Time = time.Time(*temp)
	case ora.TimeStampTZ:
		ot.NullTime.Time = time.Time(temp)
	case *ora.TimeStampTZ:
		ot.NullTime.Time = time.Time(*temp)
	default:
		return ot.NullTime.Scan(value)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (ot OraTimeStamp) Value() (driver.Value, error) {
	if !ot.Valid {
		return nil, nil
	}
	return ora.TimeStampTZ(ot.Time), nil
}
