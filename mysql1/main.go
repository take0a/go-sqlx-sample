package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

// Datatypes は、データ型テストテーブル
type Datatypes struct {
	Col01 sql.NullInt32   `db:"col01"` // tinyint
	Col02 sql.NullInt32   `db:"col02"` // smallint
	Col03 sql.NullInt64   `db:"col03"` // mediumint
	Col04 sql.NullInt64   `db:"col04"` // int
	Col05 sql.NullInt64   `db:"col05"` // bigint
	Col06 sql.NullString  `db:"col06"` // decimal(10,4)
	Col07 sql.NullFloat64 `db:"col07"` // float
	Col08 sql.NullFloat64 `db:"col08"` // double
	Col09 []byte          `db:"col09"` // bit(3)
	Col10 sql.NullTime    `db:"col10"` // date
	Col11 sql.NullTime    `db:"col11"` // datetime
	Col12 sql.NullTime    `db:"col12"` // timestamp
	Col13 sql.NullString  `db:"col13"` // time
	Col14 sql.NullInt32   `db:"col14"` // year
	Col15 sql.NullString  `db:"col15"` // char(10)
	Col16 sql.NullString  `db:"col16"` // varchar(10)
	Col17 []byte          `db:"col17"` // binary(10)
	Col18 []byte          `db:"col18"` // varbinary(10)
	Col19 []byte          `db:"col19"` // tinyblob
	Col20 []byte          `db:"col20"` // blob
	Col21 []byte          `db:"col21"` // mediumblob
	Col22 []byte          `db:"col22"` // longblob
	Col23 sql.NullString  `db:"col23"` // tinytext
	Col24 sql.NullString  `db:"col24"` // text
	Col25 sql.NullString  `db:"col25"` // mediumtext
	Col26 sql.NullString  `db:"col26"` // longtext
	Col27 sql.NullString  `db:"col27"` // enum('S', 'M', 'L')
	Col28 sql.NullString  `db:"col28"` // set('R', 'G', 'B')
	Col29 []byte          `db:"col29"` // geometry
	Col30 []byte          `db:"col30"` // point
	Col31 []byte          `db:"col31"` // linestring
	Col32 []byte          `db:"col32"` // polygon
	Col33 []byte          `db:"col33"` // multipoint
	Col34 []byte          `db:"col34"` // multilinestring
	Col35 []byte          `db:"col35"` // multipolygon
	Col36 []byte          `db:"col36"` // geometrycollection
	Col37 sql.NullString  `db:"col37"` // json
}

// Key は、データ型テストテーブルのキー
type Key struct {
	Col01 int32 `db:"col01"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	JST, _ := time.LoadLocation("Asia/Tokyo")
	key := Key{1}
	src := Datatypes{
		Col01: sql.NullInt32{Int32: 1, Valid: true},
		Col02: sql.NullInt32{Int32: 2, Valid: true},
		Col03: sql.NullInt64{Int64: 3, Valid: true},
		Col04: sql.NullInt64{Int64: 4, Valid: true},
		Col05: sql.NullInt64{Int64: 5, Valid: true},
		Col06: sql.NullString{String: "123456.1234", Valid: true},
		Col07: sql.NullFloat64{Float64: 1.5, Valid: true},
		Col08: sql.NullFloat64{Float64: 1.25, Valid: true},
		Col09: []byte{3},
		Col10: sql.NullTime{Time: time.Date(2001, 2, 3, 0, 0, 0, 0, JST), Valid: true}, // loc 指定で JST
		Col11: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 0, JST), Valid: true}, // loc 指定で JST, ミリ秒は0
		Col12: sql.NullTime{Time: time.Date(2001, 2, 3, 4, 5, 6, 0, JST), Valid: true}, // MariaDB は default が Now
		Col13: sql.NullString{String: "12:34:56", Valid: true},
		Col14: sql.NullInt32{Int32: 2000, Valid: true},
		Col15: sql.NullString{String: "char(10)", Valid: true}, // 末尾空白なし
		Col16: sql.NullString{String: "varchar", Valid: true},
		Col17: []byte{1, 2, 3, 4, 5, 6, 7, 8, 0, 0}, // 末尾 0 埋め
		Col18: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		Col19: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		Col20: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		Col21: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		Col22: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		Col23: sql.NullString{String: "tinytext", Valid: true},
		Col24: sql.NullString{String: "text", Valid: true},
		Col25: sql.NullString{String: "mediumtext", Valid: true},
		Col26: sql.NullString{String: "longtext", Valid: true},
		Col27: sql.NullString{String: "S", Valid: true},
		Col28: sql.NullString{String: "R,G,B", Valid: true},
		Col29: []byte{
			0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0,
			0x3F, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0,
			0xBF,
		}, // POINT(1,-1)
		Col37: sql.NullString{String: `{"key": 1}`, Valid: true}, // : の後ろのスペースは１つ
	}

	_, err = db.NamedExec(`
		INSERT INTO datatypes (
			col01, col02, col03, col04, col05, col06, col07, col08, col09, col10,
			col11, col12, col13, col14, col15, col16, col17, col18, col19, col20,
			col21, col22, col23, col24, col25, col26, col27, col28, col29, col30,
			col31, col32, col33, col34, col35, col36, col37
		) VALUES (
			:col01, :col02, :col03, :col04, :col05, :col06, :col07, :col08, :col09, :col10,
			:col11, :col12, :col13, :col14, :col15, :col16, :col17, :col18, :col19, :col20,
			:col21, :col22, :col23, :col24, :col25, :col26, :col27, :col28, :col29, :col30,
			:col31, :col32, :col33, :col34, :col35, :col36, :col37
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
			col31, col32, col33, col34, col35, col36, col37
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
