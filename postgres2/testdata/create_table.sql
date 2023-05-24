DROP TABLE 日ほンｺﾞ表;

CREATE TABLE 日ほンｺﾞ表 (
    列１ varchar(10),
    れつ2 character(10),
    レツ３ json,
    ﾚﾂ4 text,
    Ｒｅｔｓｕ５ xml,
    CONSTRAINT pk_日ほンｺﾞ表 PRIMARY KEY(列１)
);

COMMENT ON TABLE 日ほンｺﾞ表 IS '日本語テスト';
COMMENT ON COLUMN 日ほンｺﾞ表.列１ IS '可変長文字列';
COMMENT ON COLUMN 日ほンｺﾞ表.れつ2 IS '固定長文字列';
COMMENT ON COLUMN 日ほンｺﾞ表.レツ３ IS 'JSONデータ';
COMMENT ON COLUMN 日ほンｺﾞ表.ﾚﾂ4 IS '可変長文字列';
COMMENT ON COLUMN 日ほンｺﾞ表.Ｒｅｔｓｕ５ IS 'XMLデータ';
