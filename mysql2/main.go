package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/roboninc/sqlx"
)

// S日ほンｺﾞ表 は、日本語テストテーブル
type S日ほンｺﾞ表 struct {
	F列１     sql.NullString `db:"列１"`     // char(10)
	Fれつ2    sql.NullString `db:"れつ2"`    // varchar(10)
	Fレツ３    sql.NullString `db:"レツ３"`    // tinytext
	Fﾚﾂ4    sql.NullString `db:"ﾚﾂ4"`    // text
	FＲｅｔｓｕ５ sql.NullString `db:"Ｒｅｔｓｕ５"` // mediumtext
	F列６     sql.NullString `db:"列６"`     // longtext
	F列７     sql.NullString `db:"列７"`     // enum('小', '中', '大')
	F列８     sql.NullString `db:"列８"`     // set('あか', 'みどり', 'あお')
	F列９     sql.NullString `db:"列９"`     // json
}

// Key は、日本語テストテーブルのキー
type Key struct {
	Col01 string `db:"col01"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := Key{"ｺﾃｲ長文字列"}
	src := S日ほンｺﾞ表{
		F列１:     sql.NullString{String: "ｺﾃｲ長文字列", Valid: true},
		Fれつ2:    sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		Fレツ３:    sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		Fﾚﾂ4:    sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		FＲｅｔｓｕ５: sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		F列６:     sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		F列７:     sql.NullString{String: "小", Valid: true},
		F列８:     sql.NullString{String: "あか,みどり,あお", Valid: true},
		F列９:     sql.NullString{String: `{"日本語キー": "日本語文字列"}`, Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO 日ほンｺﾞ表 (
			列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５, 列６, 列７, 列８, 列９
		) VALUES (
			:列１, :れつ2, :レツ３, :ﾚﾂ4, :Ｒｅｔｓｕ５, :列６, :列７, :列８, :列９
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := S日ほンｺﾞ表{}
	query, args, err := db.BindNamed(`
		SELECT 
		列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５, 列６, 列７, 列８, 列９
		FROM 日ほンｺﾞ表 
		WHERE 列１ = :col01`,
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
		DELETE FROM 日ほンｺﾞ表
		WHERE 列１ = :col01`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
