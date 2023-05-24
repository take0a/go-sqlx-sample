package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/take0a/go-sqlx-sample/patch"
)

// S日ほンｺﾞ表 は、日本語テストテーブル
type S日ほンｺﾞ表 struct {
	F列１     sql.NullString `db:"列１"`     // varchar(10)
	Fれつ2    sql.NullString `db:"れつ2"`    // character(10)
	Fレツ３    sql.NullString `db:"レツ３"`    // json
	Fﾚﾂ4    sql.NullString `db:"ﾚﾂ4"`    // text
	FＲｅｔｓｕ５ sql.NullString `db:"Ｒｅｔｓｕ５"` // xml
}

// Key は、日本語テストテーブルのキー
type Key struct {
	Col01 string `db:"col1"`
}

func main() {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	key := Key{"ｶﾍﾝ長文字列"}
	src := S日ほンｺﾞ表{
		F列１:     sql.NullString{String: "ｶﾍﾝ長文字列", Valid: true},
		Fれつ2:    sql.NullString{String: "ｺﾃｲ長文字列   ", Valid: true},
		Fレツ３:    sql.NullString{String: `{"日本語キー":"日本語文字列"}`, Valid: true},
		Fﾚﾂ4:    sql.NullString{String: "あぁアｱｶﾞＡａ漢〇€㈱ー～―‐－", Valid: true},
		FＲｅｔｓｕ５: sql.NullString{String: "<xml /><ルートノード></ルートノード>", Valid: true},
	}

	_, err = patch.NamedExec(db, `
		INSERT INTO 日ほンｺﾞ表 (
			列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５
		) VALUES (
			:列１, :れつ2, :レツ３, :ﾚﾂ4, :Ｒｅｔｓｕ５
		)`,
		src,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}

	dst := S日ほンｺﾞ表{}
	query, args, err := db.BindNamed(`
		SELECT 
			列１, れつ2, レツ３, ﾚﾂ4, Ｒｅｔｓｕ５
		FROM 日ほンｺﾞ表 
		WHERE 列１ = :col1`,
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
		WHERE 列１ = :col1`,
		key,
	)
	if err != nil {
		log.Printf("db.Exec error %s", err)
	}
}
