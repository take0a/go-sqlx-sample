package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/roboninc/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

// S日ほンｺﾞ表 は、日本語テストテーブル
type S日ほンｺﾞ表 struct {
	F列１     sql.NullString `db:"列１"`     // varchar2
	Fれつ2    sql.NullString `db:"れつ2"`    // nvarchar2
	Fレツ３    sql.NullString `db:"レツ３"`    // long
	Fﾚﾂ4    sql.NullString `db:"ﾚﾂ4"`    // char(10)
	FＲｅｔｓｕ５ sql.NullString `db:"ＲＥＴＳＵ５"` // nchar(10) 全角でも大文字
	F列６     sql.NullString `db:"列６"`     // clob
	F列７     sql.NullString `db:"列７"`     // nclob
}

// Key は、日本語テストテーブルのキー
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

	key := Key{"ｶﾍﾝ長文字列"}
	src := S日ほンｺﾞ表{
		F列１:     sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		Fれつ2:    sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		Fレツ３:    sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		Fﾚﾂ4:    sql.NullString{String: "ｺﾃｲ長文字列   ", Valid: true},
		FＲｅｔｓｕ５: sql.NullString{String: "ｺﾃｲ長文字列   ", Valid: true},
		F列６:     sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		F列７:     sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO 日ほンｺﾞ表 (
			列１, れつ2, レツ３, ﾚﾂ4, ＲＥＴＳＵ５, 列６, 列７
		) VALUES (
			:列１, :れつ2, :レツ３, :ﾚﾂ4, :ＲＥＴＳＵ５, :列６, :列７
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := S日ほンｺﾞ表{}
	query, args, err := db.BindNamed(`
		SELECT 
		列１, れつ2, レツ３, ﾚﾂ4, ＲＥＴＳＵ５, 列６, 列７
		FROM 日ほンｺﾞ表 
		WHERE 列１ = :COL01`,
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
		WHERE 列１ = :COL01`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
