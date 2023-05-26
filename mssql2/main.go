package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	_ "github.com/microsoft/go-mssqldb"
)

// S日ほンｺﾞ表 は、日本語テストテーブル
type S日ほンｺﾞ表 struct {
	F列１     sql.NullString `db:"列１"`     // char(10)
	Fれつ2    sql.NullString `db:"れつ2"`    // varchar(10)
	Fレツ３    sql.NullString `db:"レツ３"`    // text
	Fﾚﾂ4    sql.NullString `db:"ﾚﾂ4"`    // nchar(10)
	FＲｅｔｓｕ５ sql.NullString `db:"Ｒｅｔｓｕ５"` // nvarchar(10)
	F列６     sql.NullString `db:"列６"`     // ntext
	F列７     sql.NullString `db:"列７"`     // xml
}

// Key は、日本語テストテーブルのキー
type Key struct {
	Col01 string `db:"col01"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("sqlserver", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := Key{"ｺﾃｲ長文字列                   "}
	src := S日ほンｺﾞ表{
		F列１:     sql.NullString{String: "ｺﾃｲ長文字列                   ", Valid: true},
		Fれつ2:    sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		Fレツ３:    sql.NullString{String: "あぁアｱｶﾞＡａ漢〇㈱ー～―‐－", Valid: true}, // cp932 に € は無い
		Fﾚﾂ4:    sql.NullString{String: "ｺﾃｲ長文字列   ", Valid: true},
		FＲｅｔｓｕ５: sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		F列６:     sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		F列７:     sql.NullString{String: "<xml/><ルートノード>あ</ルートノード>", Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO 日ほンｺﾞ表 (
			列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５, 列６, 列７
		) VALUES (
			:列１, :れつ2, :レツ３, :ﾚﾂ4, :Ｒｅｔｓｕ５, :列６, :列７
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := S日ほンｺﾞ表{}
	query, args, err := db.BindNamed(`
		SELECT 
			列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５, 列６, 列７
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
