package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/google/go-cmp/cmp"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/roboninc/sqlx"
)

// Datatypes は、データ型テストテーブル
type Datatypes struct {
	Col01 sql.NullInt64   `db:"col01"` // bigint
	Col02 sql.NullInt64   `db:"col02"` // int
	Col03 sql.NullInt64   `db:"col03"` // smallint
	Col04 sql.NullInt64   `db:"col04"` // tinyint
	Col05 sql.NullBool    `db:"col05"` // bit
	Col06 sql.NullString  `db:"col06"` // decimal(10,4)
	Col07 sql.NullString  `db:"col07"` // money
	Col08 sql.NullString  `db:"col08"` // smallmoney
	Col09 sql.NullFloat64 `db:"col09"` // float(53)
	Col10 sql.NullFloat64 `db:"col10"` // real
	Col11 sql.NullTime    `db:"col11"` // date
	Col12 sql.NullTime    `db:"col12"` // datetime
	Col13 sql.NullTime    `db:"col13"` // datetime2
	Col14 sql.NullTime    `db:"col14"` // datetimeoffset
	Col15 sql.NullTime    `db:"col15"` // smalldatetime
	Col16 sql.NullTime    `db:"col16"` // time
	Col17 sql.NullString  `db:"col17"` // char(10)
	Col18 sql.NullString  `db:"col18"` // varchar(10)
	Col19 sql.NullString  `db:"col19"` // text
	Col20 sql.NullString  `db:"col20"` // nchar(10)
	Col21 sql.NullString  `db:"col21"` // nvarchar(10)
	Col22 sql.NullString  `db:"col22"` // ntext
	Col23 []byte          `db:"col23"` // binary(10)
	Col24 []byte          `db:"col24"` // varbinary(10)
	Col25 []byte          `db:"col25"` // image
	Col26 []byte          `db:"col26"` // uniqueidentifier
	// Col27 []byte          `db:"col27"` // sql_variant mssql: Operand type clash: varbinary(max) is incompatible with sql_variant
	Col28 sql.NullString `db:"col28"` // xml
}

// Key は、データ型テストテーブルのキー
type Key struct {
	Col01 int64 `db:"col01"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("sqlserver", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	JST, _ := time.LoadLocation("Asia/Tokyo")
	key := Key{1}
	src := Datatypes{
		Col01: sql.NullInt64{Int64: 1, Valid: true},
		Col02: sql.NullInt64{Int64: 2, Valid: true},
		Col03: sql.NullInt64{Int64: 3, Valid: true},
		Col04: sql.NullInt64{Int64: 4, Valid: true},
		Col05: sql.NullBool{Bool: true, Valid: true},
		Col06: sql.NullString{String: "123456.1234", Valid: true},
		Col07: sql.NullString{String: "123456.1234", Valid: true},
		Col08: sql.NullString{String: "123456.1234", Valid: true},
		Col09: sql.NullFloat64{Float64: 1.5, Valid: true},
		Col10: sql.NullFloat64{Float64: 2.125, Valid: true},
		Col11: sql.NullTime{Time: time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC), Valid: true},
		Col12: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7000000, time.UTC), Valid: true},      // ミリ秒あり
		Col13: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7008000, time.UTC), Valid: true},      // マイクロ秒あり
		Col14: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 7008000, JST), Valid: true},           // TZあり
		Col15: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 0, 0, time.UTC), Valid: true},            // 秒丸め
		Col16: sql.NullTime{Time: time.Date(0001, 1, 1, 12, 34, 56, 123456700, time.UTC), Valid: true}, // 0001-01-01固定（string は NG）
		Col17: sql.NullString{String: "char(10)  ", Valid: true},
		Col18: sql.NullString{String: "varchar", Valid: true},
		Col19: sql.NullString{String: "text", Valid: true},
		Col20: sql.NullString{String: "nchar(10) ", Valid: true},
		Col21: sql.NullString{String: "nvarchar", Valid: true},
		Col22: sql.NullString{String: "ntext", Valid: true},
		Col23: []byte{1, 2, 3, 0, 0, 0, 0, 0, 0, 0}, // 末尾ゼロ埋め
		Col24: []byte{1, 2, 3, 4},
		Col25: []byte{1, 2, 3, 4, 5},
		Col26: []byte{
			0x6F, 0x96, 0x19, 0xFF, 0x8B, 0x86, 0xD0, 0x11, 0xB4, 0x2D,
			0x00, 0xC0, 0x4F, 0xC9, 0x64, 0xFF,
		},
		Col28: sql.NullString{String: "<xml/><root>a</root>", Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO datatypes (
			col01, col02, col03, col04, col05, col06, col07, col08, col09, col10,
			col11, col12, col13, col14, col15, col16, col17, col18, col19, col20,
			col21, col22, col23, col24, col25, col26,        col28
		) VALUES (
			:col01, :col02, :col03, :col04, :col05, :col06, :col07, :col08, :col09, :col10,
			:col11, :col12, :col13, :col14, :col15, :col16, :col17, :col18, :col19, :col20,
			:col21, :col22, :col23, :col24, :col25, :col26,         :col28
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
			col21, col22, col23, col24, col25, col26,        col28
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
