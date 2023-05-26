DROP TABLE 日ほンｺﾞ表;

CREATE TABLE 日ほンｺﾞ表 (
    列１ VARCHAR2(10 CHAR),
    れつ2 NVARCHAR2(10),
    レツ３ LONG,
    ﾚﾂ4 CHAR(10 CHAR),
    Ｒｅｔｓｕ５ NCHAR(10),
    列６ CLOB,
    列７ NCLOB,
    CONSTRAINT pk_日ほンｺﾞ表 PRIMARY KEY(列１)
);

COMMENT ON TABLE 日ほンｺﾞ表 IS '日本語テスト';
COMMENT ON COLUMN 日ほンｺﾞ表.列１ IS '最大長が10文字の可変長文字列';
COMMENT ON COLUMN 日ほンｺﾞ表.れつ2 IS '最大長が10文字の可変長Unicode文字列';
COMMENT ON COLUMN 日ほンｺﾞ表.レツ３ IS '最大2GB(231から1を引いたバイト数)の可変長文字データ';
COMMENT ON COLUMN 日ほンｺﾞ表.ﾚﾂ4 IS '長さ10文字の固定長文字データ';
COMMENT ON COLUMN 日ほンｺﾞ表.Ｒｅｔｓｕ５ IS '長さ10文字の固定長文字データ';
COMMENT ON COLUMN 日ほンｺﾞ表.列６ IS 'シングルバイト文字またはマルチバイト・キャラクタを含むキャラクタ・ラージ・オブジェクト';
COMMENT ON COLUMN 日ほンｺﾞ表.列７ IS 'Unicodeキャラクタを含むキャラクタ・ラージ・オブジェクト。';
