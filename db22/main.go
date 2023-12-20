package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/google/go-cmp/cmp"
	_ "github.com/ibmdb/go_ibm_db"
	"github.com/roboninc/sqlx"
)

// S日ほンｺﾞ表 は、日本語テストテーブル
type S日ほンｺﾞ表 struct {
	F列１     sql.NullString `db:"列１"`     // char(30)
	Fれつ2    sql.NullString `db:"れつ2"`    // varchar(30)
	Fレツ３    sql.NullString `db:"レツ３"`    // clob(30)
	Fﾚﾂ4    sql.NullString `db:"ﾚﾂ4"`    // graphic(10)
	FＲｅｔｓｕ５ sql.NullString `db:"ＲＥＴＳＵ５"` // vargraphic(10)
}

// Key は、日本語テストテーブルのキー
type Key struct {
	Col01 string `db:"col01"`
}

func main() {
	sqlx.BindDriver("go_ibm_db", sqlx.QUESTION)

	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("go_ibm_db", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := Key{"ｺﾃｲ長文字列         "}
	src := S日ほンｺﾞ表{
		F列１:  sql.NullString{String: "ｺﾃｲ長文字列         ", Valid: true},
		Fれつ2: sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		// Fレツ３: sql.NullString{String: "あぁアｱｶﾞＡａ漢〇㈱ー～―‐－", Valid: true}, // clob は日本語NGみたい
		Fﾚﾂ4:    sql.NullString{String: "ｺﾃｲ長文字列　　　", Valid: true},
		FＲｅｔｓｕ５: sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
	}

	_, err = db.NamedExec(`
		INSERT INTO 日ほンｺﾞ表 (
			列１, れつ2, レツ３, ﾚﾂ4, ＲＥＴＳＵ５
		) VALUES (
			:列１, :れつ2, :レツ３, :ﾚﾂ4, :ＲＥＴＳＵ５
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := S日ほンｺﾞ表{}
	query, args, err := db.BindNamed(`
		SELECT 
			列１, れつ2, レツ３, ﾚﾂ4, ＲＥＴＳＵ５
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
